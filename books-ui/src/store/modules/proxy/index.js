import sites from './sites.json'
import routes from './routes.json'
import routeOptions from './routeOptions.json'
import { get } from '../apiHelpers';

const initState = {
  sites: [],
  sitesState: "loading",

  routes: [],
  routesState: "await",
  selectedSite: null,
  routesLib: {},

  newSite: null,
  newRoute: null,
};

const state = {
  ...initState
};

const getters = {
  sites:        (state) => state.sites,
  routes:       (state) => state.routes,
  routesLib:    (state) => state.routesLib,
  routesState:  (state) => state.routesState,
  sitesState:   (state) => state.sitesState,
  newSite:      (state) => state.newSite,
  newRoute:     (state) => state.newRoute,
  routeOptions: () => routeOptions,
  selectedSite: (state) => state.selectedSite,
};

const mutations = {
};

const actions = {
  addNewSiteLayout: ({state}) => {
    const newSite = {"port": "", "name": "", id: -1}
    state.newSite = newSite
    state.selectedSite = newSite
  },
  createNewSite: ({state}, payload) => {
    state.selectedSite = null
    state.sites = [{
      ...state.newSite,
      id: state.sites.reduce((acc, {id}) => Math.max(id, acc) + 1, 0)
    }, ...state.sites]
    state.newSite = null
  },
  breakeAddingSite: ({state}) => {
    state.newSite = null
    state.selectedSite = null
  },
  addNewRouteLayout: ({state}) => {
    const newRoute = {"path": null, "role": null, id: -1}
    state.newRoute = newRoute
  },
  createNewRoute: ({state}, payload) => {
    console.log(payload)
    state.newRoute = null
  },
  breateAddingRoute: ({state}) => {
    state.newRoute = null
  },
  uploadSites: async ({state}) => {

    const sites = await get("/servers")
    console.log(sites)


    setTimeout(() => {
      state.sitesState = 'active'
      state.sites = sites
    }, 2000);
  },
  uploadSiteRoutes: ({state}, site_props) => {
    const {site_id} = site_props
    state.routesState = "loading"
    setTimeout(() => {
      state.routes = routes
      state.routesLib = routes.reduce((acc, cur) => ({
        ...acc,
        [cur.id]: cur
      }), {})
      state.selectedSite = site_props
      state.routesState = "active"
    }, 1000);
  },
  updateRouteRow: ({state}, row) => {
    console.log(row)

  },
  removeRoute: ({state}, {id}) => {
    state.routes = state.routes.filter((el) => el.id !== id)
  },
  closeRoutesList: ({state}) => {
    state.routesState = "await"
  },
  saveSite: ({state}, payload) => {
    console.log(payload)
    state.selectedSite = null
  },
  breakSavingSite: ({state}) => {
    state.selectedSite = null
  },

};

export const store = {
  namespaced: true,
  state,
  getters,
  mutations,
  actions,
};
