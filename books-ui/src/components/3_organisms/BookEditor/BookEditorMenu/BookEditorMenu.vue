<script>
import BookEditorMenuItem from "./BookEditorMenuItem.vue";
import {IPageConfig} from "@organisms/BookEditor/interfaces.js";

export default {
  name: 'BookEditorMenu',
  components: {BookEditorMenuItem},
  props: {
    menuItems: {
      type: Array,
      default: () => []
    },
    pageConfig: {
      type: IPageConfig,
      required: true
    }
  },
  data() {
    return {
      isOpen: true
    }
  },
  mounted() {},
  methods: {
    toggle() {
      this.isOpen = !this.isOpen;
    }
  }
}

</script>


<template>
  <div class="w-[200px] flex flex-col h-full border-r border-gray-200 p-2">
    <BookEditorMenuItem
      v-for="item in menuItems"
      :id="item.id"
      :key="item.id"
      :title="item.title"
      :order="item.order"
      :level="item.level"
      :itemClass="item.class"
      :bookId="pageConfig.bookId"
      :sectionId="pageConfig.sectionId"
    />
  </div>
</template>
