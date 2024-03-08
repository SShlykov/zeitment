import {expect, describe, test, vi} from 'vitest'
import {mount} from "@vue/test-utils";
import SideMenu from "@frames/SideMenu/SideMenu.vue";
import SideMenuList from "@frames/SideMenu/SideMenuList.vue";
import SideMenuHead from "@frames/SideMenu/SideMenuHead.vue";
import MenuItems from "@frames/SideMenu/MenuItems/MenuItems.vue";
import ItemLink from "@frames/SideMenu/MenuItems/ItemLink.vue";
import ItemLine from "@frames/SideMenu/MenuItems/ItemLine.vue";
import ItemButton from "@frames/SideMenu/MenuItems/ItemButton.vue";
import {createStore} from "vuex";
import {bookMock} from '@helpers/staticData.js'
import menuList from "@store/modules/layout/menuList.js";
import Router from "@router"

const store = createStore({
  plugins: [],
  modules: {
    books: {
      state: {
        userBooks: []
      },
      mutations: {
        setBooksList() {
        },
        resetStore() {
        }
      },
      actions: {
        async fetchBooks() {
        }
      },
      getters: {
        userBooks: () => [bookMock]
      },
      namespaced: true,
    },
    layout: {
      namespaced: true,
      getters: {
        isSideMenuOpen: () => true,
        menuList: () => menuList,
      },
      actions: {
        toggleMenu() {
        }
      }
    }
  }
})

describe("tests of SideMenu", () => {

  test('mount test of SideMenu', async () => {
    const wrapper = mount(SideMenu, {
      shallow: true,
      global: {
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
    await Router.push('/')
    await Router.isReady()

    const wrapper = mount(SideMenu, {
      plugins: [Router],
      shallow: true,
      global: {
        mocks: {
          $store: store
        },
      }
    })

    console.log(wrapper.html())
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
  test('mount test of MenuItems', async () => {
    const wrapper = mount(MenuItems, {
      shallow: true,
    })

    expect(wrapper.exists()).toBe(true)
  })

  test('mount test of ItemLink', async () => {
    await Router.push('/')
    await Router.isReady()

    const wrapper = mount(ItemLink, {
      shallow: true,
      plugins: [Router],
      global: {
        mocks: {
          $store: store
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
          $store: store
        }
      }
    })

    expect(wrapper.vm.itemFunction()).toBe("22")
    expect(wrapper.exists()).toBe(true)
  })
})

