<script>
import BookEditorHeader from "@organisms/BookEditor/BookEditorHeader/BookEditorHeader.vue";
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
  computed: {
    ...mapGetters('books', ['editableBook'])
  },
  mounted() {

  },
  methods: {
  }
}

</script>

<template>
  <div
    v-if="editableBook"
    class="w-full h-full flex flex-col"
  >
    <BookEditorHeader
      :serviceOfBooks="serviceOfBooks"
      :editableBook="editableBook"
      :bookManager="bookManager"
    />
    <div class="w-full flex flex-grow">
      <BookEditorChaptersMenu />
      <BookEditorBody />
    </div>
  </div>
  <div v-if="!editableBook">
    <div class="flex justify-center items-center h-full">
      <div class="text-2xl text-slate-400">
        Книги не существует
      </div>
    </div>
  </div>
</template>
