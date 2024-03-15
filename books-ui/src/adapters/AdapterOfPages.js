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
  "text": "text",
  "chapter_id": "chapterId",
  "is_public": "isPublic",
  "map_params_id": "mapParamsId"
}

/**
 * PagesApi class
 *
 * @class
 */

class AdapterOfPages {

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
  async getPages(params) {
    const defaultParams = {
      "page_size": 10,
      "page": 1
    }
    const {page, page_size} = fetchParamsByDefaultObject(params, defaultParams)
    if (!is(Number, page) || !is(Number, page_size) || page < 0 || page_size < 0) return  []
    let {data: books} = await post(`${this.url}/pages/list`, {
      "options": {page, page_size}
    })
    books = convertList(books, {config: this.adapterFromApiConfig})
    return books
  }

  /**
   *
   * @param {String} chapterId
   * @param {Object | undefined} params
   * @returns {Promise<Object[]>}
   */
  async getPagesByChapterId(chapterId, params) {
    const defaultParams = {
      "page_size": 10,
      "page": 1
    }
    const {page, page_size} = fetchParamsByDefaultObject(params, defaultParams)
    if (!is(Number, page) || !is(Number, page_size) || page < 0 || page_size < 0) return  []
    let {data: pages} = await post(`${this.url}/pages/chapters/${chapterId}`, {
      "options": {page, page_size}
    })
    pages = pages || []
    return convertList(pages, {config: this.adapterFromApiConfig})
  }

  /**
   *
   * @param {Object} book
   * @returns {Promise<*|void|Object>}
   */
  async updatePage(book) {
    const pageToApi = convertObject(book, {config: this.adapterToApiConfig})
    const {data: pageFromApi} = await put(`${this.url}/pages/${book.id}`, {
      page: pageToApi
    })

    const updatedPage = convertObject(pageFromApi, {config: this.adapterFromApiConfig})
    return updatedPage
  }

  /**
   *
   * @param {Object} page
   * @returns {Promise<*|void|Object>}
   */
  async createPage(page) {
    const pageToApi = convertObject(page, {config: this.adapterToApiConfig})
    const {data: pageFromApi} = await post(`${this.url}/pages`, {
      page: pageToApi
    })
    const createdPage = convertObject(pageFromApi, {config: this.adapterFromApiConfig})
    return createdPage
  }


  /**
   *
   * @param {Number} id
   * @returns {Promise<Object>}
   */
  async getPageById(id) {
    const {data: page} = await get(`${this.url}/pages/${id}`)
    return convertObject(page, {config: this.adapterFromApiConfig})
  }

  /**
   *
   * @param {Number} id
   * @returns {Promise<Object>}
   */
  async deletePageById(id) {
    const {data: page} = await remove(`${this.url}/pages/${id}`)
    const updatedPage = convertObject(page, {config: this.adapterFromApiConfig})
    return updatedPage
  }

  /**
   * @param {Function} logFunction
   * @returns {Promise<null>}
   */
  async integrationTests(logFunction) {
    logFunction("Список страниц")
    const books = await this.getPages()
    logFunction(books)
    logFunction("Список страниц в главе")
    const pagesByBook = await this.getPagesByChapterId("af3ff4ad-bc7d-4e64-acf8-bbd874d4516b")
    console.log(pagesByBook)
    logFunction("Создание страницы")
    const newPage = await this.createPage({
      "title": "Page 4",
      "text": "s adsadaseqwd sadsaqweq dsa",
      "chapterId": "af3ff4ad-bc7d-4e64-acf8-bbd874d4516b",
      "isPublic": false
    })
    logFunction(newPage)
    logFunction("Получение страницы по id(Созданной)")
    const bookById = await this.getPageById(newPage.id)
    logFunction(bookById)
    logFunction("Обновление страницы")
    const updatedPage = await this.updatePage({...bookById, title: "Обновленная страница"})
    logFunction(updatedPage)
    logFunction("Удаление страницы")
    const deletedPage = await this.deletePageById(bookById.id)
    logFunction(deletedPage)
  }
}

export default AdapterOfPages
