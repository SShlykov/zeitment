import {expect, describe, test, vi, beforeEach} from 'vitest'
import {mount} from "@vue/test-utils";
import SideMenu from "@frames/SideMenu/SideMenu.vue";
import SideMenuList from "@frames/SideMenu/SideMenuList.vue";
import SideMenuHead from "@frames/SideMenu/SideMenuHead.vue";
import MenuItems from "@frames/SideMenu/MenuItems/MenuItems.vue";
import ItemLink from "@frames/SideMenu/MenuItems/ItemLink.vue";
import ItemLine from "@frames/SideMenu/MenuItems/ItemLine.vue";
import ItemButton from "@frames/SideMenu/MenuItems/ItemButton.vue";
import ItemBook from "@frames/SideMenu/MenuItems/ItemBook.vue";
import {createStore} from "vuex";
import {appBook} from "@mocks/books.js";
import { store as books } from '@/store/modules/books';
import { store as layout } from '@/store/modules/layout';
import Router from "@router";

const store = createStore({
  plugins: [],
  modules: {
    books,
    layout
  }
})

describe("tests of SideMenu", () => {

  const mockRoute = {
    params: {
      id: 1
    }
  }
  const mockRouter = {
    push: vi.fn()
  }

  beforeEach(async () => {
    await store.dispatch('books/resetStore')
  })

  test('mount test of SideMenu', async () => {
    await store.dispatch('books/saveUserBooks', [appBook])

    const wrapper = mount(SideMenu, {
      shallow: true,
      global: {
        plugins: [Router],
        mocks: {
          $store: store
        }
      }
    })

    const menuList = wrapper.vm.flattenMenuList
    expect(menuList[0].title).contains("Создать книгу")
    expect(menuList[1].title).contains("Тестовая книга")

    expect(menuList[menuList.length - 1].title).contains("Выход")
  })

  test('render menu items', async () => {
    const wrapper = mount(SideMenu, {
      shallow: true,
      global: {
        plugins: [Router],
        mocks: {
          $store: store
        }
      }
    })

    const flattenMenuList = wrapper.vm.flattenMenuList

    flattenMenuList.forEach((item) => {
      expect(item.title).not.toBe(null)
    })

    expect(flattenMenuList.length).toBeGreaterThan(0)
  })

  test('create book from menu', async () => {

    const wrapper = mount(SideMenu, {
      shallow: true,
      global: {
        plugins: [Router],
        mocks: {
          $store: store,
          $route: mockRoute,
          $router: mockRouter
        },
      }
    })

    expect(wrapper.vm.flattenMenuList[0].title).toBe("Создать книгу")
  })
})

describe("tests of SideMenuList", () => {
  test('mount test of SideMenuList', async () => {
    const wrapper = mount(SideMenuList, {
      shallow: true,
    })

    expect(wrapper.exists()).toBe(true)
  })

})

describe("tests of SideMenuHead", () => {
  test('mount test of SideMenuHead', async () => {
    const wrapper = mount(SideMenuHead, {
      shallow: true,
    })

    expect(wrapper.exists()).toBe(true)
  })
})

describe("tests of MenuItems", () => {
  const mockRoute = {
    params: {
      id: 1
    }
  }
  const mockRouter = {
    push: vi.fn()
  }

  test('mount test of MenuItems', async () => {
    const wrapper = mount(MenuItems, {
      shallow: true,
      global: {
        plugins: [Router],
      },
      props: {
        menuList: []
      }
    })
    expect(wrapper.exists()).toBe(true)
  })

  test('mount test of ItemLink', async () => {
    await Router.isReady()

    const wrapper = mount(ItemLink, {
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

  test('mount test of ItemLine', async () => {
    const wrapper = mount(ItemLine, {
      shallow: true,
    })

    expect(wrapper.exists()).toBe(true)
  })

  test('mount test of ItemBook', async () => {
    const wrapper = mount(ItemBook, {
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

  test('ItemBook remove book by click', async () => {

    const wrapper = mount(ItemBook, {
      shallow: true,
      props: {
        serviceOfBooks: {
          removeBook: () => {
            return appBook
          }
        }
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

    expect(wrapper.vm.serviceOfBooks.removeBook(appBook.id)).toBe(appBook)
  })

  test('mount test of ItemButton', async () => {
    const wrapper = mount(ItemButton, {
      shallow: true,
      props: {
        itemFunction: () => {
          return "22"
        }
      },
      global: {
        mocks: {
          $store: store,
          $route: mockRoute,
          $router: mockRouter
        }
      }
    })

    expect(wrapper.vm.itemFunction()).toBe("22")
    expect(wrapper.exists()).toBe(true)
  })
})

