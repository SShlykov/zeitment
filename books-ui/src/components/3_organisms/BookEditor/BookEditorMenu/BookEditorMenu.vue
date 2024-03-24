<script>
import BookEditorMenuItem from "./BookEditorMenuItem.vue";
import {IPageConfig} from "@organisms/BookEditor/interfaces.js";
import {VueDraggableNext} from 'vue-draggable-next'

export default {
  name: 'BookEditorMenu',

  components: {BookEditorMenuItem, draggable: VueDraggableNext},
  props: {
    menuItems: {
      type: Array,
      default: () => []
    },
    pageConfig: {
      type: IPageConfig,
      required: true
    },
    onSort: {
      type: Function,
      default: () => null
    },
  },
  data() {
    return {
      isOpen: true
    }
  },
  computed: {
    bottomMenuItems() {
      return this.menuItems.filter(item => item.level === "button")
    },
    pagesAndChapters() {
      return this.menuItems.filter(item => item.level !== "button")
    }
  },
  mounted() {},
  methods: {
    toggle() {
      this.isOpen = !this.isOpen;
    },
    log(event) {
      console.log(event)
    },
  }
}

</script>

<template>
  <div
    class="w-[200px] flex flex-col h-full border-r border-gray-200 p-2"
  >
    <draggable
      class="w-full flex flex-col "
      :list="pagesAndChapters"
      animation="150"
      @change="(event) => onSort(event, pagesAndChapters)"
    >
      <div
        v-for="element in pagesAndChapters"
        :key="`chapter-${element.id}`"
        class="w-full"
      >
        <BookEditorMenuItem
          :id="element.id"
          :key="element.id"
          :title="element.title"
          :order="element.order"
          :level="element.level"
          :itemClass="element.class"
          :icon="element.icon"
          :onClick="element.onClick"
          :bookId="pageConfig.bookId"
          :sectionId="pageConfig.sectionId"
        />
        <div class="w-full">
          <draggable
            :list="element.items"
            animation="150"
            group="nested"
            class="flex flex-col h-full w-full"
            @change="(event) => onSort(event, element.items)"
          >
            <BookEditorMenuItem
              v-for="item in element.items"
              :id="item.id"
              :key="`page-${item.id}`"
              :title="item.title"
              :order="item.order"
              :level="item.level"
              :itemClass="item.class"
              :icon="item.icon"
              :onClick="item.onClick"
              :bookId="pageConfig.bookId"
              :sectionId="pageConfig.sectionId"
            />
          </draggable>
        </div>
      </div>
    </draggable>
    <BookEditorMenuItem
      v-for="item in bottomMenuItems"
      :id="item.id"
      :key="`button-${item.id}`"
      :title="item.title"
      :order="item.order"
      :level="item.level"
      :itemClass="item.class"
      :icon="item.icon"
      :onClick="item.onClick"
      :bookId="pageConfig.bookId"
      :sectionId="pageConfig.sectionId"
    />
  </div>
</template>
