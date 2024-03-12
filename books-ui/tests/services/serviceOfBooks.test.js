import {expect, test, describe, beforeEach} from 'vitest'
import ServiceOfBooks from '@services/ServiceOfBooks.js'
import {createStore} from "vuex";
import {AdapterOfBooks, appBook} from "@mocks/books.js"


describe('serviceOfBooks', () => {
  const store = createStore({
    modules: {
      books: {
        namespaced: true,
        state: {
          userBooks: []
        },
        mutations: {
          setUserBooks(state, userBooks) {
            state.userBooks = userBooks
          },
          resetStore(state) {
            state.userBooks = []
          }
        },
        getters: {
          userBooks: (state) => state.userBooks
        },
        actions: {
          async saveUserBooks({commit}, userBooks) {
            commit('setUserBooks', userBooks)
          },
          resetStore({commit}) {
            commit('resetStore')
          }
        }
      }
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

  test("remove book", async () => {
    const book = await serviceOfBooks.removeBook(appBook.id)
    expect(book).toEqual(appBook)
  })
})