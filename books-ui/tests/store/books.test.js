import { expect, test, describe, vi, beforeEach} from "vitest";
import { createStore } from 'vuex';
import { store as books } from '@/store/modules/books';
import { appBook } from "@mocks/books.js";

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

  test('select editable book', async () => {
    await store.dispatch('books/saveEditableBook', appBook)
    const editableBook = store.getters['books/editableBook']
    expect(editableBook).toEqual(appBook)
  })
})


