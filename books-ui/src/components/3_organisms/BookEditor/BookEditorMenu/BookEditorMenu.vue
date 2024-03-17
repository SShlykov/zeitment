<script>
import BookEditorMenuItem from "@organisms/BookEditor/BookEditorMenu/BookEditorMenuItem.vue";

export default {
  name: 'BookEditorMenuItem',
  components: {BookEditorMenuItem},
  props: {
    menuItems: {
      type: Array,
      default: () => []
    }
  },
  data() {
    return {
      isOpen: true
    }
  },
  computed: {},
  mounted() {},
  methods: {
    toggle() {
      this.isOpen = !this.isOpen;
    }
  }
}

</script>

<template>
  <div class="w-[100px] h-full border-r border-gray-200">
    <BookEditorMenuItem
      v-for="item in menuItems"
      :key="item.id"
      :item="item"
    />
  </div>
</template>
