<template>
  <div class="h-full flex flex-col justify-between pb-4">
    <MenuItems
      :isOpenMenu="isOpenMenu"
      :menuList="topMenuList"
    />
    <MenuItems
      :isOpenMenu="isOpenMenu"
      :menuList="bottomMenuList"
    />
  </div>
</template>

<script>
import MenuItems from './MenuItems/MenuItems.vue'

export default {
  name: 'MenuList',
  components: {MenuItems},
  props: {
    menuList: {
      default: () => [],
      type: Array
    },
    isOpenMenu: {
      default: false,
      type: Boolean
    }
  },
  computed: {
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
