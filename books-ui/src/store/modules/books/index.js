const initialState = {
  userBooks: [],
  editableBook: null,
  tableOfContent: [],
}

const state = {
  ...initialState
};

const getters = {
  userBooks:    state => state.userBooks,
  editableBook: state => state.editableBook,
  tableOfContent: state => state.tableOfContent,
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
  setTableOfContent(state, tableOfContent) {
    state.tableOfContent = tableOfContent;
  }
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
  saveTableOfContent({ commit }, tableOfContent) {
    commit('setTableOfContent', tableOfContent);
  }
};

export const store = {
  namespaced: true,
  state,
  getters,
  mutations,
  actions,
};
