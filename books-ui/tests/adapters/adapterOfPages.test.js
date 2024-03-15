import {test, describe, vi, expect} from 'vitest'
import axios from 'axios'
import AdapterOfPages from "@adapters/AdapterOfPages.js"
import {appPage, apiPageResponse, apiPagesResponse} from '@mocks/pages.js'

vi.mock('axios')

describe("tests AdapterOfPages for getPagesByPageId", () => {
  const url = "http://localhost:8000/api/v1/pages/"
  axios.post.mockResolvedValue({data: apiPagesResponse})
  const pagesAdapter = new AdapterOfPages(url)


  test("get pages list by book id", async () => {
    const pagesData = await pagesAdapter.getPagesByChapterId(appPage.id)
    expect(pagesData).toEqual([appPage])
  })

  test("get pages list by book id with bad param", async () => {
    let pagesData = await pagesAdapter.getPagesByChapterId(appPage.id, {
      foo: "bar"
    })
    expect(pagesData).toEqual([appPage])
  })

  test("get pages list by book id with bad page= -1", async () => {
    let pagesData = await pagesAdapter.getPagesByChapterId(appPage.id, {
      page: -1,
    })
    expect(pagesData).toEqual([])
  })

  test("get pages list by book id with bad page_size= -1", async () => {
    let pagesData = await pagesAdapter.getPagesByChapterId(appPage.id, {
      page_size: -1,
    })
    expect(pagesData).toEqual([])
  })

  test("get pages list by book id with bad page is string", async () => {
    let pagesData = await pagesAdapter.getPagesByChapterId(appPage.id, {
      page: "foo",
    })
    expect(pagesData).toEqual([])
  })

  test("get pages list by book id with bad page is null", async () => {
    let pagesData = await pagesAdapter.getPagesByChapterId(appPage.id, {
      page: null,
    })
    expect(pagesData).toEqual([appPage])
  })
})

describe("tests AdapterOfPages for get pages", () => {
  const url = "http://localhost:8000/api/v1/pages/"
  axios.post.mockResolvedValue({data: apiPagesResponse})
  const pagesAdapter = new AdapterOfPages(url)

  test("get pages list", async () => {
    const pagesData = await pagesAdapter.getPages()
    expect(pagesData).toEqual([appPage])
  })

  test("get pages list with bad param", async () => {
    let pagesData = await pagesAdapter.getPages({
      foo: "bar"
    })
    expect(pagesData).toEqual([appPage])
  })


  test("get pages list with bad page= -1", async () => {
    let pagesData = await pagesAdapter.getPages({
      page: -1,
    })
    expect(pagesData).toEqual([])
  })

  test("get pages list with bad page_size= -1", async () => {
    let pagesData = await pagesAdapter.getPages({
      page_size: -1,
    })
    expect(pagesData).toEqual([])
  })

  test("get pages list with bad page is string", async () => {
    let pagesData = await pagesAdapter.getPages({
      page: "foo",
    })
    expect(pagesData).toEqual([])
  })

  test("get pages list with bad page is null", async () => {
    let pagesData = await pagesAdapter.getPages({
      page: null,
    })
    expect(pagesData).toEqual([appPage])
  })
})

describe("tests of AdapterOfPages ", () => {
  const url = "http://localhost:8000/api/v1/pages/"
  const pagesAdapter = new AdapterOfPages(url)

  test("get pages list", async () => {
    axios.post.mockResolvedValue({data: apiPagesResponse})

    const pagesData = await pagesAdapter.getPages()
    expect(pagesData).toEqual([appPage])
  })

  test("create page (save page)", async () => {
    axios.post.mockResolvedValue({data: apiPageResponse})

    const pagesData = await pagesAdapter.createPage(appPage)
    expect(pagesData).toEqual(appPage)
  })

  test("update page", async () => {
    axios.put.mockResolvedValue({data: apiPageResponse})

    const pagesData = await pagesAdapter.updatePage(appPage)
    expect(appPage.id).toBeDefined()
    expect(pagesData).toEqual(appPage)
  })

  test("get page by id", async () => {
    axios.get.mockResolvedValue({data: apiPageResponse})

    const pagesData = await pagesAdapter.getPageById(1)
    expect(pagesData).toEqual(appPage)

  })

  test("delete page", async () => {
    axios.delete.mockResolvedValue({data: apiPageResponse})

    const pagesData = await pagesAdapter.deletePageById(appPage.id)
    expect(pagesData).toEqual(appPage)
  })
})

