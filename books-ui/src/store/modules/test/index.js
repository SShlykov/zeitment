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
  },
};

export const store = {
  namespaced: true,
  state,
  getters,
  mutations,
  actions,
};
