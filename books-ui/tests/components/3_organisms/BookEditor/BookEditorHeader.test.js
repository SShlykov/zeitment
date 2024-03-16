import {expect, describe, test, vi} from 'vitest'
import {mount} from "@vue/test-utils.proto";
import BookEditorHeader from "@organisms/BookEditor/BookEditorHeader/BookEditorHeader.vue";
import BookEditorHeaderContainer from "@organisms/BookEditor/BookEditorHeader/BookEditorHeaderContainer.vue";

describe("tests of BookEditorHeader", () => {
  const props = {
    serviceOfBooks: {
      storeEditableBookAttribute: vi.fn(),
      saveEditableBookToServer: vi.fn(),
      fetchEditableBook: vi.fn()
    },
    bookManager: {
      saveBookWithPage: vi.fn()
    },
    editableBook: {
      title: "test",
    }
  }

  test('mount test of BookEditorHeader', async () => {
    const wrapper = mount(BookEditorHeader, {
      shallow: true,
      props
    })

    expect(wrapper.exists()).toBe(true)
  })

  test('test functions in BookEditorHeader', async () => {
    const wrapper = mount(BookEditorHeader, {
      shallow: true,
      props
    })

    const event = {
      target: {
        value: "test"
      }
    }

    expect(wrapper.exists()).toBe(true)
    expect(wrapper.vm.updateBookTitle(event)).toBe("ok")
    expect(wrapper.vm.updateBookAuthor(event)).toBe("ok")
    expect(wrapper.vm.saveBook()).toBe("ok")
  })
})

describe("tests of BookEditorHeaderContainer", () => {
  test('mount test of BookEditorHeaderContainer', async () => {
    const wrapper = mount(BookEditorHeaderContainer, {
      shallow: true,
    })

    expect(wrapper.exists()).toBe(true)
  })
})
