import { expect, test, describe, vi, beforeEach} from "vitest";
import { createStore } from 'vuex';
import { store as books } from '@/store/modules/books';
import { appBook } from "@mocks/books.js";
import { appTableOfContent } from "@mocks/tableOfContent.js";

describe("tests books store with vuex", () => {
  const store = createStore({
    plugins: [],
    modules: {
      books
    },
  })

  beforeEach(() => {
    store.dispatch('books/resetStore')
  })

  test('test of saveUserBooks', async () => {
    await store.dispatch('books/saveUserBooks', [appBook])
    const booksList = store.getters['books/userBooks']
    expect(booksList).toEqual([appBook])
  })

  test('select current book', async () => {
    await store.dispatch('books/saveCurrentBook', appBook)
    const editableBook = store.getters['books/currentBook']
    expect(editableBook).toEqual(appBook)
  })

  test('test of saveTableOfContent', async () => {
    await store.dispatch('books/saveTableOfContent', appTableOfContent)
    const tableOfContents = store.getters['books/tableOfContents']
    expect(tableOfContents).toEqual(appTableOfContent)
  })
})


