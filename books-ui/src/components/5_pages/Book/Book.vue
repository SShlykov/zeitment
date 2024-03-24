<script>
import AdapterOfBooks from "@adapters/AdapterOfBooks/AdapterOfBooks.js";
import ServiceOfBooks from "@services/ServiceOfBooks.js";
import BookEditor from "@organisms/BookEditor/BookEditor.vue";
import ContentLoader from "@molecules/ContentLoader.vue";
import {mapGetters} from "vuex";
import BookManager from "@useCases/BookManager.js";
import ServiceOfLayout from "@services/ServiceOfLayout.js";
import ServiceOfChapters from "@services/ServiceOfChapters.js";
import ServiceOfPages from "@services/ServiceOfPages.js";
import {AdapterOfChapters} from "@mocks/chapters.js";
import {AdapterOfPages} from "@mocks/pages.js";

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
    ...mapGetters('books', ['currentBook', 'tableOfContents', "tableOfContentsSections"]),
    pageConfig() {
      return {
        bookId: this.$route.params.book_id,
        type: this.$route.params.type,
        sectionId: this.$route.params.section_id
      }
    },
    menuItems() {
      const addItemButton = {
        title: "Добавить...",
        class: "hover:bg-gray-100 transition-all cursor-pointer text-gray-500 hover:text-gray-700 text-lg p-2 rounded-md ",
        level: "button",
        icon: "ri-sticky-note-add-line",
        sort: false,
        onClick: () => {
        }
      }

      const sections = this.tableOfContentsSections
      const items = [...sections, addItemButton]
      return items
    }
  },
  watch:{
    $route (to){
      const bookId = to.params.book_id
      const type = to.params.type
      const sectionId = to.params.section_id
      this.bookManager.fetchBookWithPage(bookId, type, sectionId)
    }
  },
  mounted() {
    const url = import.meta.env.VITE_API_ADDR
    const adapterOfBooks = new AdapterOfBooks(url)
    const adapterOfChapters = new AdapterOfChapters(url)
    const adapterOfPages = new AdapterOfPages(url)
    const store = this.$store
    const bookId = this.pageConfig.bookId
    const type = this.pageConfig.type
    const sectionId = this.pageConfig.sectionId
    if (!bookId) {
      this.$router.push('/')
    }
    const serviceOfBooks = new ServiceOfBooks(adapterOfBooks, store)
    const serviceOfChapters = new ServiceOfChapters(adapterOfChapters, store)
    const serviceOfPages = new ServiceOfPages(adapterOfPages, store)
    const layoutService  = new ServiceOfLayout(store)
    const bookManager    = new BookManager(serviceOfBooks, serviceOfChapters, serviceOfPages, layoutService)

    bookManager.fetchBookWithPage(bookId, type, sectionId)

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
      :pageConfig="pageConfig"
      :menuItems="menuItems"
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
