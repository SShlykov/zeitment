import { expect, test, describe, vi, beforeEach} from "vitest";
import { createStore } from 'vuex';
import {store as userBooks} from "books-ui/src/store/modules/books";
import {apiBook, apiBooksResponse, appBook} from "@mocks/books.js";

vi.mock('axios')

describe("tests books store with vuex", () => {
  const store = createStore({
    plugins: [],
    modules: {
      userBooks
    },
  })

  beforeEach(() => {
    store.dispatch('books/resetStore')
  })

  test('fetch books list', async () => {

    await store.dispatch('userBooks/saveUserBooks', [appBook])
    expect(store.state.userBooks.booksList).toEqual([appBook])
  })
})


