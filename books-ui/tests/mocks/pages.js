const apiPage = {
  "id": "63e4e2cd-cf70-4e5d-96b0-06eb01c311f2",
  "created_at": "2024-03-08T23:10:19.037536+03:00",
  "updated_at": "2024-03-08T23:10:35.543244+03:00",
  "deleted_at": null,
  "title": "Page 8",
  "text": "fjhdsjkhfdkjghdslkhfskldjhgfskldjbvsfkdbvksfdab",
  "chapter_id": "af3ff4ad-bc7d-4e64-acf8-bbd874d4516b",
  "is_public": false,
  "map_params_id": null
}

const appPage = {
  "id": "63e4e2cd-cf70-4e5d-96b0-06eb01c311f2",
  "createdAt": "2024-03-08T23:10:19.037536+03:00",
  "updatedAt": "2024-03-08T23:10:35.543244+03:00",
  "deletedAt": null,
  "title": "Page 8",
  "text": "fjhdsjkhfdkjghdslkhfskldjhgfskldjbvsfkdbvksfdab",
  "chapterId": "af3ff4ad-bc7d-4e64-acf8-bbd874d4516b",
  "isPublic": false,
  "mapParamsId": null
}

const apiPagesResponse = {
  data: [apiPage],
  status: "ok"
}

const apiPageResponse = {
  data: apiPage,
  status: "ok"
}

class AdapterOfPages {
  constructor(url) {
    this.url = url
  }

  async fetchChapterPages() {
    return [appPage]
  }

  async createPage() {
    return appPage
  }

  async updatePage(newPage) {
    return {
      ...appPage,
      ...newPage
    }
  }

  async removePage() {
    return appPage
  }

  async getPageById() {
    return appPage
  }

  async getPagesByChapterId() {
    return [appPage]
  }

}

export {appPage, apiPage, apiPageResponse, apiPagesResponse, AdapterOfPages}