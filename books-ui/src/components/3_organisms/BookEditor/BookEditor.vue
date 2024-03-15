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
  data() {},
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
  <div class="w-full h-full flex flex-col">
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
</template>
