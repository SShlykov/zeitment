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
    "title": "Главная",
    "icon": "home-4-line",
    "link": "/",
    "type": "link",
    "name": "main"
  },
  {
    "title": "Главная",
    "icon": "home-4-line",
    "link": "/",
    "type": "link",
    "name": "main",
    "position": "top",
    "a": "v"
  }
]

const list = menuList()

export default list