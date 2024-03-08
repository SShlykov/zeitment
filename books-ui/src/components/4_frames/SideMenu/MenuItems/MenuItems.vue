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

<script>
import ItemLink from './ItemLink.vue'
import ItemLine from './ItemLine.vue'
import ItemButton from './ItemButton.vue'

export default {
  name: 'MenuItems',
  components: {ItemLink, ItemLine, ItemButton},
  props: {
    menuList: {
      type: Array,
      required: true
    }
  }
}

</script>
