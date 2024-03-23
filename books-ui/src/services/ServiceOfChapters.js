class ServiceOfChapters {
  /**
   * @param {Object} adapterOfChapters
   * @param {Object} store
   */
  constructor(adapterOfChapters, store) {
    this.adapterOfChapters = adapterOfChapters;
    this.store = store;
  }

  /**
   * Fetch chapters by book id
   * @param {String} bookId
   * @returns {Promise<void>}
   */
  async fetchChaptersByBookId(bookId) {
    const chaptersList = await this.adapterOfChapters.getChaptersByBookId(bookId);
    this.store.dispatch('chapters/saveChapters', chaptersList);
  }

  async fetchChapterById(chapterId) {
    const chapter = await this.adapterOfChapters.getChapterById(chapterId);
    this.store.dispatch('chapters/saveCurrentChapter', chapter);
    return chapter
  }
}

export default ServiceOfChapters;