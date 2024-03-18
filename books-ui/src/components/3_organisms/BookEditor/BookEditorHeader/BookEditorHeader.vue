<script>
import BookEditorHeaderContainer from "./BookEditorHeaderContainer.vue";

export default {
  name: 'BookEditorHeader',
  components: {BookEditorHeaderContainer},
  props: {
    currentBook: {
      type: Object,
      required: true
    },
    serviceOfBooks: {
      type: Object,
      required: true
    },
    bookManager: {
      type: Object,
      required: true
    }
  },
  methods: {
    updateBookTitle(e) {
      this.serviceOfBooks.storeEditableBookAttribute('title', e.target.value)
      return "ok"
    },
    updateBookAuthor(e) {
      this.serviceOfBooks.storeEditableBookAttribute('author', e.target.value)
      return "ok"
    },
    saveBook() {
      this.bookManager.saveBookWithPage()
      return "ok"
    }
  }
}

</script>

<template>
  <BookEditorHeaderContainer class="flex justify-between">
    <div class="flex items-center">
      <input
        :value="currentBook.title"
        :onInput="updateBookTitle"
        :style="`width: ${(1 + currentBook.title.length)}ch`"
        placeholder="Название"
        class="h-full pl-2 rounded-md font-w-700 text-xl !ring-0 bg-white transition-all
                 min-w-[150px] max-w-[350px]
                 border-0 !outline-0
                 focus:bg-gray-100 hover:bg-gray-100"
      >
    </div>
    <div>
      <input
        :value="currentBook.author"
        :onInput="updateBookAuthor"
        :style="`width: ${(1 + currentBook.author.length)}ch`"
        placeholder="Автор"
        class="h-full pl-2 rounded-md mr-2  !ring-0 bg-white transition-all
                 text-md text-end
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
  </BookEditorHeaderContainer>
</template>
