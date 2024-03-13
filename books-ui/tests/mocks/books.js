const appBook = {
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

const apiBook = {
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

const apiBooksResponse = {
  data: [apiBook],
  status: "ok"
}

const apiBookResponse = {
  data: apiBook,
  status: "ok"
}

class AdapterOfBooks {
  constructor(uri) {
    this.uri = uri
  }

  async getBooks() {
    return [appBook]
  }

  async updateBook(newBook) {
    return {
      ...appBook,
      ...newBook
    }
  }

  async createBook() {
    return appBook
  }

  async deleteBookById() {
    return appBook
  }

  async getBookById() {
    return appBook
  }
}

const bookStore = {
  namespaced: true,
  state: {
    userBooks: []
  },
  mutations: {
    setUserBooks(state, userBooks) {
      state.userBooks = userBooks
    },
    resetStore(state) {
      state.userBooks = []
    }
  },
  getters: {
    userBooks: (state) => state.userBooks
  },
  actions: {
    async saveUserBooks({commit}, userBooks) {
      commit('setUserBooks', userBooks)
    },
    resetStore({commit}) {
      commit('resetStore')
    }
  }
}

export { appBook, apiBook, apiBookResponse, apiBooksResponse, AdapterOfBooks, bookStore }