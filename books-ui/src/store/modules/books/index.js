const initialState = {
  userBooks: [],
  currentBook: null,
  tableOfContents: {
    sections: []
  },
}

const state = {
  ...initialState
};

const getters = {
  userBooks:    state => state.userBooks,
  currentBook: state => state.currentBook,
  tableOfContents: state => state.tableOfContents,
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
  setCurrentBook(state, book) {
    state.currentBook = book;
  },
  setTableOfContent(state, tableOfContents) {
    state.tableOfContents = tableOfContents;
  }
};

const actions = {
  saveUserBooks({ commit }, userBooks) {
    commit('setUserBooks', userBooks);
  },
  resetStore({ commit }) {
    commit('resetStore');
  },
  saveCurrentBook({ commit }, book) {
    commit('setCurrentBook', book);
  },
  saveTableOfContent({ commit }, tableOfContents) {
    commit('setTableOfContent', tableOfContents);
  }
};

export const store = {
  namespaced: true,
  state,
  getters,
  mutations,
  actions,
};
