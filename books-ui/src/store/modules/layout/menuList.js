/**
 * @typedef {'link' | 'line' | 'book'}     Type
 */

/**
 * @typedef {Object}              MenuItem
 * @param   {String}              title
 * @param   {String}              icon
 * @param   {String}              link
 * @param   {Type}                type
 * @param   {String}              name
 * @param   {'top' | 'bottom'}    position
 */


/**
 *
 * @returns {Array<MenuItem>}
 */
const menuList = () => [

  {
    "title": "Мои книги",
    "icon": "book-read-line",
    "link": "/",
    "type": "link",
    "name": "user_books",
    "position": "bottom",
  },
  {
    "title": "Найти книгу",
    "icon": "git-repository-line",
    "link": "/books_compendium",
    "type": "link",
    "name": "main",
    "position": "bottom",
  },
  {
    "title": "Найти автора",
    "icon": "user-search-line",
    "link": "/authors",
    "type": "link",
    "name": "authors",
    "position": "bottom",
  },
  {
    "title": "Настройки",
    "icon": "settings-4-line",
    "link": "/settings",
    "type": "link",
    "name": "settings",
    "position": "bottom",
  },
  {
    "title": "Выход",
    "icon": "logout-circle-r-line",
    "link": "/auth",
    "type": "link",
    "name": "logout",
    "position": "bottom",
  },
]

/**
 *
 * @returns {Array<MenuItem>}
 */
const devMenuList = () => [
  {
    "title": "Тесты",
    "icon": "test-tube-line",
    "link": "/tests",
    "type": "link",
    "name": "tests",
    "position": "bottom",
  },
]

const makeMenuList = () => {
  if (import.meta.env.VITE_ROUTES_MODE === "dev") {
    return [...devMenuList(), ...menuList()]
  } else {
    return menuList()
  }
}

const list = makeMenuList()

export default list