import { expect, test, describe, vi, beforeEach} from "vitest";
import { createStore } from 'vuex';
import { store as books } from '@/store/modules/books';
import { groupSectionsByPagesAndChapters } from '@/store/modules/books/helpers.js';
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


describe("book store helpers", () => {
  test("group tableOfContents sections by chapters and pages", () => {
    const sections = appTableOfContent.sections
    const groupedSections = groupSectionsByPagesAndChapters(sections)
    const expectedList = [
      {
        "id": "af3ff4ad-bc7d-4e64-acf8-bbd874d4516b",
        "title": "Chapter 1",
        "order": 1000,
        "level": "chapter",
        "isPublic": false,
        items: [
          {
            "id": "2b15d86b-e52c-4d6f-9629-0bf3bc940f29",
            "title": "Page 2",
            "order": 1001,
            "level": "page",
            "isPublic": false
          }
        ]
      },
      {
        "id": "99dfcc38-41d3-4967-bd6f-df22ad551cac",
        "title": "Chapter 2",
        "order": 2000,
        "level": "chapter",
        "isPublic": true,
        items: []
      },
      {
        "id": "22943bd1-bc4e-44fa-b398-5190943d3611",
        "title": "Chapter 2",
        "order": 3000,
        "level": "chapter",
        "isPublic": false,
        items: []
      }
    ]
    expect(groupedSections).toEqual(expectedList)
  })

  test("group tableOfContents sections by chapters and  with empty list", () => {
    const groupedSections = groupSectionsByPagesAndChapters([])
    const expectedList = []
    expect(groupedSections).toEqual(expectedList)
  })
})

