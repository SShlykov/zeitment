import {expect, describe, test, vi, beforeEach} from 'vitest'
import {mount} from "@vue/test-utils";
import BookEditor from "@organisms/BookEditor/BookEditor.vue";
import BookEditorHeader from "@organisms/BookEditor/BookEditorHeader.vue";
import BookEditorChaptersMenu from "@organisms/BookEditor/BookEditorChaptersMenu.vue";
import BookEditorBody from "@organisms/BookEditor/BookEditorBody.vue";
import { store as books } from '@store/modules/books/index.js';
import {createStore} from "vuex";
import axios from "axios";
import Router from "@router";
import {bookPageConfig} from "@mocks/books.js";

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

  beforeEach(() => {
    store.commit('books/resetStore')
  })

  axios.post.mockResolvedValue({data: []})
  axios.get.mockResolvedValue({data: []})


  test('mount test without book of BookEditor', async () => {
    const wrapper = mount(BookEditor, {
      props: {
        serviceOfBooks: {},
        bookManager: {},
        pageConfig: bookPageConfig
      },
      global: {
        plugins: [Router],
        mocks: {
          $store: store,
          $route: mockRoute,
          $router: mockRouter
        }
      }
    })

    expect(wrapper.vm.pageConfig.bookId).toBe("book_id")
    expect(wrapper.vm.pageConfig.sectionId).toBe("section_id")
    expect(wrapper.vm.pageConfig.type).toBe("page")
    expect(wrapper.exists()).toBe(true)
    expect(wrapper.text()).contains('Книги не существует')
  })

  test('mount test of BookEditor', async () => {
    store.dispatch('books/saveCurrentBook', {id: 1, title: "qwerty", author: "qwerty"})

    const wrapper = mount(BookEditor, {
      props: {
        serviceOfBooks: {},
        bookManager: {},
        pageConfig: bookPageConfig
      },
      global: {
        plugins: [Router],
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