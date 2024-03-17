const initialState = {
  chapterPages: [],
  editablePage: null
}

const state = {
  ...initialState
};

const getters = {
  chapterPages: state => state.chapterPages,
  editablePage: state => state.editablePage,
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
  setEditablePage(state, page) {
    state.editablePage = page;
  },
};

const actions = {
  saveChapterPages({ commit }, pages) {
    commit('setChapterPages', pages);
  },
  resetStore({ commit }) {
    commit('resetStore');
  },
  saveEditablePage({ commit }, page) {
    commit('setEditablePage', page);
  },
};

export const store = {
  namespaced: true,
  state,
  getters,
  mutations,
  actions,
};