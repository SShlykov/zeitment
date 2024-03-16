import {test, describe, expect} from 'vitest'
import { mount} from '@vue/test-utils.proto'
import UserBooks from '@pages/UserBooks/UserBooks.vue';

describe("tests of UserBooks", () => {
  test('mount test of UserBooks', async () => {

    const wrapper = mount(UserBooks, {
      shallow: true,
    })

    expect(wrapper.exists()).toBe(true)
  })
})


