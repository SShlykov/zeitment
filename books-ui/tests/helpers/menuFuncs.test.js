import {describe, test, expect} from 'vitest'
import {bookToMenuItem, booksListToMenuList} from '@helpers/menuFuncs'
import {bookMock} from "@helpers/staticData.js";

// {
//   "title": "Создать книгу",
//     "icon": "file-add-line",
//     "link": "/new_book",
//     "type": "link",
//     "name": "new_book"
// },
// {
//   "title": "Книга 1",
//     "icon": "book-marked-line",
//     "link": "/book/:link",
//     "type": "link",
//     "name": "book1"
// },  {
//   "title": "Книга 2",
//       "icon": "book-marked-line",
//       "link": "/book/:link",
//       "type": "link",
//       "name": "book2"
// },

describe("bookToMenuItem", () => {
  test('convert book to menu item', () => {
    const menuItem = bookToMenuItem(bookMock)
    const expectedMenuItem = {
      "title": "Тестовая книга",
      "icon": "book-marked-line",
      "link": `/book/${bookMock.id}`,
      "type": "book",
      "name": `${bookMock.id}`
    }

    expect(menuItem).toEqual(expectedMenuItem)
  })
})

describe("booksListToMenuList", () => {
  test('convert books list to menu list', () => {
    const menuList = booksListToMenuList([bookMock])
    const expectedMenuList = [
      {
        "title": "Тестовая книга",
        "icon": "book-marked-line",
        "link": `/book/${bookMock.id}`,
        "type": "book",
        "name": `${bookMock.id}`
      }
    ]

    expect(menuList).toEqual(expectedMenuList)
  })
})
