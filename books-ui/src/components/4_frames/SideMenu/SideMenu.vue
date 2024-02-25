<template>
  <div
    class="h-[100vh] z-40 bg-white flex flex-col px-4 pt-8 border-r border-slate-200 transition-all"
    :class="{'w-[300px]': isOpenMenu, 'w-[70px]': !isOpenMenu}"
  >
    <MenuHead />
    <MenuList />
  </div>
</template>

<script>
import {mapGetters, mapMutations} from "vuex";
import MenuHead from './MenuHead.vue'
import MenuList from './MenuList.vue'

export default {
  name: 'App',
  components: {MenuHead, MenuList},
  data() {
  },
  computed: {
    ...mapGetters('layout', ['isOpenMenu', 'menuList']),
    pageName() {
      return this.$route.name
    }
  },
  methods: {
    ...mapMutations('layout', ['toggleMenu']),
  }
}

</script>
