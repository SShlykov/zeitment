import {createRouter, createWebHistory} from "vue-router";
import Main from "@pages/Main/index.vue"
import Auth from "@pages/Auth/index.vue"
import {cond, T} from 'ramda'

const routes = [
  {
    path: "/",
    name: "main",
    component: Main,
  },
  {
    path: "/auth",
    name: "auth",
    component: Auth,
  },
];

const dev_routes = ["user_manage", "offers", "form", "statistic"]

const routerHistory = createWebHistory();

const router = createRouter({
    history: routerHistory,
    routes
});

const logout = (from, next) => {
  localStorage.removeItem("token")
  sessionStorage.setItem("pathname", from.fullPath)
  next()
}

// router.beforeEach(async (to, from, next) => {
//   const token = localStorage.getItem('token')
//   return cond([
//     [() => to.name === "auth" && from.name === "auth", () => next()],
//     [() => to.name === "auth", () => logout(from, () => next())],
//     [() => !token, () => logout(from, () => next({name: "auth"}))],
//     [T, () => next()],
//   ])()
// })


export {routes, dev_routes};
export default router;
