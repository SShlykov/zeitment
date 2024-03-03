import BooksApi from "@apiServices/BooksApi.js"
import {adapterConfig} from "./StaticData.js";

const BooksService = new BooksApi(adapterConfig)

const initialState = {
  booksList: [],
}

const state = {
  ...initialState
};

const getters = {
  booksList: state => state.booksList,
};

const mutations = {
  resetStore(state) {
    Object.assign(state, initialState);
  },
  setBooksList(state, booksList) {
    state.booksList = booksList;
  },
};

const actions = {
  async fetchUserBooks({ commit }) {
    const booksList = await BooksService.getBooks()
    commit('setBooksList', booksList)
  },
};

export const store = {
  namespaced: true,
  state,
  getters,
  mutations,
  actions,
};
