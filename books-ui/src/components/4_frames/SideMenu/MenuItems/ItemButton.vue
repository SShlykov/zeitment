<template>
  <div
    class="flex p-2 rounded-md hover:bg-gray-200 flex items-center"
    :class="{'bg-gray-200': pageName === name}"
    @click="itemFunction"
  >
    <i :class="`ri-${icon} text-slate-600 mr-2 `" />
    <span
      v-if="isOpenMenu"
      class="text-slate-700 truncate"
    >{{ title }}</span>
  </div>
</template>

<script>
import {mapGetters} from "vuex";

export default {
  name: 'ItemButton',
  components: {},
  props: {
    icon: {
      default: "",
      type: String
    },
    title: {
      default: "",
      type: String
    },
    itemFunction: {
      default: () => null,
      type: Function
    },
    name: {
      default: "",
      type: String
    },
  },
  computed: {
    ...mapGetters('layout', ['isOpenMenu', 'menuList']),
    pageName() {
      if (!this.$route) return ""
      return this.$route.name
    }
  },
  methods: {
  }
}

</script>
