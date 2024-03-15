class ServiceOfChapters {
  /**
   * @param {Object} adapterOfChapters
   * @param {Object} store
   */
  constructor(adapterOfChapters, store) {
    this.adapterOfBooks = adapterOfChapters;
    this.store = store;
  }

  /**
   * Fetch chapters by book id
   * @param {String} bookId
   * @returns {Promise<void>}
   */
  async fetchChaptersByBookId(bookId) {
    const chaptersList = await this.adapterOfBooks.getChaptersByBookId(bookId);
    this.store.dispatch('chapters/saveChapters', chaptersList);
  }
}

export default ServiceOfChapters;