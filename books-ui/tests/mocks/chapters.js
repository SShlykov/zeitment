const apiChapter = {
  "id": "99dfcc38-41d3-4967-bd6f-df22ad551cac",
  "created_at": "2024-03-03T23:14:21.8742+03:00",
  "updated_at": "2024-03-03T23:18:12.257658+03:00",
  "deleted_at": null,
  "title": "Chapter 2",
  "number": 1,
  "text": "tests",
  "book_id": "fb5e7d1d-38cd-4831-bae9-07b36080e3e7",
  "is_public": true,
  "map_link": null,
  "map_params_id": null
}

const appChapter = {
  "id": "99dfcc38-41d3-4967-bd6f-df22ad551cac",
  "createdAt": "2024-03-03T23:14:21.8742+03:00",
  "updatedAt": "2024-03-03T23:18:12.257658+03:00",
  "deletedAt": null,
  "title": "Chapter 2",
  "number": 1,
  "text": "tests",
  "bookId": "fb5e7d1d-38cd-4831-bae9-07b36080e3e7",
  "isPublic": true,
  "mapLink": null,
  "mapParamsId": null
}

const apiChaptersResponse = {
  data: [apiChapter],
  status: "ok"
}

const apiChapterResponse = {
  data: apiChapter,
  status: "ok"
}


class AdapterOfChapters {
  constructor(url) {
    this.url = url
  }

  async getChapters() {
    return [appChapter]
  }

  async updateChapter(newChapter) {
    return {
      ...appChapter,
      ...newChapter
    }
  }

  async createChapter() {
    return appChapter
  }

  async getChapterById() {
    return appChapter
  }

  async deleteChapterById() {
    return appChapter
  }

  async getChaptersByBookId() {
    return [appChapter]
  }
}

export {apiChapter, apiChapterResponse, apiChaptersResponse, appChapter, AdapterOfChapters}