import {createRouter, createWebHistory} from "vue-router";
import Auth from "@pages/Auth/index.vue"
import UserBooks from "@pages/UserBooks/UserBooks.vue"
import NewBook from "@pages/NewBook/NewBook.vue"
import UserSettings from "@pages/UserSettings/UserSettings.vue"
import {cond, T} from 'ramda'

const routes = [
  {
    path: "/",
    name: "user_books",
    component: UserBooks,
  },
  {
    path: "/settings",
    name: "setting",
    component: UserSettings,
  },  {
    path: "/new_book",
    name: "new_book",
    component: NewBook,
  },
  {
    path: "/auth",
    name: "auth",
    component: Auth,
  },
];

const dev_routes = []

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
