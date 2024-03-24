import {test, describe, vi, expect} from 'vitest'
import axios from 'axios'
import AdapterOfBooks from "@adapters/AdapterOfBooks/AdapterOfBooks.js"
import {appBook, apiBookResponse, apiBooksResponse} from '@mocks/books.js'
import {apiTableOfContentResponse, appTableOfContent} from "@mocks/tableOfContent.js";

vi.mock('axios')

describe("tests AdapterOfBooks for get books", () => {
  const url = "http://localhost:8000/api/v1/books/"
  axios.post.mockResolvedValue({data: apiBooksResponse})
  const BooksService = new AdapterOfBooks(url)

  test("get books list", async () => {
    const booksData = await BooksService.getBooks()
    expect(booksData).toEqual([appBook])
  })

  test("get books list with bad param", async () => {
    let booksData = await BooksService.getBooks({
      foo: "bar"
    })
    expect(booksData).toEqual([appBook])
  })

  test("get books list with bad page= -1", async () => {
    let booksData = await BooksService.getBooks({
      page: -1,
    })
    expect(booksData).toEqual([])
  })

  test("get books list with bad page_size= -1", async () => {
    let booksData = await BooksService.getBooks({
      page_size: -1,
    })
    expect(booksData).toEqual([])
  })

  test("get books list with bad page is string", async () => {
    let booksData = await BooksService.getBooks({
      page: "foo",
    })
    expect(booksData).toEqual([])
  })

  test("get books list with bad page is null", async () => {
    let booksData = await BooksService.getBooks({
      page: null,
    })
    expect(booksData).toEqual([appBook])
  })
})

describe("tests of AdapterOfBooks ", () => {
  const url = "http://localhost:8000/api/v1/books/"

  const booksAdapter = new AdapterOfBooks(url)

  test("get books list", async () => {
    axios.post.mockResolvedValue({data: apiBooksResponse})

    const booksData = await booksAdapter.getBooks(null)
    expect(booksData).toEqual([appBook])
  })

  test("get books table of content", async () => {
    axios.post.mockResolvedValue({data: apiBooksResponse})

    const booksData = await booksAdapter.getBooks(null)
    expect(booksData).toEqual([appBook])
  })

  test("create book (save book)", async () => {
    axios.post.mockResolvedValue({data: apiBookResponse})

    const booksData = await booksAdapter.createBook(appBook)
    expect(booksData).toEqual(appBook)
  })

  test("update book", async () => {
    axios.put.mockResolvedValue({data: apiBookResponse})

    const booksData = await booksAdapter.updateBook(appBook)
    expect(appBook.id).toBeDefined()
    expect(booksData).toEqual(appBook)
  })

  test("get book by id", async () => {
    axios.get.mockResolvedValue({data: apiBookResponse})

    const booksData = await booksAdapter.getBookById(1)
    expect(booksData).toEqual(appBook)

  })

  test("delete book", async () => {
    axios.delete.mockResolvedValue({data: apiBookResponse})

    const booksData = await booksAdapter.deleteBookById(appBook.id)
    expect(booksData).toEqual(appBook)
  })

  test("get book table of content", async () => {
    axios.post.mockResolvedValue({data: apiTableOfContentResponse})

    const TOC = await booksAdapter.getTableOfContent(appBook.id)
    expect(TOC).toEqual(appTableOfContent)
  })
})
