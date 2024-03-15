const initialState = {
  chapters: [],
  editableChapter: null,
}

const state = {
  ...initialState
};

const getters = {
  chapters:        state => state.chapters,
  editableChapter: state => state.editableChapter,
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
  setEditableChapter(state, chapter) {
    state.editableChapter = chapter;
  },
};

const actions = {
  saveChapters({ commit }, chapters) {
    commit('setChapters', chapters);
  },
  resetStore({ commit }) {
    commit('resetStore');
  },
  saveEditableChapter({ commit }, chapter) {
    commit('setEditableChapter', chapter);
  },
};

export const store = {
  namespaced: true,
  state,
  getters,
  mutations,
  actions,
};
