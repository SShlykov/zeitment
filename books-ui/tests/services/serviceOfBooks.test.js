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
    await serviceOfBooks.fetchUserBooks()
    const booksList = store.getters['books/userBooks']
    expect(booksList).toEqual([appBook])
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

  test("test of storeEditableBookAttribute", async () => {
    await store.dispatch('books/saveCurrentBook', appBook)
    await serviceOfBooks.storeEditableBookAttribute("title", "new title")

    const editableBook = store.getters['books/currentBook']

    expect(editableBook).toEqual({
      ...appBook,
      title: "new title"
    })
  })

  test("test of storeEditableBookAttribute with no editableBook", async () => {
    await store.dispatch('books/saveCurrentBook', null)
    const editableBook = store.getters['books/currentBook']
    expect(editableBook).toEqual(null)
    const updatedBook = await serviceOfBooks.storeEditableBookAttribute("title", "new title")
    expect(updatedBook).toEqual(null)
  })

  test("get book by id", async () => {
    const book = await serviceOfBooks.getBookById(appBook.id)
    expect(book).toEqual(appBook)
  })

  test("save editable book to server", async () => {
    await store.dispatch('books/saveCurrentBook', appBook)
    await serviceOfBooks.storeEditableBookAttribute("title", "new title")
    const storedEditableBook = store.getters['books/currentBook']

    expectTypeOf(storedEditableBook.id).toBeString()
    await serviceOfBooks.saveCurrentBookToServer()

    let book = await serviceOfBooks.getBookById(appBook.id)
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

    const updatedBook = await serviceOfBooks.saveCurrentBookToServer()
    expect(updatedBook).toEqual(null)
  })

  test("fetch editable book", async () => {
    const book = await serviceOfBooks.fetchEditableBook(appBook.id)
    const editableBook = store.getters['books/currentBook']
    expect(editableBook).toEqual(appBook)
    expect(book).toEqual(appBook)
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