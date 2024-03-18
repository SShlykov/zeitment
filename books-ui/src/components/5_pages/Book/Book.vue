<script>
import AdapterOfBooks from "@adapters/AdapterOfBooks/AdapterOfBooks.js";
import ServiceOfBooks from "@services/ServiceOfBooks.js";
import BookEditor from "@organisms/BookEditor/BookEditor.vue";
import ContentLoader from "@molecules/ContentLoader.vue";
import {mapGetters} from "vuex";
import BookManager from "@useCases/BookManager.js";
import ServiceOfLayout from "@services/ServiceOfLayout.js";

export default {
  name: 'BookPage',
  components: {BookEditor, ContentLoader},
  data() {
    return {
      serviceOfBooks: null,
      bookManager: null
    }
  },
  computed: {
    ...mapGetters('books', ['currentBook'])
  },
  mounted() {
    const url = import.meta.env.VITE_API_ADDR
    const adapterOfBooks = new AdapterOfBooks(url)
    const store = this.$store
    const bookId = this.$route.params.id
    if (!bookId) {
      this.$router.push('/')
    }
    const serviceOfBooks = new ServiceOfBooks(adapterOfBooks, store)
    const layoutService  = new ServiceOfLayout(store)
    const bookManager    = new BookManager(serviceOfBooks, layoutService)

    serviceOfBooks.fetchCurrentBook(bookId)

    this.bookManager = bookManager
    this.serviceOfBooks = serviceOfBooks
  },
  methods: {}
}

</script>

<template>
  <div class="relative w-full h-full">
    <BookEditor
      v-if="currentBook"
      :serviceOfBooks="serviceOfBooks"
      :bookManager="bookManager"
    />
    <ContentLoader
      v-if="!currentBook"
      :loaderSize="30"
      class="flex"
    >
      <span class="text-2xl">Загрузка книги...</span>
    </ContentLoader>
  </div>
</template>
