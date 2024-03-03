<template>
  <div class="h-full flex flex-col justify-between pb-4">
    <MenuItems :menuList="topMenuList" />
    <MenuItems :menuList="bottomMenuList" />
  </div>
</template>

<script>
import {mapGetters} from "vuex";
import MenuItems from './MenuItems/MenuItems.vue'

export default {
  name: 'MenuList',
  components: {MenuItems},
  computed: {
    ...mapGetters('layout', ['menuList']),
    topMenuList(){
      return this.menuList.filter(({position}) => position !== 'bottom')
    },
    bottomMenuList(){
      return this.menuList.filter(({position}) => position === 'bottom')
    },
  },
  methods: {
  }
}

</script>
