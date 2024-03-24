<script>
import BookEditorHeader from "@organisms/BookEditor/BookEditorHeader/BookEditorHeader.vue";
import BookEditorMenu from "@organisms/BookEditor/BookEditorMenu/BookEditorMenu.vue";
import BookEditorBody from "@organisms/BookEditor/BookEditorBody.vue";
import {mapGetters} from "vuex";
import {IPageConfig} from "@organisms/BookEditor/interfaces.js";

export default {
  name: 'BookEditor',
  components: {BookEditorHeader, BookEditorMenu, BookEditorBody},
  props: {
    serviceOfBooks: {
      type: Object,
      required: true
    },
    bookManager: {
      type: Object,
      required: true
    },
    pageConfig: {
      type: IPageConfig,
      required: true
    }
  },
  computed: {
    ...mapGetters('books', ['currentBook', 'tableOfContents']),
    menuItems() {
      const addItemButton = {
        // <i class='ri-sticky-note-add-line'></i>
        title: "Добавить...",
        class: "hover:bg-gray-100 transition-all cursor-pointer text-gray-500 hover:text-gray-700 text-lg p-2 rounded-md ",
        level: "button",
        function: () => {
        }
      }

      const sections = this.tableOfContents.sections
      const items = [...sections, addItemButton]
      return items
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
        :pageConfig="pageConfig"
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
