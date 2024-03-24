const initialState = {
  chapters: [],
  currentChapter: null,
}

const state = {
  ...initialState
};

const getters = {
  chapters:        state => state.chapters,
  currentChapter: state => state.currentChapter,
};

const mutations = {
  resetStore(state) {
    for (let key in state) {
      state[key] = initialState[key]
    }
  },
  setChapters(state, chapters) {
    state.chapters = chapters;
  },
  setCurrentChapter(state, chapter) {
    state.currentChapter = chapter;
  },
};

const actions = {
  saveChapters({ commit }, chapters) {
    commit('setChapters', chapters);
  },
  resetStore({ commit }) {
    commit('resetStore');
  },
  saveCurrentChapter({ commit }, chapter) {
    commit('setCurrentChapter', chapter);
  },
};

export const store = {
  namespaced: true,
  state,
  getters,
  mutations,
  actions,
};
