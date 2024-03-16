import {test, describe, expect} from 'vitest'
import { mount} from '@vue/test-utils.proto'
import Landing from '@pages/Landing/Landing.vue';

describe("tests of Landing", () => {
  test('mount test of Landing', async () => {

    const wrapper = mount(Landing, {
      shallow: true,
    })

    expect(wrapper.exists()).toBe(true)
  })
})


