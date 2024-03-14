import {expect, test, describe, beforeEach} from 'vitest'
import {createStore} from "vuex";
import { store as layout } from '@/store/modules/layout';
import ServiceOfLayout from '@services/ServiceOfLayout.js'

describe('serviceOfBooks', () => {
  const store = createStore({
    modules: {
      layout
    }
  })

  beforeEach(async () => {
    await store.dispatch('layout/resetStore')
  })

  const serviceOfLayout = new ServiceOfLayout(store);

  beforeEach(async () => {
    await store.dispatch('books/resetStore')
  })

  test('serviceOfBooks is exist', () => {
    expect(serviceOfLayout).toBeDefined()
  })

  test("add notification", async () => {
    const notification = {
      message: "test",
      type: "success"
    }

    await serviceOfLayout.addNotification(notification)
    const notifications = store.getters['layout/notifications']
    const [notificationFromStore] = notifications
    delete notificationFromStore.id
    expect(notificationFromStore).toEqual(notification)
  })
})