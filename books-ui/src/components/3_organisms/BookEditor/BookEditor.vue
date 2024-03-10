<script>
import BookEditorHeader from "@organisms/BookEditor/BookEditorHeader.vue";
import BookEditorChaptersMenu from "@organisms/BookEditor/BookEditorChaptersMenu.vue";
import BookEditorBody from "@organisms/BookEditor/BookEditorBody.vue";

export default {
  name: 'BookEditor',
  components: {BookEditorHeader, BookEditorChaptersMenu, BookEditorBody},
  data() {},
  computed: {},
  mounted() {},
  methods: {}
}

</script>

<template>
  <div class="w-full h-full">
    <BookEditorHeader />
    <div>
      <BookEditorChaptersMenu />
      <BookEditorBody />
    </div>
  </div>
</template>
