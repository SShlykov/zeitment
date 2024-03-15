import {test, describe, vi, expect} from 'vitest'
import axios from 'axios'
import AdapterOfParagraphs from "@adapters/AdapterOfParagraphs.js"
import {appParagraph, apiParagraphResponse, apiParagraphsResponse} from '@mocks/paragraphs.js'

vi.mock('axios')

describe("tests AdapterOfParagraphs for getParagraphsByPageId", () => {
  const url = "http://localhost:8000/api/v1/paragraphs/"
  axios.post.mockResolvedValue({data: apiParagraphsResponse})
  const paragraphsAdapter = new AdapterOfParagraphs(url)


  test("get paragraphs list by page id", async () => {
    const paragraphsData = await paragraphsAdapter.getParagraphsByPageId(appParagraph.id)
    expect(paragraphsData).toEqual([appParagraph])
  })

  test("get paragraphs list by page id with bad param", async () => {
    let paragraphsData = await paragraphsAdapter.getParagraphsByPageId(appParagraph.id, {
      foo: "bar"
    })
    expect(paragraphsData).toEqual([appParagraph])
  })

  test("get paragraphs list by page id with bad paragraph= -1", async () => {
    let paragraphsData = await paragraphsAdapter.getParagraphsByPageId(appParagraph.id, {
      page: -1,
    })
    expect(paragraphsData).toEqual([])
  })

  test("get paragraphs list by page id with bad page_size= -1", async () => {
    let paragraphsData = await paragraphsAdapter.getParagraphsByPageId(appParagraph.id, {
      page_size: -1,
    })
    expect(paragraphsData).toEqual([])
  })

  test("get paragraphs list by page id with bad page is string", async () => {
    let paragraphsData = await paragraphsAdapter.getParagraphsByPageId(appParagraph.id, {
      page: "foo",
    })
    expect(paragraphsData).toEqual([])
  })

  test("get paragraphs list by page id with bad page is null", async () => {
    let paragraphsData = await paragraphsAdapter.getParagraphsByPageId(appParagraph.id, {
      page: null,
    })
    expect(paragraphsData).toEqual([appParagraph])
  })
})

describe("tests AdapterOfParagraphs for get paragraphs", () => {
  const url = "http://localhost:8000/api/v1/paragraphs/"
  axios.post.mockResolvedValue({data: apiParagraphsResponse})
  const paragraphsAdapter = new AdapterOfParagraphs(url)

  test("get paragraphs list", async () => {
    const paragraphsData = await paragraphsAdapter.getParagraphs()
    expect(paragraphsData).toEqual([appParagraph])
  })

  test("get paragraphs list with bad param", async () => {
    let paragraphsData = await paragraphsAdapter.getParagraphs({
      foo: "bar"
    })
    expect(paragraphsData).toEqual([appParagraph])
  })


  test("get paragraphs list with bad page= -1", async () => {
    let paragraphsData = await paragraphsAdapter.getParagraphs({
      page: -1,
    })
    expect(paragraphsData).toEqual([])
  })

  test("get paragraphs list with bad page_size= -1", async () => {
    let paragraphsData = await paragraphsAdapter.getParagraphs({
      page_size: -1,
    })
    expect(paragraphsData).toEqual([])
  })

  test("get paragraphs list with bad page is string", async () => {
    let paragraphsData = await paragraphsAdapter.getParagraphs({
      page: "foo",
    })
    expect(paragraphsData).toEqual([])
  })

  test("get paragraphs list with bad page is null", async () => {
    let paragraphsData = await paragraphsAdapter.getParagraphs({
      page: null,
    })
    expect(paragraphsData).toEqual([appParagraph])
  })
})

describe("tests of AdapterOfParagraphs ", () => {
  const url = "http://localhost:8000/api/v1/paragraphs/"
  const paragraphsAdapter = new AdapterOfParagraphs(url)

  test("get paragraphs list", async () => {
    axios.post.mockResolvedValue({data: apiParagraphsResponse})

    const paragraphsData = await paragraphsAdapter.getParagraphs()
    expect(paragraphsData).toEqual([appParagraph])
  })

  test("create paragraph (save paragraph)", async () => {
    axios.post.mockResolvedValue({data: apiParagraphResponse})

    const paragraphsData = await paragraphsAdapter.createParagraph(appParagraph)
    expect(paragraphsData).toEqual(appParagraph)
  })

  test("update paragraph", async () => {
    axios.put.mockResolvedValue({data: apiParagraphResponse})

    const paragraphsData = await paragraphsAdapter.updateParagraph(appParagraph)
    expect(appParagraph.id).toBeDefined()
    expect(paragraphsData).toEqual(appParagraph)
  })

  test("get paragraph by id", async () => {
    axios.get.mockResolvedValue({data: apiParagraphResponse})

    const paragraphsData = await paragraphsAdapter.getParagraphById(1)
    expect(paragraphsData).toEqual(appParagraph)

  })

  test("delete paragraph", async () => {
    axios.delete.mockResolvedValue({data: apiParagraphResponse})

    const paragraphsData = await paragraphsAdapter.deleteParagraphById(appParagraph.id)
    expect(paragraphsData).toEqual(appParagraph)
  })
})

