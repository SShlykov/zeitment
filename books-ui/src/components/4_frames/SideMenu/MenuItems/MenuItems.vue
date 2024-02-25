<template>
  <div class="flex flex-col justify-between">
    <div
      v-for="{title, icon, link, type, name} in menuList"
      :key="`${title}${name}`"
    >
      <Link
        :icon="icon"
        :title="title"
        :link="link"
        :name="name"
      />
      <Line v-if="type === 'line'" />
    </div>
  </div>
</template>

<script>
import Link from './Link.vue'
import Line from './Line.vue'

export default {
  name: 'MenuItems',
  components: {Link, Line},
  props: {
    menuList: {
      type: Array,
      required: true
    }
  }
}

</script>
