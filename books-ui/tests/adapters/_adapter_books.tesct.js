import {test, describe, vi, expect} from 'vitest'
import axios from 'axios'
import BooksApi, {adapterToApiFromParams} from "@adapters/AdapterOfBooks.js"
import {reverseObject} from '@helpers/objectUtils'
import {appBook, apiBook, apiBookResponse, apiBooksResponse} from '@mocks/books.js'

vi.mock('axios')

describe("test BooksApi class with mocks data and adapterConfig", () => {

  const bookMock = {
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

  const expectedBook = {
    "id": "fb5e7d1d-38cd-4831-bae9-07b36080e3e7",
    "createdAt": "2024-03-01T23:47:35.711668+03:00",
    "updatedAt": "2024-03-01T23:47:35.711668+03:00",
    "deletedAt": null,
    "owner": "e75aae0d-c1eb-4199-a1d8-2177f57d6a1e",
    "title": "Тестовая книга",
    "author": "Васильев А.В.",
    "description": "test description",
    "isPublic": false,
    "publication": null,
    "imageLink": null,
    "mapLink": null,
    "mapParamsId": null,
    "variables": []
  }

  const url = "http://localhost:8000/api/v1/books/"

  const BooksService = new BooksApi(url)

  test("get books list", async () => {
    axios.get.mockResolvedValue({data: apiBooksResponse})

    const booksData = await BooksService.getBooks()
    expect(booksData).toEqual([appBook])
  })

  test("create book (save book)", async () => {
    axios.post.mockResolvedValue({data: apiBookResponse})

    const booksData = await BooksService.createBook(expectedBook)
    expect(booksData).toEqual(appBook)
  })

  test("update book", async () => {
    axios.put.mockResolvedValue({data: apiBookResponse})

    const booksData = await BooksService.updateBook(expectedBook)
    expect(booksData).toEqual(expectedBook)
  })

  test("get book by id", async () => {
    axios.get.mockResolvedValue({data: apiBookResponse})

    const booksData = await BooksService.getBookById(1)
    expect(booksData).toEqual(expectedBook)

  })

  test("delete book", async () => {
    axios.delete.mockResolvedValue({data: apiBookResponse})

    const booksData = await BooksService.deleteBookById(bookMock.id)
    expect(booksData).toEqual(expectedBook)
  })
})


// describe("test BooksApi class with mocks data and without adapterConfig", () => {
//
//   const bookMock = {
//     "id": "fb5e7d1d-38cd-4831-bae9-07b36080e3e7",
//     "created_at": "2024-03-01T23:47:35.711668+03:00",
//     "updated_at": "2024-03-01T23:47:35.711668+03:00",
//     "deleted_at": null,
//     "owner": "e75aae0d-c1eb-4199-a1d8-2177f57d6a1e",
//     "title": "Тестовая книга",
//     "author": "Васильев А.В.",
//     "description": "test description",
//     "is_public": false,
//     "publication": null,
//     "image_link": null,
//     "map_link": null,
//     "map_params_id": null,
//     "variables": []
//   }
//
//   const bookConfig = null
//
//   const expectedBook = bookMock
//
//   const BooksService = new BooksApi(bookConfig)
//
//   test("get books list", async () => {
//     const booksMock = [
//       bookMock
//     ]
//
//     axios.get.mockResolvedValue({data: booksMock})
//
//     const booksData = await BooksService.getBooks()
//     expect(booksData).toEqual([expectedBook])
//   })
//
//   test("create book (save book)", async () => {
//     axios.post.mockResolvedValue({data: bookMock})
//
//     const booksData = await BooksService.createBook(expectedBook)
//     expect(booksData).toEqual(expectedBook)
//   })
//
//   test("update book", async () => {
//     axios.put.mockResolvedValue({data: bookMock})
//
//     const booksData = await BooksService.updateBook(expectedBook)
//     expect(booksData).toEqual(expectedBook)
//   })
//
//   test("get book by id", async () => {
//     axios.get.mockResolvedValue({data: bookMock})
//
//     const booksData = await BooksService.getBookById(1)
//     expect(booksData).toEqual(expectedBook)
//
//   })
//
//   test("delete book", async () => {
//     axios.delete.mockResolvedValue({data: bookMock})
//
//     const booksData = await BooksService.deleteBookById(bookMock.id)
//     expect(booksData).toEqual(expectedBook)
//   })
// })
//
// describe("test adapterToApiFromParams", () => {
//   test("adapterToApiFromParams", () => {
//     const params = null
//     const adapterConfig = {
//       "id": "id",
//       "created_at": "createdAt",
//       "updated_at": "updatedAt",
//       "deleted_at": "deletedAt",
//       "owner": "owner",
//       "title": "title",
//       "author": "author",
//       "description": "description",
//       "is_public": "isPublic",
//       "publication": "publication",
//       "image_link": "imageLink",
//       "map_link": "mapLink",
//       "map_params_id": "mapParamsId",
//       "variables": "variables"
//     }
//
//     const config = adapterToApiFromParams(adapterConfig, params)
//     expect(config).toEqual(reverseObject(adapterConfig))
//   })
// })