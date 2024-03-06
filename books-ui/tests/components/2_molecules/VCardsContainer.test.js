import {test, describe, expect} from 'vitest'
import { mount} from '@vue/test-utils'
import VCardsContainer from '@molecules/VCardsContainer.vue';

describe("tests of VCardsContainer", () => {
  test('mount test of VCardsContainer', async () => {

    const wrapper = mount(VCardsContainer, {
      shallow: true,
    })

    expect(wrapper.exists()).toBe(true)
  })
})

