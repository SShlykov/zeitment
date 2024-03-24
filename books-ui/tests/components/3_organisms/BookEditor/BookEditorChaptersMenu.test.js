import {expect, describe, test, vi} from 'vitest'
import {mount} from "@vue/test-utils";
import BookEditorMenu from "@organisms/BookEditor/BookEditorMenu/BookEditorMenu.vue";
import BookEditorMenuItem from "@organisms/BookEditor/BookEditorMenu/BookEditorMenuItem.vue";
import Router from "@router";
import {bookPageConfig} from "@mocks/books.js";
import book from "@pages/Book/Book.vue";


vi.mock('axios')

const mockRoute = {
  params: {
    book_id: "book_id",
    type: "page",
    section_id: "section_id"
  }
}
const mockRouter = {
  push: vi.fn()
}

describe('tests of BookEditorMenu', () => {

  test('mount test of BookEditorMenu', async () => {
    const wrapper = mount(BookEditorMenu, {
      props: {
        pageConfig: bookPageConfig
      },
      global: {
        plugins: [Router],
        mocks: {
          $route: mockRoute,
          $router: mockRouter
        }
      }
    })

    expect(wrapper.exists()).toBe(true)
  })

  test('BookEditorMenu toggle', async () => {
    const wrapper = mount(BookEditorMenu, {
      props: {
        pageConfig: bookPageConfig
      },
      shallow: true,
      global: {
        plugins: [Router],
        mocks: {
          $route: mockRoute,
          $router: mockRouter
        }
      }
    })

    expect(wrapper.vm.isOpen).toBe(true)
    wrapper.vm.toggle()
    expect(wrapper.vm.isOpen).toBe(false)
  })

  test('BookEditorMenu menu list', async () => {
    const wrapper = mount(BookEditorMenu, {
      global: {
        plugins: [Router],
        mocks: {
          $route: mockRoute,
          $router: mockRouter
        }
      },
      props: {
        menuItems: [
          {
            title: "test",
            link: "/test",
            section_id: "section_id",
            id: "id",
            level: "chapter"
          }
        ],
        pageConfig: bookPageConfig
      }
    })

    expect(wrapper.text()).toContain("test")
  })
})

describe('tests of BookEditorMenuItem', () => {
  test('mount test of BookEditorMenuItem', async () => {
    const wrapper = mount(BookEditorMenuItem, {
      props: {
        title: "title",
        id: "section_id",
        level: "chapter",
        bookId: bookPageConfig.bookId,
        sectionId: bookPageConfig.sectionId

      },
      shallow: true,
      global: {
        plugins: [Router],
        mocks: {
          $route: mockRoute,
          $router: mockRouter
        }
      }
    })

    expect(wrapper.exists()).toBe(true)
  })

  test('mount test of BookEditorMenuItem with selected item', async () => {
    const wrapper = mount(BookEditorMenuItem, {
      props: {
        title: "title",
        id: "section_id",
        level: "chapter",
        bookId: bookPageConfig.bookId,
        sectionId: bookPageConfig.sectionId
      },
      global: {
        plugins: [Router],
        mocks: {
          $route: mockRoute,
          $router: mockRouter
        }
      }
    })

    expect(wrapper.html()).contains('bg-gray-100')
    expect(wrapper.vm.sectionId).contains('section_id')
    expect(wrapper.exists()).toBe(true)
  })

  test('mount test of BookEditorMenuItem without selected item', async () => {
    const wrapper = mount(BookEditorMenuItem, {
      props: {
        title: "title",
        id: "bad_id",
        level: "chapter",
        bookId: bookPageConfig.bookId,
        sectionId: bookPageConfig.sectionId
      },
      global: {
        plugins: [Router],
        mocks: {
          $route: mockRoute,
          $router: mockRouter
        }
      }
    })

    expect(wrapper.html()).not.contains('bg-slate-100')
    expect(wrapper.exists()).toBe(true)
  })
})