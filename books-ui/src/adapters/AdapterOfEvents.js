import {get, post, put, remove} from '@helpers/apiHelpers.js'
import {convertList, convertObject} from '@helpers/adapter/adapter.js'
import {fetchParamsByDefaultObject, reverseObject} from '@helpers/objectUtils'
import {is} from "ramda";


const adapterConfig = {
  "id": "id",
  "created_at": "createdAt",
  "updated_at": "updatedAt",
  "book_id": "bookId",
  "chapter_id": "chapterId",
  "page_id": "pageId",
  "paragraph_id": "paragraphId",
  "event_type": "eventType",
  "is_public": "isPublic",
  "key": "key",
  "value": "value",
  "link": "link",
  "link_text": "linkText",
  "link_type": "linkType",
  "link_image": "linkImage",
  "description": "description"
}


/**
 * ChaptersApi class
 *
 * @class
 */

class AdapterOfEvents {
  
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
   * @param {Object} event
   * @returns {Promise<Object>}
   */
  async updateEvent(event) {
    const eventToApi = convertObject(event, {config: this.adapterToApiConfig})
    const {data: eventFromApi} = await put(`${this.url}/bookevents/${event.id}`, {
      book_event: eventToApi
    })
    const updatedEvent = convertObject(eventFromApi, {config: this.adapterFromApiConfig})
    return updatedEvent
  }

  /**
   *
   * @param {Object} event
   * @returns {Promise<Object>}
   */
  async createEvent(event) {
    const eventToApi = convertObject(event, {config: this.adapterToApiConfig})
    let {data: eventFromApi} = await post(`${this.url}/bookevents`, {
      book_event: eventToApi
    })
    eventFromApi = convertObject(eventFromApi, {config: this.adapterFromApiConfig})
    return eventFromApi
  }

  /**
   *
   * @param {String} paragraphId
   * @returns {Promise<Object[]>}
   */
  async deleteEventById(id) {
    const {data: eventFromApi} = await remove(`${this.url}/bookevents/${id}`)
    return convertObject(eventFromApi, {config: this.adapterFromApiConfig})
  }

  /**
   *
   * @param {String} paragraphId
   * @returns {Promise<Object[]>}
   */
  async getEventById(id) {
    const {data: event} = await get(`${this.url}/bookevents/${id}`)
    return convertObject(event, {config: this.adapterFromApiConfig})
  }

  /**
   *
   * @param {String} paragraphId
   * @returns {Promise<Object[]>}
   */
  async getEventsByBookId(bookId) {
    const {data: events} = await post(`${this.url}/bookevents/book/${bookId}`)
    return convertList(events, {config: this.adapterFromApiConfig})
  }

  async getEventsByChapterId(chapterId) {
    const {data: events} = await post(`${this.url}/bookevents/chapter/${chapterId}`)
    return convertList(events, {config: this.adapterFromApiConfig})
  }

  /**
   *
   * @param {String} paragraphId
   * @returns {Promise<Object[]>}
   */
  async getEventsByPageId(pageId) {
    const {data: events} = await post(`${this.url}/bookevents/page/${pageId}`)
    return convertList(events, {config: this.adapterFromApiConfig})
  }

  /**
   *
   * @param {String} paragraphId
   * @returns {Promise<Object[]>}
   */
  async getEventsByParagraphId(paragraphId) {
    const {data: events} = await post(`${this.url}/bookevents/paragraph/${paragraphId}`)
    return convertList(events, {config: this.adapterFromApiConfig})
  }

  /**
   *
   * @param {Function} logFunction
   * @returns {Promise<null>}
   */
  async integrationTests(logFunction) {
    logFunction("Создать событие")
    const randomKey = Math.random().toString(36).slice(4)
    let eventParams = {
      "bookId": "fb5e7d1d-38cd-4831-bae9-07b36080e3e7",
      "isPublic": true,
      "key": randomKey,
      "value": Math.random().toString(36).slice(4),
      "eventType": "test"
    }
    const createdEvent = await this.createEvent(eventParams)
    logFunction(createdEvent)
    logFunction("Получить событие по id")
    const eventById = await this.getEventById(createdEvent.id)
    logFunction(eventById)
    logFunction("Обновить событие")
    eventParams = {
      ...eventById,
      "key": randomKey,
      value: Math.random().toString(36).slice(4)
    }
    const updatedEvent = await this.updateEvent(eventParams)
    logFunction(updatedEvent)
    logFunction("Удалить событие")
    const deletedEvent = await this.deleteEventById(updatedEvent.id)
    logFunction(deletedEvent)
    logFunction("Список событий по книге")
    const eventsByBook = await this.getEventsByBookId("fb5e7d1d-38cd-4831-bae9-07b36080e3e7")
    logFunction(eventsByBook)
    logFunction("Список событий по главе")
    const eventsByChapter = await this.getEventsByChapterId("af3ff4ad-bc7d-4e64-acf8-bbd874d4516b")
    logFunction(eventsByChapter)
    logFunction("Список событий по странице")
    const eventsByPage = await this.getEventsByPageId("2b15d86b-e52c-4d6f-9629-0bf3bc940f29")
    logFunction(eventsByPage)
    logFunction("Список событий по параграфу")
    const eventsByParagraph = await this.getEventsByParagraphId("12b9b045-0845-462c-b372-0fca3180a6af")
    logFunction(eventsByParagraph)

  }

}

export default AdapterOfEvents