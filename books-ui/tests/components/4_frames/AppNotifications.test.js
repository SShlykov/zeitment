import {test, describe, expect} from 'vitest'
import { mount} from '@vue/test-utils.proto'
import AppNotifications from '@frames/AppNotifications/AppNotifications.vue';
import AppNotification from '@frames/AppNotifications/AppNotificationsItem.vue';

describe("tests of AppNotifications", () => {
  test('mount test of AppNotifications', async () => {
    const wrapper = mount(AppNotifications, {
      shallow: true,
    })
    expect(wrapper.exists()).toBe(true)
  })

  test('render notifications', async () => {
    const wrapper = mount(AppNotifications, {
      props: {
        notifications: [
          {
            message: "test",
            type: "success"
          },
          {
            message: "foo_bar test",
            type: "error"
          },
        ]
      }
    })
    expect(wrapper.text()).contains("test")
    expect(wrapper.text()).contains("foo_bar test")
  })
})

describe("tests of AppNotification", () => {
  test('mount test of AppNotification', async () => {

    const wrapper = mount(AppNotification, {
      shallow: true,
    })

    expect(wrapper.exists()).toBe(true)
  })
})

