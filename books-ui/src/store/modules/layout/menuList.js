/**
 * @typedef {'link' | 'line'}     Link
 */

/**
 * @typedef {Object}              MenuItem
 * @param   {String}              title
 * @param   {String}              icon
 * @param   {Link}                link
 * @param   {String}              type
 * @param   {String}              name
 * @param   {'top' | 'bottom'}    position
 */


/**
 *
 * @returns {Array<MenuItem>}
 */
const menuList = () => [
  {
    "title": "Создать книгу",
    "icon": "file-add-line",
    "link": "/new_book",
    "type": "link",
    "name": "new_book"
  },
  {
    "title": "Книга 1",
    "icon": "book-marked-line",
    "link": "/book/:link",
    "type": "link",
    "name": "book1"
  },  {
    "title": "Книга 2",
    "icon": "book-marked-line",
    "link": "/book/:link",
    "type": "link",
    "name": "book2"
  },
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

const list = menuList()

export default list