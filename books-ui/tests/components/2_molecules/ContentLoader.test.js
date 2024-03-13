import {test, describe, expect} from 'vitest'
import { mount} from '@vue/test-utils'
import ContentLoader from '@molecules/VCardsContainer.vue';

describe("tests of ContentLoader", () => {
  test('mount test of ContentLoader', async () => {

    const wrapper = mount(ContentLoader, {
      shallow: true,
    })

    expect(wrapper.exists()).toBe(true)
  })
})

