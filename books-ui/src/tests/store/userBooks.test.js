import { expect, test, describe, vi, beforeEach} from "vitest";
import { createStore } from 'vuex';
import {store as userBooks} from "@/store/modules/userBooks";
import axios from "axios";

vi.mock('axios')

describe("tests books store with vuex", () => {
  const store = createStore({
    plugins: [],
    modules: {
      userBooks
    },
  })

  beforeEach(() => {
    store.dispatch('userBooks/resetStore')
  })

  test('fetch books list', async () => {
    const booksMock = [
      {
        id: 1,
        title: "book1",
      },
      {
        id: 2,
        title: "book2",
      }
    ]

    axios.get.mockResolvedValue({data: booksMock})

    await store.dispatch('userBooks/fetchBooks')
    expect(store.state.userBooks.booksList).toEqual(booksMock)
  })
})


