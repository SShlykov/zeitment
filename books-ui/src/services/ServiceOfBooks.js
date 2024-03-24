class ServiceOfBooks {
  /**
   * @param {Object} adapterOfBooks
   * @param {Object} store
   */
  constructor(adapterOfBooks, store) {
    this.adapterOfBooks = adapterOfBooks;
    this.store = store;
  }

  /**
   * функция получения списка книг пользователя
   * @returns {Promise<void>}
   */
  async fetchUserBooks() {
    const booksList = await this.adapterOfBooks.getBooks();
    this.store.dispatch('books/saveUserBooks', booksList);
  }

  /**
   * функция создания книги
   * @param {Object} book
   * @returns {Promise<*|void|Object|{owner: string, variables: [], author: string, description: string, title: string, mapParamsId: null, mapLink: null, createdAt: string, imageLink: null, deletedAt: null, publication: null, isPublic: boolean, id: string, updatedAt: string}>}
   */
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

  /**
   * функция обновления книги
   * @param {Object} book
   * @returns {Promise<*|void|Object|(*&{owner: string, variables: [], author: string, description: string, title: string, mapParamsId: null, mapLink: null, createdAt: string, imageLink: null, deletedAt: null, publication: null, isPublic: boolean, id: string, updatedAt: string})>}
   */
  async updateBook(book) {
    return await this.adapterOfBooks.updateBook(book);
  }

  /**
   * функция удаления книги по id
   * @param {Number} id
   * @returns {Promise<Object|{owner: string, variables: [], author: string, description: string, title: string, mapParamsId: null, mapLink: null, createdAt: string, imageLink: null, deletedAt: null, publication: null, isPublic: boolean, id: string, updatedAt: string}|*>}
   */
  async removeBook(id) {
    const book = await this.adapterOfBooks.deleteBookById(id);
    const books = await this.adapterOfBooks.getBooks();
    this.store.dispatch('books/saveUserBooks', books);
    return book;
  }

  /**
   * функция сохранения атрибута книги в store
   * @param {String} attribute - атрибут книги
   * @param value - значение атрибута
   * @returns {Promise<{[p: string]: *}|null>}
   */
  async storeEditableBookAttribute(attribute, value) {
    const book = this.store.getters['books/currentBook'];
    if (!book) return null
    const updatedBook = {
      ...book,
      [attribute]: value
    };
    await this.store.dispatch('books/saveCurrentBook', updatedBook);
    return updatedBook
  }

  /**
   * функция получения книги по id
   * @param {Number} id
   * @returns {Promise<*|Object|{owner: string, variables: [], author: string, description: string, title: string, mapParamsId: null, mapLink: null, createdAt: string, imageLink: null, deletedAt: null, publication: null, isPublic: boolean, id: string, updatedAt: string}>}
   */
  async getBookById(id) {
    return await this.adapterOfBooks.getBookById(id);
  }

  /**
   * функция сохранения редактируемой книги на сервер
   * @returns {Promise<*|void|Object|(*&{owner: string, variables: [], author: string, description: string, title: string, mapParamsId: null, mapLink: null, createdAt: string, imageLink: null, deletedAt: null, publication: null, isPublic: boolean, id: string, updatedAt: string})|null>}
   */
  async saveCurrentBookToServer() {
    const book = this.store.getters['books/currentBook'];
    if (!book) return null
    const updatedBook = await this.adapterOfBooks.updateBook(book);
    await this.store.dispatch('books/saveCurrentBook', updatedBook);
    return updatedBook
  }

  /**
   * функция получения оглавление книги по id и сохранения его в store
   * @param {String} bookId
   * @returns {Promise<void>}
   */
  async fetchTableOfContent(bookId) {
    const tableOfContents = await this.adapterOfBooks.getTableOfContent(bookId);
    this.store.dispatch('books/saveTableOfContent', tableOfContents);
    return tableOfContents
  }

  /**
   * функция получения книги по id и сохранения ее в store
   * @param {Number} id
   * @returns {Promise<*|Object|{owner: string, variables: [], author: string, description: string, title: string, mapParamsId: null, mapLink: null, createdAt: string, imageLink: null, deletedAt: null, publication: null, isPublic: boolean, id: string, updatedAt: string}|{owner: string, variables: [], author: string, description: string, title: string, mapParamsId: null, mapLink: null, createdAt: string, imageLink: null, deletedAt: null, publication: null, isPublic: boolean, id: string, updatedAt: string}>}
   */
  async fetchEditableBook(id) {
    const book = await this.adapterOfBooks.getBookById(id);
    await this.store.dispatch('books/saveCurrentBook', book);
    return book;
  }

  /**
   * функция получения книги по bookId и оглавления ее и сохранения их в store
   * @param {Number} id
   * @returns {Promise<{book: {owner: string, variables: [], author: string, description: string, title: string, mapParamsId: null, mapLink: null, createdAt: string, imageLink: null, deletedAt: null, publication: null, isPublic: boolean, id: string, updatedAt: string}, tableOfContents: *}>}
   */
  async fetchCurrentBook(bookId) {
    const book = await this.fetchEditableBook(bookId);
    const tableOfContents = await this.fetchTableOfContent(bookId)
    return {
      book,
      tableOfContents
    }
  }
}

export default ServiceOfBooks;