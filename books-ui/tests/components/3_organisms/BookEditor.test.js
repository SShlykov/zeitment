import {expect, describe, test, vi} from 'vitest'
import {mount} from "@vue/test-utils";
import BookEditor from "@organisms/BookEditor/BookEditor.vue";
import BookEditorHeader from "@organisms/BookEditor/BookEditorHeader.vue";
import BookEditorChaptersMenu from "@organisms/BookEditor/BookEditorChaptersMenu.vue";
import BookEditorBody from "@organisms/BookEditor/BookEditorBody.vue";
import { store as books } from '@/store/modules/books';
import {createStore} from "vuex";
import axios from "axios";
import {apiBookResponse} from "@mocks/books.js";

vi.mock('axios')

describe("tests of BookEditor", () => {
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

  test('mount test of BookEditor', async () => {
    axios.get.mockResolvedValue({data: apiBookResponse})

    const wrapper = mount(BookEditor, {
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

describe("tests of BookEditorHeader", () => {
  test('mount test of BookEditorHeader', async () => {
    const wrapper = mount(BookEditorHeader, {
      shallow: true,
    })

    expect(wrapper.exists()).toBe(true)
  })
})

describe("tests of BookEditorChaptersMenu", () => {
  test('mount test of BookEditorChaptersMenu', async () => {
    const wrapper = mount(BookEditorChaptersMenu, {
      shallow: true,
    })

    expect(wrapper.exists()).toBe(true)
  })
})

describe("tests of BookEditorBody", () => {
  test('mount test of BookEditorBody', async () => {
    const wrapper = mount(BookEditorBody, {
      shallow: true,
    })

    expect(wrapper.exists()).toBe(true)
  })
})