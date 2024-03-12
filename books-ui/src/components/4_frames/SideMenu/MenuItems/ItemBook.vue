<template>
  <router-link
    class="flex p-2 rounded-md hover:bg-gray-200 items-center"
    :class="{'bg-gray-200': pageName === name}"
    :to="link"
  >
    <i :class="`ri-${icon} text-slate-600 mr-2 `" />
    <span
      v-if="isOpenMenu"
      class="text-slate-700 truncate"
    >{{ title }}</span>
  </router-link>
</template>

<script>

export default {
  name: 'ItemBook',
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
    link: {
      default: "",
      type: String
    },
    name: {
      default: "",
      type: String
    },
    isOpenMenu: {
      type: Boolean,
      default: false
    }
  },
  computed: {
    pageName() {
      if (!this.$route) return ""
      return this.$route.name
    }
  },
  methods: {
  }
}

</script>
