import menuList from "./menuList.js"

const initialState = {
  height:      0,
  width:       0,
  inited:      false,
  isOpenMenu:  true,
  notifications: []
}

const state = {
  ...initialState
};

const getters = {
  height:                  (state) => state.height,
  width:                   (state) => state.width,
  isOpenMenu:              (state) => state.isOpenMenu,
  menuWidth:               (state) => state.isOpenMenu ? 300 : 70,
  notifications:           (state) => state.notifications,
  menuList:                ()      => menuList,
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
  },
  setNotifications: (state, notifications) => {
    state.notifications = notifications
  },
  resetStore: (state) => {
    for (let key in state) {
      state[key] = initialState[key]
    }
  }
};

const actions = {
  initScreenSizeRecalc: ({commit}) => {
    commit("enableScreenListener")
  },
  setNotifications: ({commit}, notifications) => {
    commit("setNotifications", notifications)
  },
  addNotification: ({commit, state}, notification) => {
    const notifications = [...state.notifications, notification]
    commit("setNotifications", notifications)
  },
  removeNotification: ({commit, state}, id) => {
    const notifications = state.notifications.filter(n => n.id !== id)
    commit("setNotifications", notifications)
  },
  resetStore: ({commit}) => {
    commit("setNotifications", [])
  },
};

export const store = {
  namespaced: true,
  state,
  getters,
  mutations,
  actions,
};
