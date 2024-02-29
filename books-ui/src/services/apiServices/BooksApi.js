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
}

const BooksInstance = new BooksApi()
export default BooksApi
export {BooksInstance}