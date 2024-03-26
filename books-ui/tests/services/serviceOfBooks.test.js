import {expect, test, describe, beforeEach, expectTypeOf} from 'vitest'
import {createStore} from "vuex";
import { store as books } from '@/store/modules/books';
import ServiceOfBooks from '@services/ServiceOfBooks.js'
import {AdapterOfBooks, appBook} from "@mocks/books.js"
import {appTableOfContent, appTableOfContents} from "@mocks/tableOfContent.js"

describe('serviceOfBooks', () => {
  const store = createStore({
    modules: {
      books
    }
  })

  beforeEach(async () => {
    await store.dispatch('books/resetStore')
  })

  const adapterOfBooks = new AdapterOfBooks('');
  const serviceOfBooks = new ServiceOfBooks(adapterOfBooks, store);

  test('serviceOfBooks is exist', () => {
    expect(serviceOfBooks).toBeDefined()
  })

  test("fetch booksList", async () => {
    const booksList = await serviceOfBooks.fetchUserBooks()
    expect(booksList).toEqual([appBook])
  })

  test("get booksList from store", async () => {
    await store.dispatch('books/saveUserBooks', [appBook])
    const booksList = await serviceOfBooks.getUserBooks()
    expect(booksList).toEqual([appBook])
  })

  test("get current book", () => {
    store.dispatch('books/saveCurrentBook', appBook)
    const currentBook = serviceOfBooks.getCurrentBook()
    expect(currentBook).toEqual(appBook)
  })

  test("put booksList", async () => {
    await serviceOfBooks.putUserBooks([appBook])
    const booksList = store.getters['books/userBooks']
    expect(booksList).toEqual([appBook])
  })

  test("make fetch and put booksList", async () => {
    const booksList = await serviceOfBooks.makeFetchAndPutUserBooks()
    const storeBooksList = store.getters['books/userBooks']
    expect(booksList).toEqual([appBook])
    expect(storeBooksList).toEqual([appBook])
  })

  test("create book", async () => {
    let books = store.getters['books/userBooks']
    expect(books).toEqual([])

    await serviceOfBooks.createBook(appBook)

    books = store.getters['books/userBooks']
    expect(books).toEqual([appBook])
  })

  test("update book", async () => {
    const updatedBook = {
      ...appBook,
      title: "new title"
    }
    const book = await serviceOfBooks.updateBook(updatedBook)
    expect(book).toEqual(updatedBook)
  })

  test("remove book", async () => {
    const book = await serviceOfBooks.removeBook(appBook.id)
    expect(book).toEqual(appBook)
  })

  test("test of putBookAttribute", async () => {
    await store.dispatch('books/saveCurrentBook', appBook)
    await serviceOfBooks.putBookAttribute("title", "new title")

    const editableBook = store.getters['books/currentBook']

    expect(editableBook).toEqual({
      ...appBook,
      title: "new title"
    })
  })

  test("test of putBookAttribute with no editableBook", async () => {
    await store.dispatch('books/saveCurrentBook', null)
    const editableBook = store.getters['books/currentBook']
    expect(editableBook).toEqual(null)
    const updatedBook = await serviceOfBooks.putBookAttribute("title", "new title")
    expect(updatedBook).toEqual(null)
  })

  test("get book by id", async () => {
    const book = await serviceOfBooks.fetchBookById(appBook.id)
    expect(book).toEqual(appBook)
  })

  test("save editable book to server", async () => {
    await store.dispatch('books/saveCurrentBook', appBook)
    await serviceOfBooks.putBookAttribute("title", "new title")
    const storedEditableBook = store.getters['books/currentBook']

    expectTypeOf(storedEditableBook.id).toBeString()
    await serviceOfBooks.saveCurrentBook()

    let book = await serviceOfBooks.fetchBookById(appBook.id)
    book = {
      ...book,
      title: "new title"
    }
    expect(book).toEqual(storedEditableBook)
  })

  test("save editable book to server with no editableBook", async () => {
    await store.dispatch('books/saveCurrentBook', null)
    const storedEditableBook = store.getters['books/currentBook']
    expect(storedEditableBook).toEqual(null)

    const updatedBook = await serviceOfBooks.saveCurrentBook()
    expect(updatedBook).toEqual(null)
  })

  test("fetch editable book", async () => {
    const book = await serviceOfBooks.fetchEditableBook(appBook.id)
    const editableBook = store.getters['books/currentBook']
    expect(editableBook).toEqual(appBook)
    expect(book).toEqual(appBook)
  })

  test("fetch table of contents", async () => {
    store.dispatch('books/saveTableOfContent', appTableOfContent)
    const tableOfContents = await serviceOfBooks.fetchTableOfContents(appBook.id)
    const storedTableOfContents = store.getters['books/tableOfContents']
    expect(storedTableOfContents).toEqual(appTableOfContent)
    expect(tableOfContents).toEqual(appTableOfContent)
  })

  test("put table of contents", async () => {
    await serviceOfBooks.putTableOfContents(appTableOfContent)
    const storedTableOfContents = store.getters['books/tableOfContents']
    expect(storedTableOfContents).toEqual(appTableOfContent)
  })

  test("make fetch and put table of contents", async () => {
    const tableOfContents = await serviceOfBooks.makeFetchAndPutTableOfContents(appBook.id)
    const storedTableOfContents = store.getters['books/tableOfContents']
    expect(tableOfContents).toEqual(appTableOfContent)
    expect(storedTableOfContents).toEqual(appTableOfContent)
  })

  test("fetch currentBook", async () => {
    const book = await serviceOfBooks.fetchCurrentBook(appBook.id)
    const editableBook = store.getters['books/currentBook']
    const tableOfContents = store.getters['books/tableOfContents']
    expect(editableBook).toEqual(appBook)
    expect(tableOfContents).toEqual(appTableOfContent)
    expect(book.book).toEqual(appBook)
    expect(book.tableOfContents).toEqual(appTableOfContent)
  })

})