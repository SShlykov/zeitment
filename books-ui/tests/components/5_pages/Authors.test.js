import {test, describe, expect} from 'vitest'
import { mount} from '@vue/test-utils.proto'
import AuthorsPage from '@pages/Authors/Authors.vue';

describe("tests of AuthorsPage", () => {
  test('mount test of AuthorsPage', async () => {

    const wrapper = mount(AuthorsPage, {
      shallow: true,
    })

    expect(wrapper.exists()).toBe(true)
  })
})


