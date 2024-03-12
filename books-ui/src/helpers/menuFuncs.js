/**
 * Convert book to menu item
 * @param book
 * @returns {{icon: string, link: string, name: string, title, type: string}}
 */
const bookToMenuItem = (book) => {
  return {
    title: book.title,
    icon: "book-marked-line",
    link: `/book/${book.id}`,
    type: "book",
    name: `${book.id}`,
  }
}

/**
 * Convert books list to menu list
 * @param booksList
 * @returns {*}
 */
const booksListToMenuList = (booksList) => {
  return booksList.map(bookToMenuItem)
}

export {bookToMenuItem, booksListToMenuList}