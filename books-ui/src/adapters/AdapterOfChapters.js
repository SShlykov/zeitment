import {get, post, put, remove} from '@helpers/apiHelpers.js'
import {convertList, convertObject} from '@helpers/adapter/adapter.js'
import {fetchParamsByDefaultObject, reverseObject} from '@helpers/objectUtils'
import {is} from "ramda";

const adapterConfig = {
  "id": "id",
  "created_at": "createdAt",
  "updated_at": "updatedAt",
  "deleted_at": "deletedAt",
  "title": "title",
  "number": "number",
  "text": "text",
  "book_id": "bookId",
  "is_public": "isPublic",
  "map_link": "mapLink",
  "map_params_id": "mapParamsId"
}


/**
 * ChaptersApi class
 *
 * @class
 */

class AdapterOfChapters {

  /**
   *
   * @param {String} url
   */
  constructor(url) {
    this.adapterFromApiConfig = adapterConfig
    this.adapterToApiConfig = reverseObject(this.adapterFromApiConfig)
    this.url = url
  }


  /**
   *
   * @param {Object | undefined} params
   * @returns {Promise<Object[]|*[]>}
   */
  async getChapters(params) {
    const defaultParams = {
      "page_size": 10,
      "page": 1
    }
    const {page, page_size} = fetchParamsByDefaultObject(params, defaultParams)
    if (!is(Number, page) || !is(Number, page_size) || page < 0 || page_size < 0) return  []
    let {data: books} = await post(`${this.url}/chapters/list`, {
      "options": {page, page_size}
    })
    books = convertList(books, {config: this.adapterFromApiConfig})
    return books
  }

  /**
   *
   * @param {String} bookId
   * @param {Object | undefined} params
   * @returns {Promise<Object[]>}
   */
  async getChaptersByBookId(bookId, params) {
    const defaultParams = {
      "page_size": 10,
      "page": 1
    }
    const {page, page_size} = fetchParamsByDefaultObject(params, defaultParams)
    if (!is(Number, page) || !is(Number, page_size) || page < 0 || page_size < 0) return  []
    const {data: chapters} = await post(`${this.url}/chapters/book/${bookId}`, {
      "options": {page, page_size}
    })
    return convertList(chapters, {config: this.adapterFromApiConfig})
  }

  /**
   *
   * @param {Object} book
   * @returns {Promise<*|void|Object>}
   */
  async updateChapter(book) {
    const chapterToApi = convertObject(book, {config: this.adapterToApiConfig})
    const {data: chapterFromApi} = await put(`${this.url}/chapters/${book.id}`, {
      chapter: chapterToApi
    })

    const updatedChapter = convertObject(chapterFromApi, {config: this.adapterFromApiConfig})
    return updatedChapter
  }

  /**
   *
   * @param {Object} chapter
   * @returns {Promise<*|void|Object>}
   */
  async createChapter(chapter) {
    const chapterToApi = convertObject(chapter, {config: this.adapterToApiConfig})
    const {data: chapterFromApi} = await post(`${this.url}/chapters`, {
      chapter: chapterToApi
    })
    const createdChapter = convertObject(chapterFromApi, {config: this.adapterFromApiConfig})
    return createdChapter
  }


  /**
   *
   * @param {Number} id
   * @returns {Promise<Object>}
   */
  async getChapterById(id) {
    const {data: chapter} = await get(`${this.url}/chapters/${id}`)
    return convertObject(chapter, {config: this.adapterFromApiConfig})
  }

  /**
   *
   * @param {Number} id
   * @returns {Promise<Object>}
   */
  async deleteChapterById(id) {
    const {data: chapter} = await remove(`${this.url}/chapters/${id}`)
    const updatedChapter = convertObject(chapter, {config: this.adapterFromApiConfig})
    return updatedChapter
  }

  /**
   * @param {Function} logFunction
   * @returns {Promise<null>}
   */
  async integrationTests(logFunction) {
    logFunction("Список глав")
    const books = await this.getChapters()
    logFunction(books)
    logFunction("Список глав по книге")
    const chaptersByBook = await this.getChaptersByBookId("fb5e7d1d-38cd-4831-bae9-07b36080e3e7")
    console.log(chaptersByBook)
    logFunction("Создание главы")
    const newChapter = await this.createChapter({
      "title": "Chapter 1",
      "number": 1,
      "text": "tests",
      "bookId": "fb5e7d1d-38cd-4831-bae9-07b36080e3e7",
      "isPublic": false
    })
    logFunction(newChapter)
    logFunction("Получение главы по id(Созданной)")
    const bookById = await this.getChapterById(newChapter.id)
    logFunction(bookById)
    logFunction("Обновление главы")
    const updatedChapter = await this.updateChapter({...bookById, title: "Обновленная глава"})
    logFunction(updatedChapter)
    logFunction("Удаление главы")
    const deletedChapter = await this.deleteChapterById(bookById.id)
    logFunction(deletedChapter)
  }
}

export default AdapterOfChapters
