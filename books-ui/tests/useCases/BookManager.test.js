import {describe, expect, test} from 'vitest'
import {createStore} from "vuex";
import {store as books} from "@store/modules/books/index.js";
import ServiceOfBooks from '@services/ServiceOfBooks.js'
import {AdapterOfBooks, appBook} from "@mocks/books.js"
import BookManager from "@useCases/BookManager.js"

describe('BookManager', () => {
  const store = createStore({
    modules: {
      books
    }
  })

  const bookAdapter = new AdapterOfBooks('')
  const bookService = new ServiceOfBooks(bookAdapter, store)
  const bookManager = new BookManager(bookService)

  test('bookManager is exist', async () => {
    expect(bookManager).toBeDefined()
  })

  test('test of saveBookWithPage', async () => {
    const {book} = await bookManager.saveBookWithPage(appBook)
    expect(book).toEqual(appBook)
  })
})