import {expect, test, describe, beforeEach} from 'vitest'
import {createStore} from "vuex";
import { store as pages } from '@/store/modules/pages';
import ServiceOfPages from '@services/ServiceOfPages.js'
import {AdapterOfPages, appPage} from "@mocks/pages.js"

describe('serviceOfPages', () => {
  const store = createStore({
    modules: {
      pages
    }
  })

  beforeEach(async () => {
    await store.dispatch('pages/resetStore')
  })

  const adapterOfPages = new AdapterOfPages('');
  const serviceOfPages = new ServiceOfPages(adapterOfPages, store);

  test('serviceOfPages is exist', () => {
    expect(serviceOfPages).toBeDefined()
  })

  test("fetch pagesList", async () => {
    await serviceOfPages.fetchChapterPages()
    const pagesList = store.getters['pages/chapterPages']
    expect(pagesList).toEqual([appPage])
  })

  test("create page", async () => {
    let editablePage = store.getters['pages/editablePage']
    expect(editablePage).toEqual(null)

    await serviceOfPages.createPage(appPage)

    editablePage = store.getters['pages/editablePage']
    expect(editablePage).toEqual(appPage)
  })

  test("update page", async () => {
    const updatedPage = {
      ...appPage,
      title: "new title"
    }
    const page = await serviceOfPages.updatePage(updatedPage)
    expect(page).toEqual(updatedPage)
  })

  test("remove page", async () => {
    const page = await serviceOfPages.removePage(appPage.id)
    expect(page).toEqual(appPage)
  })

  test("test of storeEditablePageAttribute", async () => {
    await store.dispatch('pages/saveEditablePage', appPage)
    await serviceOfPages.storeEditablePageAttribute("title", "new title")

    const editablePage = store.getters['pages/editablePage']
    expect(editablePage.title).toEqual("new title")
  })

  test("test of get page by id", async () => {
    await store.dispatch('pages/saveChapterPages', [appPage])
    const page = await serviceOfPages.getPageById(appPage.id)
    expect(page).toEqual(appPage)
  })

  test("test of save editable page to server", async () => {
    await store.dispatch('pages/saveEditablePage', appPage)
    const page = await serviceOfPages.saveEditablePageToServer()
    expect(page).toEqual(appPage)
  })
})