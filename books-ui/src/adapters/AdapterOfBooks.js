import {get, post, put, remove} from '@helpers/apiHelpers.js'
import {convertList, convertObject} from '@helpers/adapter/adapter.js'
import {fetchParamsByDefaultObject, reverseObject} from '@helpers/objectUtils'
import {is} from "ramda";

const adapterConfig = {
  "id": "id",
  "created_at": "createdAt",
  "updated_at": "updatedAt",
  "deleted_at": "deletedAt",
  "owner": "owner",
  "title": "title",
  "author": "author",
  "description": "description",
  "is_public": "isPublic",
  "publication": "publication",
  "image_link": "imageLink",
  "map_link": "mapLink",
  "map_params_id": "mapParamsId",
  "variables": "variables"
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

class AdapterOfBooks {

  /**
   * @constructor
   * @param {Object | undefined } adapterFromApiConfig - конфигурация адаптера для преобразования объектов из API
   * @param {Object | undefined | any} params - параметры
   * @param {Object} params.adapterToApiConfig - конфигурация адаптера для преобразования объектов в API, по дефолту преобразование происходит в обратном порядке из adapterFromApiConfig
   */
  constructor(url) {
    this.adapterFromApiConfig = adapterConfig
    this.adapterToApiConfig = reverseObject(this.adapterFromApiConfig)
    this.url = url
  }

  /**
   *
   * @returns {Promise<Object[]>}
   */
  async getBooks(params) {
    const defaultParams = {
      "page_size": 10,
      "page": 1
    }
    const {page, page_size} = fetchParamsByDefaultObject(params, defaultParams)
    console.log(params)
    if (!is(Number, page) || !is(Number, page_size) || page < 0 || page_size < 0) return  []
    console.log(`${this.url}/books/list`)
    let {data: books} = await post(`${this.url}/books/list`, {
      "options": {page, page_size}
    })
    books = convertList(books, {config: this.adapterFromApiConfig})
    return books
  }

  /**
   *
   * @param {Object} book
   * @returns {Promise<*|void|Object>}
   */
  async updateBook(book) {
    const bookToApi = convertObject(book, {config: this.adapterToApiConfig})
    const {data: bookFromApi} = await put(`${this.url}/books/${book.id}`, {
      book: bookToApi
    })

    const updatedBook = convertObject(bookFromApi, {config: this.adapterFromApiConfig})
    return updatedBook
  }

  /**
   *
   * @param {Object} book
   * @returns {Promise<*|void|Object>}
   */
  async createBook(book) {
    const bookToApi = convertObject(book, {config: this.adapterToApiConfig})
    let {data: bookFromApi} = await post(`${this.url}/books`, {
      book: bookToApi
    })
    bookFromApi = convertObject(bookFromApi, {config: this.adapterFromApiConfig})
    return bookFromApi
  }


  /**
   *
   * @param {Number} id
   * @returns {Promise<Object>}
   */
  async getBookById(id) {
    const {data: book} = await get(`${this.url}/books/${id}`)
    return convertObject(book, {config: this.adapterFromApiConfig})
  }

  /**
   *
   * @param {Number} id
   * @returns {Promise<Object>}
   */
  async deleteBookById(id) {
    const {data: book} = await remove(`${this.url}/books/${id}`)
    const updatedBook = convertObject(book, {config: this.adapterFromApiConfig})
    return updatedBook
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

export default AdapterOfBooks
