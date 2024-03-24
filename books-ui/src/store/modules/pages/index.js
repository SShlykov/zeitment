const initialState = {
  chapterPages: [],
  currentPage: null
}

const state = {
  ...initialState
};

const getters = {
  chapterPages: state => state.chapterPages,
  currentPage: state => state.currentPage,
};

const mutations = {
  resetStore(state) {
    for (let key in state) {
      state[key] = initialState[key]
    }
  },
  setChapterPages(state, pages) {
    state.chapterPages = pages;
  },
  setCurrentPage(state, page) {
    state.currentPage = page;
  }
};

const actions = {
  saveChapterPages({ commit }, pages) {
    commit('setChapterPages', pages);
  },
  resetStore({ commit }) {
    commit('resetStore');
  },
  saveCurrentPage({ commit }, page) {
    commit('setCurrentPage', page);
  },
};

export const store = {
  namespaced: true,
  state,
  getters,
  mutations,
  actions,
};