import {expect, test, describe, beforeEach, expectTypeOf} from 'vitest'
import {createStore} from "vuex";
import { store as chapters } from '@/store/modules/chapters';
import ServiceOfChapters from '@services/ServiceOfChapters.js'
import {AdapterOfChapters, appChapter} from "@mocks/chapters"

describe('serviceOfChapters', () => {
  const store = createStore({
    modules: {
      chapters
    }
  })

  beforeEach(async () => {
    await store.dispatch('chapters/resetStore')
  })

  const adapterOfChapters = new AdapterOfChapters('');
  const serviceOfChapters = new ServiceOfChapters(adapterOfChapters, store);

  test('serviceOfChapters is exist', () => {
    expect(serviceOfChapters).toBeDefined()
  })

  test('fetch chaptersList by book id', async () => {
    await serviceOfChapters.fetchChaptersByBookId(1)
    const chaptersList = store.getters['chapters/chapters']
    expect(chaptersList).toEqual([appChapter])
  })
})