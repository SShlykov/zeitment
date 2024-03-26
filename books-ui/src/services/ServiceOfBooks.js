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
   * функция получения книги по id
   * @param {Number} id
   * @returns {Promise<*|Object|{owner: string, variables: [], author: string, description: string, title: string, mapParamsId: null, mapLink: null, createdAt: string, imageLink: null, deletedAt: null, publication: null, isPublic: boolean, id: string, updatedAt: string}>}
   */
  async fetchBookById(id) {
    return await this.adapterOfBooks.getBookById(id);
  }

  /**
   * функция получения списка книг пользователя
   * @returns {Promise<Object[]|*[]|[{owner: string, variables: [], author: string, description: string, title: string, mapParamsId: null, mapLink: null, createdAt: string, imageLink: null, deletedAt: null, publication: null, isPublic: boolean, id: string, updatedAt: string}]|*>}
   */
  async fetchUserBooks() {
    return await this.adapterOfBooks.getBooks()
  }

  /**
   * функция получения оглавления книги по bookId и сохранения его в store
   * @param bookId
   * @returns {Promise<{[p: string]: *}|{author: string, authorship: string, sections: [{level: string, isPublic: boolean, id: string, title: string, order: number},{level: string, isPublic: boolean, id: string, title: string, order: number},{level: string, isPublic: boolean, id: string, title: string, order: number},{level: string, isPublic: boolean, id: string, title: string, order: number}], bookId: string, bookTitle: string, tags: []}|*>}
   */
  async fetchTableOfContents(bookId) {
    const tableOfContents = await this.adapterOfBooks.getTableOfContent(bookId);
    return tableOfContents
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
   *
   * @param booksList
   * @returns {Array}
   */
  putUserBooks(booksList) {
    this.store.dispatch('books/saveUserBooks', booksList);
    return booksList
  }

  /**
   * функция получения списка книг пользователя из store
   * @returns {Promise<*>}
   */
  async getUserBooks() {
    return this.store.getters['books/userBooks']
  }

  /**
   *
   * @returns {Promise<*>}
   */
  getCurrentBook() {
    return this.store.getters['books/currentBook']
  }

  /**
   * функция сохранения оглавления книги в store
   * @param tableOfContents
   * @returns {Promise<*>}
   */
  async putTableOfContents(tableOfContents) {
    await this.store.dispatch('books/saveTableOfContent', tableOfContents);
    return tableOfContents
  }

  async makeFetchAndPutTableOfContents(bookId){
    const tableOfContents = await this.fetchTableOfContents(bookId)
    await this.putTableOfContents(tableOfContents)
    return tableOfContents
  }


  async makeFetchAndPutUserBooks(){
    const booksList = await this.fetchUserBooks()
    this.putUserBooks(booksList)
    return booksList
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
    const books = await this.fetchUserBooks();
    this.putUserBooks(books)
    return newBook;
  }


  /**
   * функция удаления книги по id
   * @param {Number} id
   * @returns {Promise<Object|{owner: string, variables: [], author: string, description: string, title: string, mapParamsId: null, mapLink: null, createdAt: string, imageLink: null, deletedAt: null, publication: null, isPublic: boolean, id: string, updatedAt: string}|*>}
   */
  async removeBook(id) {
    const book = await this.adapterOfBooks.deleteBookById(id);
    const books = await this.fetchUserBooks()
    this.putUserBooks(books)
    return book;
  }

  /**
   * функция сохранения атрибута книги в store
   * @param {String} attribute - атрибут книги
   * @param value - значение атрибута
   * @returns {Promise<{[p: string]: *}|null>}
   */
  async putBookAttribute(attribute, value) {
    const book = this.getCurrentBook();
    if (!book) return null
    const updatedBook = {
      ...book,
      [attribute]: value
    };
    await this.store.dispatch('books/saveCurrentBook', updatedBook);
    return updatedBook
  }


  /**
   * функция сохранения редактируемой книги на сервер
   * @returns {Promise<*|void|Object|(*&{owner: string, variables: [], author: string, description: string, title: string, mapParamsId: null, mapLink: null, createdAt: string, imageLink: null, deletedAt: null, publication: null, isPublic: boolean, id: string, updatedAt: string})|null>}
   */
  async saveCurrentBook() {
    const book = this.getCurrentBook();
    if (!book) return null
    const updatedBook = await this.adapterOfBooks.updateBook(book);
    await this.store.dispatch('books/saveCurrentBook', updatedBook);
    return updatedBook
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
    const tableOfContents = await this.makeFetchAndPutTableOfContents(bookId)
    return {
      book,
      tableOfContents
    }
  }
}

export default ServiceOfBooks;