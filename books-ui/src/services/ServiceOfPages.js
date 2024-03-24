class ServiceOfPages {
  constructor(adapter, store) {
    this.adapter = adapter;
    this.store = store;
  }

  async fetchChapterPages(id) {
    const pages = await this.adapter.getPagesByChapterId(id);
    this.store.dispatch('pages/saveChapterPages', pages);
  }

  async createPage(aPage) {
    const thePage = await this.adapter.createPage(aPage);
    this.store.dispatch('pages/saveCurrentPage', thePage);
  }

  async updatePage(page) {
    await this.adapter.updatePage(page);
    this.store.dispatch('pages/saveChapterPages', [page]);
    return page;
  }

  async removePage(id) {
    const page = await this.adapter.removePage(id);
    return page;
  }

  async storeEditablePageAttribute(attribute, value) {
    const page = this.store.getters['pages/currentPage'];
    if (!page) return null;
    const updatedPage = {
      ...page,
      [attribute]: value
    };
    await this.store.dispatch('pages/saveCurrentPage', updatedPage);
  }

  async saveEditablePageToServer() {
    const page = this.store.getters['pages/currentPage'];
    await this.adapter.updatePage(page);
    return page;
  }

  async getPageById(id) {
    const page = await this.adapter.getPageById(id);
    this.store.dispatch('pages/saveCurrentPage', page);
    return page
  }

  currentPage() {
    return this.store.getters['pages/currentPage'];
  }
}

export default ServiceOfPages;