import { expect, test, describe, beforeEach} from "vitest";
import { createStore } from 'vuex';
import { store as layout } from '@/store/modules/layout';

describe("tests layout store with vuex", () => {
  const store = createStore({
    plugins: [],
    modules: {
      layout
    },
  })

  beforeEach(() => {
    store.dispatch('layout/resetStore')
  })

  test('test of setNotifications', async () => {
    await store.dispatch('layout/setNotifications', [{id: 1, message: "test"}])
    const notifications = store.getters['layout/notifications']
    expect(notifications).toEqual([{id: 1, message: "test"}])
  })

  test('test of addNotification', async () => {
    await store.dispatch('layout/addNotification', {id: 1, message: "test"})
    await store.dispatch('layout/addNotification', {id: 2, message: "test2"})
    const notifications = store.getters['layout/notifications']
    expect(notifications).toEqual([{id: 1, message: "test"}, {id: 2, message: "test2"}])
  })

  test('test of removeNotification', async () => {
    await store.dispatch('layout/addNotification', {id: 1, message: "test"})
    await store.dispatch('layout/addNotification', {id: 2, message: "test2"})
    await store.dispatch('layout/removeNotification', 1)
    const notifications = store.getters['layout/notifications']
    expect(notifications).toEqual([{id: 2, message: "test2"}])
  })
})