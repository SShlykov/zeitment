import {expect, describe, test, vi, beforeEach} from 'vitest'
import {mount} from "@vue/test-utils";
import BookEditorMenu from "@organisms/BookEditor/BookEditorMenu/BookEditorMenu.vue";
import BookEditorMenuItem from "@organisms/BookEditor/BookEditorMenu/BookEditorMenuItem.vue";


vi.mock('axios')

describe('tests of BookEditorMenu', () => {
  test('mount test of BookEditorMenu', async () => {
    const wrapper = mount(BookEditorMenu, {
    })

    expect(wrapper.exists()).toBe(true)
  })

  test('BookEditorMenu toggle', async () => {
    const wrapper = mount(BookEditorMenu, {
      shallow: true,
    })

    expect(wrapper.vm.isOpen).toBe(true)
    wrapper.vm.toggle()
    expect(wrapper.vm.isOpen).toBe(false)
  })

  test('BookEditorMenu menu list', async () => {
    const wrapper = mount(BookEditorMenu, {
      shallow: true,
      props: {
        menuItems: [
          {
            title: "test",
            link: "/test"
          }
        ]
      }
    })

    // expect(wrapper.text()).toContain("test")
  })
})

describe('tests of BookEditorMenuItem', () => {
  test('mount test of BookEditorMenuItem', async () => {
    const wrapper = mount(BookEditorMenuItem, {
      shallow: true,
    })

    expect(wrapper.exists()).toBe(true)
  })
})