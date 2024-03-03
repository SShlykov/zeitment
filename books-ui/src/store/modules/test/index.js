import BooksApi from "@apiServices/BooksApi.js"

const BooksService = new BooksApi()

const initialState = {
  integrationTestLog: [],
}

const state = {
  ...initialState
};

const getters = {
  integrationTestLog: (state) => state.integrationTestLog,
  jsonIntegrationTestLog: (state) => JSON.stringify(state.integrationTestLog),
};

const mutations = {
  setIntegrationTestLog(state, log) {
    state.integrationTestLog = log;
  }
};

const actions = {
  async startIntegrationTest({ commit, state }) {
    commit('setIntegrationTestLog', [])
    const logTest = (result) => commit('setIntegrationTestLog', [...state.integrationTestLog, result])
    await BooksService.integrationTests(logTest)
  },
};

export const store = {
  namespaced: true,
  state,
  getters,
  mutations,
  actions,
};
