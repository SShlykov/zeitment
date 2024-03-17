import { expect, test, describe, vi, beforeEach} from "vitest";
import { createStore } from 'vuex';
import { store as pages } from '@/store/modules/pages';
import { appPage } from "@mocks/pages.js";

describe("tests pages store with vuex", () => {
  const store = createStore({
    plugins: [],
    modules: {
      pages
    },
  })

  beforeEach(() => {
    store.dispatch('pages/resetStore')
  })

  test('test of saveChapterPages', async () => {
    await store.dispatch('pages/saveChapterPages', [appPage])
    const pagesList = store.getters['pages/chapterPages']
    expect(pagesList).toEqual([appPage])
  })

  test('select editable page', async () => {
    await store.dispatch('pages/saveEditablePage', appPage)
    const editablePage = store.getters['pages/editablePage']
    expect(editablePage).toEqual(appPage)
  })
})