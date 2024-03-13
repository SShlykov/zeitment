<script>
import BookEditorHeader from "@organisms/BookEditor/BookEditorHeader.vue";
import BookEditorChaptersMenu from "@organisms/BookEditor/BookEditorChaptersMenu.vue";
import BookEditorBody from "@organisms/BookEditor/BookEditorBody.vue";
import {mapGetters} from "vuex";


export default {
  name: 'BookEditor',
  components: {BookEditorHeader, BookEditorChaptersMenu, BookEditorBody},
  props: {
    serviceOfBooks: {
      type: Object,
      default: () => {}
    },
    bookManager: {
      type: Object,
      default: () => {}
    }
  },
  data() {},
  computed: {
    ...mapGetters('books', ['editableBook'])
  },
  mounted() {

  },
  methods: {
    updateBookTitle(e) {
      this.serviceOfBooks.storeEditableBookAttribute('title', e.target.value)
    },
    updateBookAuthor(e) {
      this.serviceOfBooks.storeEditableBookAttribute('author', e.target.value)
    },
    saveBook() {
      this.bookManager.saveBookWithPage()
    }
  }
}

</script>

<template>
  <div class="w-full h-full">
    <BookEditorHeader class="flex justify-between">
      <div class="flex items-center">
        <input
          :value="editableBook.title"
          :onInput="updateBookTitle"
          :style="`width: ${(1 + editableBook.title.length)}ch`"
          placeholder="Название"
          class="h-full pl-2 rounded-md font-w-700 text-xl !ring-0 bg-white transition-all
                 min-w-[150px] max-w-[350px]
                 border-0 !outline-0
                 focus:bg-gray-100 hover:bg-gray-100"
        >
      </div>
      <div>
        <input
          :value="editableBook.author"
          :onInput="updateBookAuthor"
          :style="`width: ${(1 + editableBook.author.length)}ch`"
          placeholder="Автор"
          class="h-full pl-2 rounded-md mr-2 text-md !ring-0 bg-white transition-all
                 min-w-[100px] max-w-[350px]
                 border-0 !outline-0
                 focus:bg-gray-100 hover:bg-gray-100"
        >
        <button
          type="button"
          class="text-white bg-teal-700 hover:bg-teal-800 focus:outline-none focus:ring-4 focus:ring-teal-300 font-medium rounded-md text-sm px-4 py-1.5 text-center me-2 mb-2 dark:bg-teal-600 dark:hover:bg-teal-700 dark:focus:ring-teal-800"
          @click="saveBook"
        >
          <i class="ri-save-line" />
          Сохранить
        </button>
      </div>
    </BookEditorHeader>
    <div>
      <BookEditorChaptersMenu />
      <BookEditorBody />
    </div>
  </div>
</template>
