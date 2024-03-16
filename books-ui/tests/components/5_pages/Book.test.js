import {test, describe, expect, vi} from 'vitest'
import { mount} from '@vue/test-utils.proto'
import BookPage from '@pages/Book/Book.vue';
import {createStore} from "vuex";
import {store as books} from "@store/modules/books/index.js";
import axios from "axios";
import {apiBookResponse} from "@mocks/books.js";

vi.mock('axios')

describe("tests of BookPage", () => {
  const store = createStore({
    plugins: [],
    modules: {
      books
    },
  })

  const mockRoute = {
    params: {
      id: 1
    }
  }
  const mockRouter = {
    push: vi.fn()
  }

  test('mount test of BookPage', async () => {
    axios.get.mockResolvedValue({data: apiBookResponse})
    const wrapper = mount(BookPage, {
      shallow: true,
      global: {
        mocks: {
          $store: store,
          $route: mockRoute,
          $router: mockRouter
        }
      }
    })

    expect(wrapper.exists()).toBe(true)
  })
})


