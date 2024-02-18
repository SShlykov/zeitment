const state = {
  height:      0,
  width:       0,
  inited:      false,
  isOpenMenu:  false
};

const getters = {
  height:                  (state) => state.height,
  width:                   (state) => state.width,
  isOpenMenu:              (state) => state.isOpenMenu,
  menuWidth:               (state) => state.isOpenMenu ? 300 : 70
};

const mutations = {
  enableScreenListener: (state) => {
    state.height = window.innerHeight
    state.width = window.innerWidth
    const acc = !!state.inited
    if (!acc) {
      window.addEventListener("resize", (e) => {
        state.height= e.target.innerHeight
        state.width = e.target.innerWidth
      });
      state.inited = true
    }
  },
  toggleMenu: (state) => {
    state.isOpenMenu = !state.isOpenMenu
  }
};

const actions = {
  initScreenSizeRecalc: ({commit}) => {
    commit("enableScreenListener")
  },

};

export const store = {
    namespaced: true,
    state,
    getters,
    mutations,
    actions,
};
