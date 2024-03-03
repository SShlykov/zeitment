import {get, post, put, remove} from './apiHelpers.js'
import {convertList, convertObject} from '@helpers/adapter/adapter.js'
import {reverseObject} from '@helpers/objectUtils'
import {path} from 'ramda'

/**
 *
 * @param {Object} adapterFromApiConfig
 * @param {Object | undefined} params
 * @returns {Object | null | undefined}
 */
const adapterToApiFromParams = (adapterFromApiConfig, params) => {
  let adapterToApiConfig = path(['adapterToApiConfig'], params)

  if (adapterToApiConfig) {
    return adapterToApiConfig
  } else if (adapterFromApiConfig) {
    adapterToApiConfig = reverseObject(adapterFromApiConfig)
  } else {
    adapterToApiConfig = {}
  }

  return adapterToApiConfig
}

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
   * @constructor
   * @param {Object | null} adapterFromApiConfig - конфигурация адаптера для преобразования объектов из API
   * @param {Object | undefined} params - параметры
   * @param {Object} params.adapterToApiConfig - конфигурация адаптера для преобразования объектов в API, по дефолту преобразование происходит в обратном порядке из adapterFromApiConfig
   */
  constructor(adapterFromApiConfig = {}, params) {
    this.adapterFromApiConfig = adapterFromApiConfig
    this.adapterToApiConfig = adapterToApiFromParams(adapterFromApiConfig, params)
  }

  /**
   *
   * @returns {Promise<Object[]>}
   */
  async getBooks() {
    const books = await get('/books')
    if (this.adapterFromApiConfig) {
      return convertList(books, {config: this.adapterFromApiConfig})
    }
    return books
  }

  /**
   *
   * @param {Object} book
   * @returns {Promise<*|void|Object>}
   */
  async updateBook(book) {
    const bookToApi = convertObject(book, {config: this.adapterToApiConfig})
    const bookFromApi = await put(`/books/${book.id}`, bookToApi)
    if (this.adapterFromApiConfig) {
      return convertObject(bookFromApi, {config: this.adapterFromApiConfig})
    }
    return bookFromApi
  }

  /**
   *
   * @param {Object} book
   * @returns {Promise<*|void|Object>}
   */
  async createBook(book) {
    const bookToApi = convertObject(book, {config: this.adapterToApiConfig})
    const bookFromApi = await post('/books', bookToApi)
    if (this.adapterFromApiConfig) {
      return convertObject(bookFromApi, {config: this.adapterFromApiConfig})
    }
    return bookFromApi
  }


  /**
   *
   * @param {Number} id
   * @returns {Promise<Object>}
   */
  async getBookById(id) {
    const book = await get(`/books/${id}`)
    if (this.adapterFromApiConfig) {
      return convertObject(book, {config: this.adapterFromApiConfig})
    }
    return book
  }

  /**
   *
   * @param {Number} id
   * @returns {Promise<Object>}
   */
  async deleteBookById(id) {
    const book = await remove(`/books/${id}`)
    if (this.adapterFromApiConfig) {
      return convertObject(book, {config: this.adapterFromApiConfig})
    }
    return book
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
export {BooksInstance, adapterToApiFromParams}