import {test, describe, vi, expect} from 'vitest'
import axios from 'axios'
import BooksApi from "@adapters/AdapterOfBooks.js"
import {appBook, apiBookResponse, apiBooksResponse} from '@mocks/books.js'

vi.mock('axios')

describe("test BooksApi class with mocks data and adapterConfig", () => {
  const url = "http://localhost:8000/api/v1/books/"

  const BooksService = new BooksApi(url)

  test("get books list", async () => {
    axios.get.mockResolvedValue({data: apiBooksResponse})

    const booksData = await BooksService.getBooks()
    expect(booksData).toEqual([appBook])
  })

  test("create book (save book)", async () => {
    axios.post.mockResolvedValue({data: apiBookResponse})

    const booksData = await BooksService.createBook(appBook)
    expect(booksData).toEqual(appBook)
  })

  test("update book", async () => {
    axios.put.mockResolvedValue({data: apiBookResponse})

    const booksData = await BooksService.updateBook(appBook)
    expect(booksData).toEqual(appBook)
  })

  test("get book by id", async () => {
    axios.get.mockResolvedValue({data: apiBookResponse})

    const booksData = await BooksService.getBookById(1)
    expect(booksData).toEqual(appBook)

  })

  test("delete book", async () => {
    axios.delete.mockResolvedValue({data: apiBookResponse})

    const booksData = await BooksService.deleteBookById(appBook.id)
    expect(booksData).toEqual(appBook)
  })
})

