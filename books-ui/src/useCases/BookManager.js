class BookManager {
  constructor(bookService) {
    this.bookService = bookService
  }

  async saveBookWithPage(bookData) {
    const book = await this.bookService.updateBook(bookData)
    const answer = {
      book
    }
    return answer
  }
}

export default BookManager