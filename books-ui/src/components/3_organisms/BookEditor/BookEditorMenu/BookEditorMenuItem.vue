<script>

export default {
  name: 'BookEditorMenuItem',
  components: {},
  props: {
    title: {
      type: String,
      default: '123'
    },
    level: {
      type: String,
      default: ""
    },
    id: {
      type: String,
      default: ""
    },
    order: {
      type: Number,
      default: 0
    },
    itemClass: {
      type: String,
      default: ""
    },
    bookId: {
      type: String,
      required: true
    },
    sectionId: {
      type: String,
      required: true
    },
    icon: {
      type: String,
      default: ""
    },
    onClick: {
      type: Function,
      default: () => {
      }
    }
  },
  mounted() {},
  methods: {}
}

</script>

<template>
  <div class="w-full">
    <router-link
      v-if="level === 'chapter'"
      :to="`/book/${bookId}/${level}/${id}`"
      class="block h-9 transition-all w-full cursor-pointer rounded-md p-2 pl-2 hover:bg-gray-100 font-w-600"
      :class="[{'bg-gray-100': sectionId === id }, {itemClass: true}]"
    >
      {{ title }}
    </router-link>
    <router-link
      v-if="level === 'page'"
      :to="`/book/${bookId}/${level}/${id}`"
      class="block h-9 pl-2 w-full   cursor-pointer"
    >
      <div
        class="w-full pl-2 p-2 hover:bg-gray-100 rounded-md transition-all"
        :class="[{'bg-gray-100': sectionId === id }, {itemClass: true}]"
      >
        {{ title }}
      </div>
    </router-link>
    <div
      v-if="level === 'button'"
      :class="itemClass"
      @click="onClick"
    >
      <i
        v-if="icon"
        :class="icon"
      />
      {{ title }}
    </div>
  </div>
</template>

