import {test, describe, expect} from 'vitest'
import { mount} from '@vue/test-utils.proto'
import BooksCompendium from '@pages/BooksCompendium/BooksCompendium.vue';

describe("tests of BooksCompendium", () => {
  test('mount test of BooksCompendium', async () => {

    const wrapper = mount(BooksCompendium, {
      shallow: true,
    })

    expect(wrapper.exists()).toBe(true)
  })
})


