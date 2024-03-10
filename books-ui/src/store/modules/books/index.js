const initialState = {
  userBooks: [],
}

const state = {
  ...initialState
};

const getters = {
  userBooks: state => state.userBooks,
};

const mutations = {
  resetStore(state) {
    Object.assign(state, initialState);
  },
  setUserBooks(state, userBooks) {
    state.userBooks = userBooks;
  },
};

const actions = {
  async saveUserBooks({ commit }, userBooks) {
    commit('setUserBooks', userBooks);
  },
};

export const store = {
  namespaced: true,
  state,
  getters,
  mutations,
  actions,
};
