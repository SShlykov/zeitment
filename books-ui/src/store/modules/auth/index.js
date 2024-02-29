import axios from 'axios';

const initState = {
  user: null, 
  errorMsg: null
};

const state = {
  ...initState
};

const getters = {
  errorMsg: (state) => state.errorMsg
};

const mutations = {};

const actions = {
  signin: async ({state}, payload) => {
    const {data} = await axios.post(`${import.meta.env.VITE_API_ADDR}/login`, {
      "login": payload.login, "password": payload.password
    }, {
      headers: {
        'Content-Type': 'application/json'
      }
    })
    if (data.token != null) {
      const token = data.token
      localStorage.setItem('token', token)
      const pathname = sessionStorage.getItem("pathname") || "/"
      sessionStorage.removeItem("pathname")
      window.location.replace(pathname);
    } else {
      state.errorMsg = "неверный логин/пароль"
    }
  }, 
  resetStore: ({commit}) => {
    commit('resetStore')
  }
};

export const store = {
  namespaced: true, state, getters, mutations, actions,
};
