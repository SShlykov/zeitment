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


}

export default ServiceOfBooks;