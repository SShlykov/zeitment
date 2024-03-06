import {test, describe, expect} from 'vitest'
import { mount} from '@vue/test-utils'
import BackgroundImage from '@atoms/BackgroundImage.vue';

describe("tests of BackgroundImage", () => {
  test('mount test of BackgroundImage', async () => {

    const wrapper = mount(BackgroundImage, {
      shallow: true,
    })

    expect(wrapper.exists()).toBe(true)
  })
})


