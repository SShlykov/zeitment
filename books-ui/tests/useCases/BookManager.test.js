import {describe, expect, test, beforeEach} from 'vitest'
import {createStore} from "vuex";
import {store as books} from "@store/modules/books/index.js";
import {store as layout} from "@store/modules/layout/index.js";
import {store as chapters} from "@store/modules/chapters/index.js";
import {store as pages} from "@store/modules/pages/index.js";
import ServiceOfBooks from '@services/ServiceOfBooks.js'
import {AdapterOfBooks, appBook} from "@mocks/books.js"
import {appTableOfContent, appTableOfContents} from "@mocks/tableOfContent.js"
import {AdapterOfChapters, appChapter} from "@mocks/chapters.js"
import {AdapterOfPages, appPage} from "@mocks/pages.js"
import BookManager from "@useCases/BookManager.js"
import ServiceOfLayout from "@services/ServiceOfLayout.js";
import ServiceOfChapters from "@services/ServiceOfChapters.js";
import ServiceOfPages from "@services/ServiceOfPages.js";

describe('BookManager', () => {
  const store = createStore({
    modules: {
      books,
      layout,
      chapters,
      pages
    }
  })

  beforeEach( async() => {
    await store.dispatch('layout/resetStore')
    await store.dispatch('books/resetStore')
    await store.dispatch('chapters/resetStore')
    await store.dispatch('pages/resetStore')
  })

  const bookAdapter = new AdapterOfBooks('')
  const bookService = new ServiceOfBooks(bookAdapter, store)
  const chapterAdapter = new AdapterOfChapters('')
  const pageAdapter = new AdapterOfPages('')
  const chapterService = new ServiceOfChapters(chapterAdapter, store)
  const pageService = new ServiceOfPages(pageAdapter, store)
  const layoutService = new ServiceOfLayout(store)
  const bookManager = new BookManager(bookService, chapterService, pageService, layoutService)

  test('bookManager is exist', async () => {
    expect(bookManager).toBeDefined()
  })

  test('test of saveBookWithPage', async () => {
    await store.dispatch('books/saveCurrentBook', appBook)
    const {book} = await bookManager.saveBookWithPage(appBook)
    expect(book).toEqual(appBook)
  })

  test('fetchBookWithPage by page', async () => {
    const bookId = "bookId"
    const type = "page"
    const sectionId = "sectionId"
    const {
      tableOfContents,
      book,
      page,
      chapter
    } = await bookManager.fetchBookWithPage(bookId, type, sectionId)


    expect(tableOfContents).toEqual(appTableOfContents)
    expect(book).toEqual(appBook)
    expect(chapter).toBe(null)
    expect(page).toEqual(appPage)

    expect(store.getters['books/currentBook']).toEqual(appBook)
    expect(store.getters['books/tableOfContents']).toEqual(appTableOfContents)
    expect(store.getters['pages/currentPage']).toEqual(appPage)
    expect(store.getters['chapters/currentChapter']).toBe(null)
  })

  test('fetchBookWithPage by chapter', async () => {
    const bookId = "bookId"
    const type = "chapter"
    const sectionId = "sectionId"
    const {
      tableOfContents,
      book,
      chapter,
      page,
    } = await bookManager.fetchBookWithPage(bookId, type, sectionId)

    expect(tableOfContents).toEqual(appTableOfContents)
    expect(book).toEqual(appBook)
    expect(chapter).toEqual(appChapter)
    expect(page).toEqual(null)

    expect(store.getters['books/currentBook']).toEqual(appBook)
    expect(store.getters['books/tableOfContents']).toEqual(appTableOfContents)
    expect(store.getters['pages/currentPage']).toBe(null)
    expect(store.getters['chapters/currentChapter']).toEqual(appChapter)
  })

  test('get currentSectionContent when section = page', async () => {
    store.dispatch('chapters/saveCurrentChapter', null)
    store.dispatch('pages/saveCurrentPage', appPage)

    const {
      title,
      text,
      number
    } = bookManager.getCurrentSectionContent()
    expect(title).toEqual(appPage.title)
    expect(text).toEqual(appPage.text)
    expect(number).toEqual(null)
  })

  test('get currentSectionContent when section = chapter', async () => {
    store.dispatch('chapters/saveCurrentChapter', appChapter)
    store.dispatch('pages/saveCurrentPage', null)

    const {
      title,
      text,
      number
    } = bookManager.getCurrentSectionContent()

    expect(title).toEqual(appChapter.title)
    expect(text).toEqual(appChapter.text)
    expect(number).toEqual(appChapter.number)
  })

  test("test of updateOrderTableOfContent", async () => {
    await store.dispatch('books/saveTableOfContents', appTableOfContents)
    const sortableElement = {
      id: "id",
      number: 1000,
      level: "chapter"
    }
    const updatedTableOfContent = await bookManager.updateOrderTableOfContent(sortableElement)
    expect(updatedTableOfContent).toEqual(appTableOfContent)
  })
})

