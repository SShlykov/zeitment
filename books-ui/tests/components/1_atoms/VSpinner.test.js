import {test, describe, expect} from 'vitest'
import { mount} from '@vue/test-utils'
import VSpinner from '@atoms/VSpinner.vue';

describe("tests of Spinner", () => {
  test('mount test of VSpinner', async () => {

    const wrapper = mount(VSpinner, {
      shallow: true,
    })

    expect(wrapper.exists()).toBe(true)
  })
})


