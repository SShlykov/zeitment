class BookManager {
  constructor(bookService, chapterService, pageService, layoutService) {
    this.bookService = bookService
    this.layoutService = layoutService
    this.chapterService = chapterService
    this.pageService = pageService
  }

  async saveBookWithPage() {
    const book = await this.bookService.saveCurrentBook()
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


  /**
   *
   * @returns {{number: {null | string}, text: string, title: string}}
   */
  getCurrentSectionContent() {
    let title = ""
    let text = ""
    let number = null

    const currentChapter = this.chapterService.currentChapter()
    const currentPage = this.pageService.currentPage()
    if (currentPage) {
      title = currentPage.title
      text = currentPage.text
      number = null
    } else if (currentChapter) {
      title = currentChapter.title
      text = currentChapter.text
      number = currentChapter.number
    }

    return {
      title,
      text,
      number
    }
  }

  /**
   *
   * @param {{
   *   id: string,
   * }} sortableElement
   * @returns {Promise<void>}
   */
  async updateOrderTableOfContent(sortableElement, event) {
    console.log(sortableElement)
    console.log(event)
    // const tableOfContents = this.bookService.tableOfContents()
    // const newTableOfContents = tableOfContents.map((element) => {
    //   if (element.id === sortableElement.id) {
    //     element.order = sortableElement.order
    //   }
    //   return element
    // })
    const tableOfContents = await this.bookService.fetchTableOfContents(sortableElement.id)
    return tableOfContents
  }
}

export default BookManager