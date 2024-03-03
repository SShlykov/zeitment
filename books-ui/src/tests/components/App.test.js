import {expect, describe, test, vi} from 'vitest'
import {mount} from "@vue/test-utils";
import App from "@/App.vue";
import {createStore} from "vuex";
import axios  from "axios";
import BooksApi from "@apiServices/BooksApi.js";
import {adapterConfig} from "@store/modules/userBooks/StaticData.js";

vi.mock('axios')
const BooksService = new BooksApi(adapterConfig)

describe("tests of App", async () => {
  const store = createStore({
    plugins: [],
    modules: {
      userBooks: {
        state: {
          booksList: []
        },
        mutations: {
          setBooksList(state, booksList) {
            state.booksList = booksList;
          },
          resetStore() { }
        },
        actions: {
          async fetchBooks({ commit }) {
            const booksList = await BooksService.getBooks()
            commit('setBooksList', booksList)
          }
        },
        getters: {
          booksList: (state) => state.booksList
        },
        namespaced: true,
      },
      layout: {
        actions: {
          initScreenSizeRecalc() { }
        },
      }
    },
  })

  const bookMock = {
    "id": "fb5e7d1d-38cd-4831-bae9-07b36080e3e7",
    "created_at": "2024-03-01T23:47:35.711668+03:00",
    "updated_at": "2024-03-01T23:47:35.711668+03:00",
    "deleted_at": null,
    "owner": "e75aae0d-c1eb-4199-a1d8-2177f57d6a1e",
    "title": "Тестовая книга",
    "author": "Васильев А.В.",
    "description": "test description",
    "is_public": false,
    "publication": null,
    "image_link": null,
    "map_link": null,
    "map_params_id": null,
    "variables": []
  }

  axios.get.mockResolvedValue({data: [ bookMock]})

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
    const expectedBook = {
      "id": "fb5e7d1d-38cd-4831-bae9-07b36080e3e7",
      "createdAt": "2024-03-01T23:47:35.711668+03:00",
      "updatedAt": "2024-03-01T23:47:35.711668+03:00",
      "deletedAt": null,
      "owner": "e75aae0d-c1eb-4199-a1d8-2177f57d6a1e",
      "title": "Тестовая книга",
      "author": "Васильев А.В.",
      "description": "test description",
      "isPublic": false,
      "publication": null,
      "imageLink": null,
      "mapLink": null,
      "mapParamsId": null,
      "variables": []
    }

    const wrapper = mount(App, {
      shallow: true,
      global: {
        mocks: {
          $store: store
        }
      }
    })

    // await wrapper.vm.fetchBooks()
    // const userBooksList = store.getters['userBooks/booksList']
    // expect(userBooksList).toEqual([expectedBook])

    expect(wrapper.vm.booksList).toEqual([expectedBook])
  })
})