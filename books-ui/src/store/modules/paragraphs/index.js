const initialState = {
  paragraphs: [],
}

const state = {
  ...initialState
};

const getters = {
  paragraphs: state => state.paragraphs,
};

const mutations = {
  resetStore(state) {
    for (let key in state) {
      state[key] = initialState[key]
    }
  },
  setParagraphs(state, paragraphs) {
    state.paragraphs = paragraphs;
  }
};

const actions = {
  resetStore({ commit }) {
    commit('resetStore');
  },
  saveParagraphs({ commit }, paragraphs) {
    commit('setParagraphs', paragraphs);
  }
};

export const store = {
  namespaced: true,
  state,
  getters,
  mutations,
  actions,
};
