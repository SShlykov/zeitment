import {describe, expect, test, beforeEach} from 'vitest'
import {createStore} from "vuex";
import {store as books} from "@store/modules/books/index.js";
import {store as layout} from "@store/modules/layout/index.js";
import ServiceOfBooks from '@services/ServiceOfBooks.js'
import {AdapterOfBooks, appBook} from "@mocks/books.js"
import BookManager from "@useCases/BookManager.js"
import ServiceOfLayout from "@services/ServiceOfLayout.js";

describe('BookManager', () => {
  const store = createStore({
    modules: {
      books,
      layout
    }
  })

  beforeEach( async() => {
    await store.dispatch('books/resetStore')
    await store.dispatch('layout/resetStore')
  })

  const bookAdapter = new AdapterOfBooks('')
  const bookService = new ServiceOfBooks(bookAdapter, store)
  const layoutService = new ServiceOfLayout(store)
  const bookManager = new BookManager(bookService, layoutService)

  test('bookManager is exist', async () => {
    expect(bookManager).toBeDefined()
  })

  test('test of saveBookWithPage', async () => {
    await store.dispatch('books/saveCurrentBook', appBook)
    const {book} = await bookManager.saveBookWithPage(appBook)
    expect(book).toEqual(appBook)
  })
})