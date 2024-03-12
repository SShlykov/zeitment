<script>
import ItemLink from './ItemLink.vue'
import ItemLine from './ItemLine.vue'
import ItemButton from './ItemButton.vue'
import ItemBook from './ItemBook.vue'

export default {
  name: 'MenuItems',
  components: {ItemLink, ItemLine, ItemButton, ItemBook},
  props: {
    menuList: {
      type: Array,
      required: true
    },
    isOpenMenu: {
      type: Boolean,
      default: false
    },
    serviceOfBooks: {
      type: Object,
      default: () => {}
    }
  }
}

</script>


<template>
  <div class="flex flex-col justify-between">
    <div
      v-for="{title, icon, link, type, name, itemFunction} in menuList"
      :key="`${title}${name}`"
    >
      <ItemLink
        v-if="type === 'link'"
        :icon="icon"
        :title="title"
        :link="link"
        :name="name"
        :isOpenMenu="isOpenMenu"
      />
      <ItemBook
        v-if="type === 'book'"
        :icon="icon"
        :title="title"
        :link="link"
        :name="name"
        :isOpenMenu="isOpenMenu"
        :serviceOfBooks="serviceOfBooks"
      />
      <ItemLine v-if="type === 'line'" />
      <ItemButton
        v-if="type === 'button'"
        :icon="icon"
        :title="title"
        :link="link"
        :name="name"
        :itemFunction="itemFunction"
      />
    </div>
  </div>
</template>

