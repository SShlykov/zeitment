import {expect, describe, test} from 'vitest'
import {mount} from "@vue/test-utils";
import BookEditor from "@organisms/BookEditor/BookEditor.vue";
import BookEditorHeader from "@organisms/BookEditor/BookEditorHeader.vue";
import BookEditorChaptersMenu from "@organisms/BookEditor/BookEditorChaptersMenu.vue";
import BookEditorBody from "@organisms/BookEditor/BookEditorBody.vue";

describe("tests of BookEditor", () => {
  test('mount test of BookEditor', async () => {
    const wrapper = mount(BookEditor, {
      shallow: true,
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