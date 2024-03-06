class ServiceOfBooks {
  constructor(adapterOfBooks, store) {
    this.adapterOfBooks = adapterOfBooks;
    this.store = store;
  }

  async fetchBooksList() {
    const booksList = await this.adapterOfBooks.getBooks();
    this.store.dispatch('books/saveBooks', booksList);
  }
}

export default ServiceOfBooks;