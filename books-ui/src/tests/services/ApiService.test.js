import BooksApi from "@apiServices/BooksApi.js"
import {test, describe, vi, expect} from 'vitest'
import axios from 'axios'

vi.mock('axios')

describe("test BooksApi class with mocks data", () => {
  const BooksService = new BooksApi()
  console.log(BooksService)

  test("get books list", async () => {
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

    const booksData = await BooksService.getBooks()
    expect(booksData).toEqual(booksMock)
  })

  test("create book (save book)", async () => {
    const bookMock = {
      "title": "Тестовая книга",
      "author": "Васильев А.В.",
      "owner": 1,
      "description": "test description",
      "is_public": false,
      "publication": null
    }

    axios.post.mockResolvedValue({data: bookMock})

    const booksData = await BooksService.createBook(bookMock)
    expect(booksData).toEqual(bookMock)
  })

  test("update book", async () => {
    const bookMock = {
      "id": 1,
      "title": "Тестовая книга",
      "author": "Васильев А.В.",
      "owner": 1,
      "description": "test description",
      "is_public": false,
      "publication": null
    }

    axios.get.mockResolvedValue({data: bookMock})

    const booksData = await BooksService.getBookById(bookMock)
    expect(booksData).toEqual(bookMock)
  })

  test("get book by id", async () => {
    const bookMock = {
      "id": 1,
      "title": "Тестовая книга",
      "author": "Васильев А.В.",
      "owner": 1,
      "description": "test description",
      "is_public": false,
      "publication": null
    }

    axios.get.mockResolvedValue({data: bookMock})

    const booksData = await BooksService.getBookById(1)
    expect(booksData).toEqual(bookMock)

  })

  test("delete book", async () => {
    const bookMock = {
      "id": 1,
      "deleted_at": "2024-02-15T17:24:52.755254148+03:00",
      "title": "Тестовая книга",
      "author": "Васильев А.В.",
      "owner": 1,
      "description": "test description",
      "is_public": false,
      "publication": null
    }

    axios.delete.mockResolvedValue({data: bookMock})

    const booksData = await BooksService.deleteBookById(1)
    expect(booksData).toEqual(bookMock)
  })
})