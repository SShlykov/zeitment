<template>
  <router-link
    class="flex p-2 rounded-md hover:bg-gray-100 flex items-center"
    :class="{'bg-gray-200': pageName === name}"
    :to="link"
  >
    <i :class="`ri-${icon} text-slate-600 mr-2 `" />
    <span
      v-if="isOpenMenu"
      class="text-slate-700"
    >{{ title }}</span>
  </router-link>
</template>

<script>
import {mapGetters, mapMutations} from "vuex";

export default {
  name: 'MenuLink',
  props: {
    icon: {
      default: "",
      type: String
    },
    title: {
      default: "",
      type: String
    },
    link: {
      default: "",
      type: String
    },
  },
  components: {},
  computed: {
    ...mapGetters('layout', ['isOpenMenu', 'menuList']),
    pageName() {
      return this.$route.name
    }
  },
  methods: {
  }
}

</script>
