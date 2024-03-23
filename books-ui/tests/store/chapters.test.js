import { expect, test, describe, beforeEach} from "vitest";
import { createStore } from 'vuex';
import { store as chapters } from '@/store/modules/chapters';
import { appChapter } from "@mocks/chapters.js";

describe("tests chapters store with vuex", () => {
  const store = createStore({
    plugins: [],
    modules: {
      chapters
    },
  })

  beforeEach(() => {
    store.dispatch('chapters/resetStore')
  })

  test('test of saveChapters', async () => {
    await store.dispatch('chapters/saveChapters', [appChapter])
    const chaptersList = store.getters['chapters/chapters']
    expect(chaptersList).toEqual([appChapter])
  })

  test('select current chapter', async () => {
    await store.dispatch('chapters/saveCurrentChapter', appChapter)
    const currentChapter = store.getters['chapters/currentChapter']
    expect(currentChapter).toEqual(appChapter)
  })
})