import {test, describe, expect} from 'vitest'
import { mount} from '@vue/test-utils.proto'
import InteractivePopup from '@molecules/InteractivePopup.vue';

describe("tests of InteractivePopup", () => {
  test('mount test of InteractivePopup', async () => {

    const wrapper = mount(InteractivePopup, {
      shallow: true,
    })

    expect(wrapper.exists()).toBe(true)
  })
})

