import {expect, test, describe} from 'vitest'
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
          }
        },
        getters: {
          userBooks: (state) => state.userBooks
        },
        actions: {
          async saveUserBooks({commit}, userBooks) {
            commit('setUserBooks', userBooks)
          }
        }
      }
    }
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
})