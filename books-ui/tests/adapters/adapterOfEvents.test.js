import {test, describe, vi, expect} from 'vitest'
import AdapterOfEvents from "@adapters/AdapterOfEvents.js"
import axios from 'axios'
import {apiEventsResponse, apiEventResponse, appEvent} from '@mocks/events.js'
import {appPage} from '@mocks/pages.js'
import {appChapter} from '@mocks/chapters.js'
import {appBook} from '@mocks/books.js'
import {appParagraph} from '@mocks/paragraphs.js'

vi.mock('axios')

describe("tests of AdapterOfEvents ", () => {
  const url = "http://localhost:8000/api/v1/events/"
  const eventsAdapter = new AdapterOfEvents(url)

  test("get events list by book id", async () => {
    axios.post.mockResolvedValue({data: apiEventsResponse})
    const eventsData = await eventsAdapter.getEventsByBookId(appBook.id)
    expect(eventsData).toEqual([appEvent])
  })

  test("get events list by chapter id", async () => {
    axios.post.mockResolvedValue({data: apiEventsResponse})
    const eventsData = await eventsAdapter.getEventsByChapterId(appChapter.id)
    expect(eventsData).toEqual([appEvent])
  })

  test("get events list by page id", async () => {
    axios.post.mockResolvedValue({data: apiEventsResponse})
    const eventsData = await eventsAdapter.getEventsByPageId(appPage.id)
    expect(eventsData).toEqual([appEvent])
  })

  test("get events list by paragraph id", async () => {
    axios.post.mockResolvedValue({data: apiEventsResponse})
    const eventsData = await eventsAdapter.getEventsByParagraphId(appParagraph.id)
    expect(eventsData).toEqual([appEvent])
  })

  test("get event by id", async () => {
    axios.get.mockResolvedValue({data: apiEventResponse})
    const eventData = await eventsAdapter.getEventById(appEvent.id)
    expect(eventData).toEqual(appEvent)
  })

  test("create event", async () => {
    axios.post.mockResolvedValue({data: apiEventResponse})
    const eventData = await eventsAdapter.createEvent(appEvent)
    expect(eventData).toEqual(appEvent)
  })

  test("update event", async () => {
    axios.put.mockResolvedValue({data: apiEventResponse})
    const eventData = await eventsAdapter.updateEvent(appEvent)
    expect(eventData).toEqual(appEvent)
  })

  test("delete event", async () => {
    axios.delete.mockResolvedValue({data: apiEventResponse})
    const eventData = await eventsAdapter.deleteEventById(appEvent.id)
    expect(eventData).toEqual(appEvent)
  })
})