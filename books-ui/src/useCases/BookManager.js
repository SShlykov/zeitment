class BookManager {
  constructor(bookService, layoutService) {
    this.bookService = bookService
    this.layoutService = layoutService
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
}

export default BookManager