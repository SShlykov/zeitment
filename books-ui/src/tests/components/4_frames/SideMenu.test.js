import {expect, describe, test} from 'vitest'
import {mount} from "@vue/test-utils";
import SideMenu from "@frames/SideMenu/SideMenu.vue";
import SideMenuList from "@frames/SideMenu/SideMenuList.vue";
import SideMenuHead from "@frames/SideMenu/SideMenuHead.vue";
import {createStore} from "vuex";


describe("tests of SideMenu", () => {
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
          menuList: () => [],
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