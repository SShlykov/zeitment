import {test, describe, expect} from 'vitest'
import { mount} from '@vue/test-utils'
import BookPage from '@pages/Book/Book.vue';

describe("tests of BookPage", () => {
  test('mount test of BookPage', async () => {

    const wrapper = mount(BookPage, {
      shallow: true,
    })

    expect(wrapper.exists()).toBe(true)
  })
})


