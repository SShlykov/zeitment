<template>
  <div
    class="h-[100vh] z-40  flex flex-col px-4 pt-8 border-r border-slate-200 transition-all bg-gray-100"
    :class="{'w-[300px]': isOpenMenu, 'w-[70px]': !isOpenMenu}"
  >
    <MenuHead :isOpenMenu="isOpenMenu" :toggleMenu="toggleMenu"/>
    <MenuList :menuList="menuList"/>
  </div>
</template>

<script>
import {mapGetters, mapMutations} from "vuex";
import MenuHead from './SideMenuHead.vue'
import MenuList from './SideMenuList.vue'

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
