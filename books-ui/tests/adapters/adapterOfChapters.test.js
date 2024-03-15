import {test, describe, vi, expect} from 'vitest'
import axios from 'axios'
import AdapterOfChapters from "@adapters/AdapterOfChapters.js"
import {appChapter, apiChaptersResponse, apiChapterResponse} from '@mocks/chapters.js'

vi.mock('axios')

describe("tests AdapterOfChapters for getChaptersByBookId", () => {
  const url = "http://localhost:8000/api/v1/chapters/"
  axios.post.mockResolvedValue({data: apiChaptersResponse})
  const chaptersAdapter = new AdapterOfChapters(url)


  test("get chapters list by book id", async () => {
    const chaptersData = await chaptersAdapter.getChaptersByBookId(appChapter.id)
    expect(chaptersData).toEqual([appChapter])
  })

  test("get chapters list by book id with bad param", async () => {
    let chaptersData = await chaptersAdapter.getChaptersByBookId(appChapter.id, {
      foo: "bar"
    })
    expect(chaptersData).toEqual([appChapter])
  })

  test("get chapters list by book id with bad page= -1", async () => {
    let chaptersData = await chaptersAdapter.getChaptersByBookId(appChapter.id, {
      page: -1,
    })
    expect(chaptersData).toEqual([])
  })

  test("get chapters list by book id with bad page_size= -1", async () => {
    let chaptersData = await chaptersAdapter.getChaptersByBookId(appChapter.id, {
      page_size: -1,
    })
    expect(chaptersData).toEqual([])
  })

  test("get chapters list by book id with bad page is string", async () => {
    let chaptersData = await chaptersAdapter.getChaptersByBookId(appChapter.id, {
      page: "foo",
    })
    expect(chaptersData).toEqual([])
  })

  test("get chapters list by book id with bad page is null", async () => {
    let chaptersData = await chaptersAdapter.getChaptersByBookId(appChapter.id, {
      page: null,
    })
    expect(chaptersData).toEqual([appChapter])
  })
})

describe("tests AdapterOfChapters for get chapters", () => {
  const url = "http://localhost:8000/api/v1/chapters/"
  axios.post.mockResolvedValue({data: apiChaptersResponse})
  const chaptersAdapter = new AdapterOfChapters(url)

  test("get chapters list", async () => {
    const chaptersData = await chaptersAdapter.getChapters()
    expect(chaptersData).toEqual([appChapter])
  })

  test("get chapters list with bad param", async () => {
    let chaptersData = await chaptersAdapter.getChapters({
      foo: "bar"
    })
    expect(chaptersData).toEqual([appChapter])
  })


  test("get chapters list with bad page= -1", async () => {
    let chaptersData = await chaptersAdapter.getChapters({
      page: -1,
    })
    expect(chaptersData).toEqual([])
  })

  test("get chapters list with bad page_size= -1", async () => {
    let chaptersData = await chaptersAdapter.getChapters({
      page_size: -1,
    })
    expect(chaptersData).toEqual([])
  })

  test("get chapters list with bad page is string", async () => {
    let chaptersData = await chaptersAdapter.getChapters({
      page: "foo",
    })
    expect(chaptersData).toEqual([])
  })

  test("get chapters list with bad page is null", async () => {
    let chaptersData = await chaptersAdapter.getChapters({
      page: null,
    })
    expect(chaptersData).toEqual([appChapter])
  })
})

describe("tests of AdapterOfChapters ", () => {
  const url = "http://localhost:8000/api/v1/chapters/"
  const chaptersAdapter = new AdapterOfChapters(url)

  test("get chapters list", async () => {
    axios.post.mockResolvedValue({data: apiChaptersResponse})

    const chaptersData = await chaptersAdapter.getChapters()
    expect(chaptersData).toEqual([appChapter])
  })

  test("create chapter (save chapter)", async () => {
    axios.post.mockResolvedValue({data: apiChapterResponse})

    const chaptersData = await chaptersAdapter.createChapter(appChapter)
    expect(chaptersData).toEqual(appChapter)
  })

  test("update chapter", async () => {
    axios.put.mockResolvedValue({data: apiChapterResponse})

    const chaptersData = await chaptersAdapter.updateChapter(appChapter)
    expect(appChapter.id).toBeDefined()
    expect(chaptersData).toEqual(appChapter)
  })

  test("get chapter by id", async () => {
    axios.get.mockResolvedValue({data: apiChapterResponse})

    const chaptersData = await chaptersAdapter.getChapterById(1)
    expect(chaptersData).toEqual(appChapter)

  })

  test("delete chapter", async () => {
    axios.delete.mockResolvedValue({data: apiChapterResponse})

    const chaptersData = await chaptersAdapter.deleteChapterById(appChapter.id)
    expect(chaptersData).toEqual(appChapter)
  })
})

