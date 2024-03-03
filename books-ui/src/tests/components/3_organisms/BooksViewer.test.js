import {test, describe, expect, vi} from 'vitest'
import { mount} from '@vue/test-utils'
import {createStore} from "vuex";
import BooksViewer from '@organisms/BooksViewer/BooksViewer.vue';
import BooksViewerContainer from '@organisms/BooksViewer/BooksViewerContainer.vue';
import BooksViewerCard from '@organisms/BooksViewer/BooksViewerCard.vue';

describe("tests of BooksViewer", () => {
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
          booksList: () => []
        },
        namespaced: true,
      }
    },
  })


  test('mount test of BooksViewer', async () => {

    const wrapper = mount(BooksViewer, {
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

describe("tests of BooksViewerContainer", () => {
  test('mount test of BooksViewerContainer', async () => {

    const wrapper = mount(BooksViewerContainer, {
      shallow: true,
    })

    expect(wrapper.exists()).toBe(true)
  })
})

describe("tests of BooksViewerCard", () => {
  test('mount test of BooksViewerCard', async () => {

    const wrapper = mount(BooksViewerCard, {
      shallow: true,
      props: { }
    })

    expect(wrapper.exists()).toBe(true)
  })

  test('display date and title', async () => {
    const updated_at = "2024-02-12T23:47:35.711668+03:00"
    vi.setSystemTime(new Date('2024-02-20T23:47:35'))

    const wrapper = mount(BooksViewerCard, {
      shallow: true,
      props: {
        id: "fb5e7d1d-38cd-4831-bae9-07b36080e3e7",
        owner: "e75aae0d-c1eb-4199-a1d8-2177f57d6a1e",
        title: "Тестовая книга",
        updated_at,
        variables: [],
        map_params_id: null,
        map_link: null,
        image_link: null,
        publication: null,
        is_public: false,
        description: "test description",
      }
    })


    expect(wrapper.text().includes("Тестовая книга")).toBe(true)
    expect(wrapper.text().includes("Обновлено 8 дней назад")).toBe(true)
  })

  test('display date and title', async () => {
    const updated_at = "2024-02-12T23:47:35.711668+03:00"
    vi.setSystemTime(new Date('2024-07-20T23:47:35'))

    const wrapper = mount(BooksViewerCard, {
      shallow: true,
      props: {
        id: "fb5e7d1d-38cd-4831-bae9-07b36080e3e7",
        owner: "e75aae0d-c1eb-4199-a1d8-2177f57d6a1e",
        title: "Тестовая книга",
        updated_at,
        variables: [],
        map_params_id: null,
        map_link: null,
        image_link: null,
        publication: null,
        is_public: false,
        description: "test description",
      }
    })


    expect(wrapper.text().includes("Тестовая книга")).toBe(true)
    expect(wrapper.text().includes("Обновлено 5 месяцев назад")).toBe(true)
  })

  test('display date', async () => {
    const updated_at = "2024-02-12T23:47:35.711668+03:00"
    vi.setSystemTime(new Date('2025-07-20T23:47:35'))

    const wrapper = mount(BooksViewerCard, {
      shallow: true,
      props: {
        updated_at,
      }
    })

    const el = wrapper.find('[test-id="lastUpdated"]')

    expect(el.text()).toBe('Обновлено год назад')
  })
})

