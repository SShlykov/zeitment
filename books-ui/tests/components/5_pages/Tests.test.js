import {test, describe, expect, vi} from 'vitest'
import { mount} from '@vue/test-utils.proto'
import TestsPage from '@pages/Tests/TestsPage.vue';
import {createStore} from "vuex";


describe("tests of TestsPage", () => {
  test('mount test of TestsPage', async () => {

    const store = createStore({
      plugins: [],
      modules: {
        test: {
          namespaced: true,
          state: {
            integrationTestLog: [],
          },
          getters: {
            integrationTestLog:     () => "",
            jsonIntegrationTestLog: () => "",
          },
          mutations: {
            setIntegrationTestLog() {
            }
          },
          actions: {
            async startIntegrationTest() {
            },
          },
        }
      },
    })

    const wrapper = mount(TestsPage, {
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


