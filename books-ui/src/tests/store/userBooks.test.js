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
        "id": "fb5e7d1d-38cd-4831-bae9-07b36080e3e7",
        "created_at": "2024-03-01T23:47:35.711668+03:00",
        "updated_at": "2024-03-01T23:47:35.711668+03:00",
        "deleted_at": null,
        "owner": "e75aae0d-c1eb-4199-a1d8-2177f57d6a1e",
        "title": "Тестовая книга",
        "author": "Васильев А.В.",
        "description": "test description",
        "is_public": false,
        "publication": null,
        "image_link": null,
        "map_link": null,
        "map_params_id": null,
        "variables": []
      }
    ]

    axios.get.mockResolvedValue({data: booksMock})

    await store.dispatch('userBooks/fetchUserBooks')
    expect(store.state.userBooks.booksList).toEqual(booksMock)
  })
})


