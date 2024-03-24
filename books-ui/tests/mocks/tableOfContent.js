const apiTableOfContent = {
  "book_id": "fb5e7d1d-38cd-4831-bae9-07b36080e3e7",
  "book_title": "Тестова",
  "author": "e75aae0d-c1eb-4199-a1d8-2177f57d6a1e",
  "authorship": "",
  "tags": [],
  "sections": [
    {
      "id": "af3ff4ad-bc7d-4e64-acf8-bbd874d4516b",
      "title": "Chapter 1",
      "order": 1000,
      "level": "chapter",
      "is_public": false
    },
    {
      "id": "2b15d86b-e52c-4d6f-9629-0bf3bc940f29",
      "title": "Page 2",
      "order": 1001,
      "level": "page",
      "is_public": false
    },
    {
      "id": "99dfcc38-41d3-4967-bd6f-df22ad551cac",
      "title": "Chapter 2",
      "order": 2000,
      "level": "chapter",
      "is_public": true
    },
    {
      "id": "22943bd1-bc4e-44fa-b398-5190943d3611",
      "title": "Chapter 2",
      "order": 3000,
      "level": "chapter",
      "is_public": false
    }
  ]
}

const appTableOfContent = {
  "bookId": "fb5e7d1d-38cd-4831-bae9-07b36080e3e7",
  "bookTitle": "Тестова",
  "author": "e75aae0d-c1eb-4199-a1d8-2177f57d6a1e",
  "authorship": "",
  "tags": [],
  "sections": [
    {
      "id": "af3ff4ad-bc7d-4e64-acf8-bbd874d4516b",
      "title": "Chapter 1",
      "order": 1000,
      "level": "chapter",
      "isPublic": false
    },
    {
      "id": "2b15d86b-e52c-4d6f-9629-0bf3bc940f29",
      "title": "Page 2",
      "order": 1001,
      "level": "page",
      "isPublic": false
    },
    {
      "id": "99dfcc38-41d3-4967-bd6f-df22ad551cac",
      "title": "Chapter 2",
      "order": 2000,
      "level": "chapter",
      "isPublic": true
    },
    {
      "id": "22943bd1-bc4e-44fa-b398-5190943d3611",
      "title": "Chapter 2",
      "order": 3000,
      "level": "chapter",
      "isPublic": false
    }
  ]
}

const apiTableOfContentResponse = {
  data: apiTableOfContent,
  status: "ok"
}

const apiTableOfContents = apiTableOfContent
const appTableOfContents = appTableOfContent
const apiTableOfContentsResponse = apiTableOfContentResponse

export {appTableOfContent, apiTableOfContentResponse, apiTableOfContent, apiTableOfContentsResponse, apiTableOfContents, appTableOfContents}

