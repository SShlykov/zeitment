const initialState = {
  userBooks: [],
  editableBook: null
}

const state = {
  ...initialState
};

const getters = {
  userBooks:    state => state.userBooks,
  editableBook: state => state.editableBook,
};

const mutations = {
  resetStore(state) {
    for (let key in state) {
      state[key] = initialState[key]
    }
  },
  setUserBooks(state, userBooks) {
    state.userBooks = userBooks;
  },
  setEditableBook(state, book) {
    state.editableBook = book;
  },
};

const actions = {
  saveUserBooks({ commit }, userBooks) {
    commit('setUserBooks', userBooks);
  },
  resetStore({ commit }) {
    commit('resetStore');
  },
  saveEditableBook({ commit }, book) {
    commit('setEditableBook', book);
  },
};

export const store = {
  namespaced: true,
  state,
  getters,
  mutations,
  actions,
};
