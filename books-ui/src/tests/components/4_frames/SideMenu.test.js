import {expect, describe, test} from 'vitest'
import {mount} from "@vue/test-utils";
import SideMenu from "@frames/SideMenu/SideMenu.vue";
import SideMenuList from "@frames/SideMenu/SideMenuList.vue";
import SideMenuHead from "@frames/SideMenu/SideMenuHead.vue";
import {createStore} from "vuex";
import {bookMock} from '@helpers/staticData.js'
import menuList from "@store/modules/layout/menuList.js";

describe("tests of SideMenu", () => {

  const store = createStore({
    plugins: [],
    modules: {
      userBooks: {
        state: {
          booksList: []
        },
        mutations: {
          setBooksList() { },
          resetStore() { }
        },
        actions: {
          async fetchBooks() { }
        },
        getters: {
          booksList: () => [bookMock]
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
          toggleMenu() { }
        }
      }
    }
  })

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

    expect(wrapper.exists()).toBe(true)
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