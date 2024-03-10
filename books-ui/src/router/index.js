import {createRouter, createWebHistory} from "vue-router";
import Auth from "@pages/Auth/index.vue"
import UserBooks from "@pages/UserBooks/UserBooks.vue"
import NewBook from "@pages/NewBook/NewBook.vue"
import Book from "@pages/Book/Book.vue"
import UserSettings from "@pages/UserSettings/UserSettings.vue"
import TestsPage from "@pages/Tests/TestsPage.vue"


const prod_routes = [
  {
    path: "/",
    name: "user_books",
    component: UserBooks,
  },
  {
    path: "/settings",
    name: "setting",
    component: UserSettings,
  },
  {
    path: "/book/:id",
    name: "book",
    component: Book,
  },
  {
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

const dev_routes = [
  {
    path: "/tests",
    name: "tests",
    component: TestsPage,
  },
]

const makeRoutes = () => {
  if (import.meta.env.VITE_ROUTES_MODE === "dev") {
    return [...prod_routes, ...dev_routes]
  } else {
    return prod_routes
  }
}

const routes = makeRoutes()

const routerHistory = createWebHistory();

const router = createRouter({
  history: routerHistory,
  routes
});


// const logout = (from, next) => {
//   localStorage.removeItem("token")
//   sessionStorage.setItem("pathname", from.fullPath)
//   next()
// }

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
