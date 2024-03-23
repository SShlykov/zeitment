class BookManager {
  constructor(bookService, chapterService, pageService, layoutService) {
    this.bookService = bookService
    this.layoutService = layoutService
    this.chapterService = chapterService
    this.pageService = pageService
  }

  async saveBookWithPage() {
    const book = await this.bookService.saveCurrentBookToServer()
    this.layoutService.addNotification({
      message: "Книга сохранена",
      type: "success",
      timer: 2000
    })
    const answer = {
      book
    }
    return answer
  }

  /**
   *
   * @param {String} bookId
   * @param {String} type
   * @param {String} sectionId
   * @returns {Promise<{chapter: null, book: (Promise<*|Object|{owner: string, variables: [], author: string, description: string, title: string, mapParamsId: null, mapLink: null, createdAt: string, imageLink: null, deletedAt: null, publication: null, isPublic: boolean, id: string, updatedAt: string}>|any), page: null, tableOfContents: void}>}
   */
  async fetchBookWithPage(bookId, type, sectionId) {
    const {book, tableOfContents} = await this.bookService.fetchCurrentBook(bookId)

    let page = null
    let chapter = null
    if (type === "page") {
      page = await this.pageService.getPageById(sectionId)
    } else if (type === "chapter") {
      chapter = await this.chapterService.fetchChapterById(sectionId)
    }
    return {
      tableOfContents,
      book,
      page,
      chapter
    }
  }
}

export default BookManager