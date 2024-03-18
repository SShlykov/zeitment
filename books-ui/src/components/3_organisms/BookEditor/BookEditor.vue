<script>
import BookEditorHeader from "@organisms/BookEditor/BookEditorHeader/BookEditorHeader.vue";
import BookEditorMenu from "@organisms/BookEditor/BookEditorMenu/BookEditorMenu.vue";
import BookEditorBody from "@organisms/BookEditor/BookEditorBody.vue";
import {mapGetters} from "vuex";


export default {
  name: 'BookEditor',
  components: {BookEditorHeader, BookEditorMenu, BookEditorBody},
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
  computed: {
    ...mapGetters('books', ['currentBook', 'tableOfContents']),
    menuItems() {
      return this.tableOfContents.sections
    }
  },
  mounted() {

  },
  methods: {
  }
}

</script>

<template>
  <div
    v-if="currentBook"
    class="w-full h-full flex flex-col"
  >
    <BookEditorHeader
      :serviceOfBooks="serviceOfBooks"
      :currentBook="currentBook"
      :bookManager="bookManager"
    />
    <div class="w-full flex flex-grow">
      <BookEditorMenu
        :menuItems="menuItems"
      />
      <BookEditorBody />
    </div>
  </div>
  <div v-if="!currentBook">
    <div class="flex justify-center items-center h-full">
      <div class="text-2xl text-slate-400">
        Книги не существует
      </div>
    </div>
  </div>
</template>
