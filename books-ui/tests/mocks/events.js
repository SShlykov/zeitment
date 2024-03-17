const apiEvent = {
  "id": "3651c693-39b2-4be9-934d-6c926ec0023d",
  "created_at": "2024-03-03T23:39:52.61276+03:00",
  "updated_at": "2024-03-03T23:41:29.49876+03:00",
  "book_id": "fb5e7d1d-38cd-4831-bae9-07b36080e3e7",
  "chapter_id": "af3ff4ad-bc7d-4e64-acf8-bbd874d4516b",
  "page_id": "2b15d86b-e52c-4d6f-9629-0bf3bc940f29",
  "paragraph_id": "12b9b045-0845-462c-b372-0fca3180a6af",
  "event_type": "test",
  "is_public": true,
  "key": "test",
  "value": "123",
  "link": "https://images.unsplash.com/photo-1703717101037-132d2c3fdd03?q=80&w=2070&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
  "link_text": null,
  "link_type": null,
  "link_image": null,
  "description": null
}

const appEvent = {
  "id": "3651c693-39b2-4be9-934d-6c926ec0023d",
  "createdAt": "2024-03-03T23:39:52.61276+03:00",
  "updatedAt": "2024-03-03T23:41:29.49876+03:00",
  "bookId": "fb5e7d1d-38cd-4831-bae9-07b36080e3e7",
  "chapterId": "af3ff4ad-bc7d-4e64-acf8-bbd874d4516b",
  "pageId": "2b15d86b-e52c-4d6f-9629-0bf3bc940f29",
  "paragraphId": "12b9b045-0845-462c-b372-0fca3180a6af",
  "eventType": "test",
  "isPublic": true,
  "key": "test",
  "value": "123",
  "link": "https://images.unsplash.com/photo-1703717101037-132d2c3fdd03?q=80&w=2070&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
  "linkText": null,
  "linkType": null,
  "linkImage": null,
  "description": null
}

const apiEventsResponse = {
  data: [apiEvent],
  status: "ok"
}

const apiEventResponse = {
  data: apiEvent,
  status: "ok"
}

export { apiEvent, appEvent, apiEventsResponse, apiEventResponse }