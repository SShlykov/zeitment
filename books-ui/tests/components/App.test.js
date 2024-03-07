import {expect, describe, test, vi} from 'vitest'
import {mount} from "@vue/test-utils";
import App from "@/App.vue";
import {createStore} from "vuex";
import axios  from "axios";
import BooksApi from "@apiServices/BooksApi.js";
import {adapterConfig} from "@store/modules/books/StaticData.js";
import {apiBook, apiBooksResponse, appBook} from "@mocks/books.js";

vi.mock('axios')
const BooksService = new BooksApi(adapterConfig)

describe("tests of App", async () => {


  const store = createStore({
    modules: {
      layout: {
        namespaced: true,
        actions: {
          initScreenSizeRecalc() { }
        },
      },
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
            console.log(userBooks)
            commit('setUserBooks', userBooks)
          }
        }
      }
    }
  })


  axios.get.mockResolvedValue({data: apiBooksResponse})

  test('mount test of App', async () => {
    const wrapper = mount(App, {
      shallow: true,
      global: {
        mocks: {
          $store: store
        }
      }
    })

    expect(wrapper.exists()).toBe(true)
  })

  test('fetch userBooks on mount', async () => {
    mount(App, {
      shallow: true,
      global: {
        mocks: {
          $store: store
        }
      }
    })

    const userBooks = store.getters['books/userBooks']

    expect(userBooks).toEqual([appBook])
  })
})