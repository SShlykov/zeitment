import {expect, test, describe} from 'vitest'
import ServiceOfBooks from '@services/ServiceOfBooks.js'
import {createStore} from "vuex";

class AdapterOfBooks {
  constructor(uri) {
    this.uri = uri
  }

  async getBooks() {
    return [{id: 1, title: 'test'}]
  }
}

describe('serviceOfBooks', () => {
  const store = createStore({
    modules: {
      books: {
        namespaced: true,
        state: {
          booksList: []
        },
        mutations: {
          setBooksList(state, booksList) {
            state.booksList = booksList
          }
        },
        getters: {
          booksList: (state) => state.booksList
        },
        actions: {
          async saveBooks({commit}, booksList) {
            commit('setBooksList', booksList)
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
    await serviceOfBooks.fetchBooksList()
    const booksList = store.getters['books/booksList']
    expect(booksList).toEqual([{id: 1, title: 'test'}])
  })
})