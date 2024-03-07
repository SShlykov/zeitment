class ServiceOfBooks {
  constructor(adapterOfBooks, store) {
    this.adapterOfBooks = adapterOfBooks;
    this.store = store;
  }

  async fetchUserBooks() {
    const booksList = await this.adapterOfBooks.getBooks();
    this.store.dispatch('books/saveUserBooks', booksList);
  }
}

export default ServiceOfBooks;