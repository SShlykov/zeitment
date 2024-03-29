import {expect, describe, test, vi} from 'vitest'
import {mount} from "@vue/test-utils";
import App from "@/App.vue";
import {createStore} from "vuex";
import axios  from "axios";
import {apiBooksResponse, appBook} from "@mocks/books.js";
import { store as books } from '@/store/modules/books';
import { store as layout } from '@/store/modules/layout';
import Router from "@router";

vi.mock('axios')

describe("tests of App", async () => {
  const store = createStore({
    modules: {
      layout,
      books
    }
  })
  axios.post.mockResolvedValue({data: apiBooksResponse})
  // axios.post.mockImplementation((v, v1) => {
  //   console.log({v, v1})
  //   return {data: apiBooksResponse}
  // })

  test('mount test of App', async () => {
    const wrapper = mount(App, {
      shallow: true,
      global: {
        plugins: [Router],
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
        plugins: [Router],
        mocks: {
          $store: store
        }
      }
    })

    const userBooks = store.getters['books/userBooks']

    expect(userBooks).toEqual([appBook])
  })
})