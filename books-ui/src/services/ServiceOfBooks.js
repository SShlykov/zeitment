class ServiceOfBooks {
  constructor(adapterOfBooks, store) {
    this.adapterOfBooks = adapterOfBooks;
    this.store = store;
  }

  async fetchUserBooks() {
    const booksList = await this.adapterOfBooks.getBooks();
    this.store.dispatch('books/saveUserBooks', booksList);
  }

  async createBook(book) {
    const defaultBook = {
      title: 'Новая книга',
      "author": "Буглов В.А.",
      "owner": "e75aae0d-c1eb-4199-a1d8-2177f57d6a1e",
      "description": "",
      "is_public": false,
      "publication": null
    }

    book = book || defaultBook;

    const newBook = await this.adapterOfBooks.createBook(book);
    const books = await this.adapterOfBooks.getBooks();
    this.store.dispatch('books/saveUserBooks', books);
    return newBook;

  }

  async updateBook(book) {
    return await this.adapterOfBooks.updateBook(book);
  }

  async removeBook(id) {
    const book = await this.adapterOfBooks.deleteBookById(id);
    const books = await this.adapterOfBooks.getBooks();
    this.store.dispatch('books/saveUserBooks', books);
    return book;
  }

  /**
   * функция сохранения атрибута книги
   * @param {String} attribute - атрибут книги
   * @param value - значение атрибута
   * @returns {Promise<{[p: string]: *}|null>}
   */
  async storeEditableBookAttribute(attribute, value) {
    const book = this.store.getters['books/editableBook'];
    if (!book) return null
    const updatedBook = {
      ...book,
      [attribute]: value
    };
    await this.store.dispatch('books/saveEditableBook', updatedBook);
    return updatedBook
  }

  async getBookById(id) {
    return await this.adapterOfBooks.getBookById(id);
  }

  async saveEditableBookToServer() {
    const book = this.store.getters['books/editableBook'];
    if (!book) return null
    const updatedBook = await this.adapterOfBooks.updateBook(book);
    await this.store.dispatch('books/saveEditableBook', updatedBook);
    return updatedBook
  }

  async fetchEditableBook(id) {
    const book = await this.adapterOfBooks.getBookById(id);
    await this.store.dispatch('books/saveEditableBook', book);
    return book;
  }
}

export default ServiceOfBooks;