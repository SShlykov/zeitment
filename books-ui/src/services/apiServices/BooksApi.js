import {get, post, put, remove} from './apiHelpers.js'

/**
 * @deftypes Book
 * @property {String} title "Тестовая книга",
 * @property {String} author "Васильев А.В.",
 * @property {String} owner "e75aae0d-c1eb-4199-a1d8-2177f57d6a1e",
 * @property {String} description "test description",
 * @property {Boolean} is_public false,
 * @property {Any} publication null
 *
 */

/**
 * BooksApi class
 *
 * @class
 */

class BooksApi {
  /**
   *
   * @param {Number} id
   * @returns {Promise<Array<Book>>}
   */
  async getBooks() {
    return await get('/books')
  }

  /**
   *
   * @param {Book} book
   * @returns {Promise<Book>}
   */
  async updateBook(book) {
    console.log(`/books/${book.id}`, book)
    return await put(`/books/${book.id}`, book)
  }

  /**
   *
   * @param {Book} book
   * @returns {Promise<Book>}
   */
  async createBook(book) {
    return await post('/books', book)
  }


  /**
   *
   * @param {Number} id
   * @returns {Promise<Book>}
   */
  async getBookById(id) {
    return await get(`/books/${id}`)
  }

  /**
   *
   * @param {Number} id
   * @returns {Promise<Book>}
   */
  async deleteBookById(id) {
    return await remove(`/books/${id}`)
  }

  /**
   * @param {Function} logFunction
   * @returns {Promise<null>}
   */
  async integrationTests(logFunction) {
    logFunction("Список книг")
    const books = await this.getBooks()
    logFunction(books)
    logFunction("Создание книги")
    const newBook = await this.createBook({
      "title": "Тестовая книга",
      "author": "Васильев А.В.",
      "owner": "e75aae0d-c1eb-4199-a1d8-2177f57d6a1e",
      "description": "test description",
      "is_public": false,
      "publication": null
    })
    logFunction(newBook)
    logFunction("Получение книги по id(Созданной)")
    const bookById = await this.getBookById(newBook.id)
    logFunction(bookById)
    logFunction("Обновление книги")
    const updatedBook = await this.updateBook({...bookById, title: "Обновленная книга"})
    logFunction(updatedBook)
    logFunction("Удаление книги")
    const deletedBook = await this.deleteBookById(bookById.id)
    logFunction(deletedBook)
  }
}

const BooksInstance = new BooksApi()
export default BooksApi
export {BooksInstance}