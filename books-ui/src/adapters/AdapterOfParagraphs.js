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
  "type": "type",
  "is_public": "isPublic",
  "page_id": "pageId"
}

/**
 * ParagraphsApi class
 *
 * @class
 */

class AdapterOfParagraphs {

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
  async getParagraphs(params) {
    const defaultParams = {
      "page_size": 10,
      "page": 1
    }
    const {page, page_size} = fetchParamsByDefaultObject(params, defaultParams)
    if (!is(Number, page) || !is(Number, page_size) || page < 0 || page_size < 0) return  []
    let {data: paragraphs} = await post(`${this.url}/paragraphs/list`, {
      "options": {page_size, page}
    })
    paragraphs = convertList(paragraphs, {config: this.adapterFromApiConfig})
    return paragraphs
  }

  /**
   *
   * @param {String} chapterId
   * @param {Object | undefined} params
   * @returns {Promise<Object[]>}
   */
  async getParagraphsByPageId(chapterId, params) {
    const defaultParams = {
      "page_size": 10,
      "page": 1
    }
    const {page, page_size} = fetchParamsByDefaultObject(params, defaultParams)
    console.log({page, page_size})
    if (!is(Number, page) || !is(Number, page_size) || page < 0 || page_size < 0) return  []
    let {data: paragraphs} = await post(`${this.url}/paragraphs/pages/${chapterId}`, {
      "options": {page, page_size}
    })
    paragraphs = paragraphs || []
    return convertList(paragraphs, {config: this.adapterFromApiConfig})
  }

  /**
   *
   * @param {Object} book
   * @returns {Promise<*|void|Object>}
   */
  async updateParagraph(book) {
    const paragraphToApi = convertObject(book, {config: this.adapterToApiConfig})
    const {data: paragraphFromApi} = await put(`${this.url}/paragraphs/${book.id}`, {
      paragraph: paragraphToApi
    })

    const updatedParagraph = convertObject(paragraphFromApi, {config: this.adapterFromApiConfig})
    return updatedParagraph
  }

  /**
   *
   * @param {Object} paragraph
   * @returns {Promise<*|void|Object>}
   */
  async createParagraph(paragraph) {
    const paragraphToApi = convertObject(paragraph, {config: this.adapterToApiConfig})
    const {data: paragraphFromApi} = await post(`${this.url}/paragraphs`, {
      paragraph: paragraphToApi
    })
    const createdParagraph = convertObject(paragraphFromApi, {config: this.adapterFromApiConfig})
    return createdParagraph
  }


  /**
   *
   * @param {Number} id
   * @returns {Promise<Object>}
   */
  async getParagraphById(id) {
    const {data: paragraph} = await get(`${this.url}/paragraphs/${id}`)
    return convertObject(paragraph, {config: this.adapterFromApiConfig})
  }

  /**
   *
   * @param {Number} id
   * @returns {Promise<Object>}
   */
  async deleteParagraphById(id) {
    const {data: paragraph} = await remove(`${this.url}/paragraphs/${id}`)
    const updatedParagraph = convertObject(paragraph, {config: this.adapterFromApiConfig})
    return updatedParagraph
  }

  /**
   * @param {Function} logFunction
   * @returns {Promise<null>}
   */
  async integrationTests(logFunction) {
    logFunction("Список параграфов")
    const books = await this.getParagraphs()
    logFunction(books)
    logFunction("Список параграфов на странице")
    const paragraphsByBook = await this.getParagraphsByPageId("2b15d86b-e52c-4d6f-9629-0bf3bc940f29")
    logFunction(paragraphsByBook)
    logFunction("Создание параграфа")
    const newParagraph = await this.createParagraph({
      "title": "Paragraph 7",
      "text": "fjhdsjkhfdkjghdslkhfskldjhgfskldjbvsfkdbvksfdab",
      "type": "simple-text",
      "isPublic": false,
      "pageId": "2b15d86b-e52c-4d6f-9629-0bf3bc940f29"
    })
    logFunction(newParagraph)
    logFunction("Получение параграфа по id(Созданной)")
    const bookById = await this.getParagraphById(newParagraph.id)
    logFunction(bookById)
    logFunction("Обновление параграфа")
    const updatedParagraph = await this.updateParagraph({...bookById, title: "Обновленная параграфа"})
    logFunction(updatedParagraph)
    logFunction("Удаление параграфа")
    const deletedParagraph = await this.deleteParagraphById(bookById.id)
    logFunction(deletedParagraph)
  }
}

export default AdapterOfParagraphs
