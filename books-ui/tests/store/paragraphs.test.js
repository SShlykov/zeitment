import { expect, test, describe, beforeEach} from "vitest";
import { createStore } from 'vuex';
import { store as paragraphs } from '@/store/modules/paragraphs';
import { appParagraph } from "@mocks/paragraphs.js";

describe("tests chapters store with vuex", () => {
  const store = createStore({
    plugins: [],
    modules: {
      paragraphs
    },
  })

  beforeEach(() => {
    store.dispatch('paragraphs/resetStore')
  })

  test('test of saveParagraphs', async () => {
    await store.dispatch('paragraphs/saveParagraphs', [appParagraph])
    const paragraphsList = store.getters['paragraphs/paragraphs']
    expect(paragraphsList).toEqual([appParagraph])
  })
})