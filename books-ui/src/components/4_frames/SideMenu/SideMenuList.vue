<template>
  <div class="h-full flex flex-col justify-between pb-4">
    <MenuItems
      :isOpenMenu="isOpenMenu"
      :menuList="topMenuList"
      :serviceOfBooks="serviceOfBooks"
    />
    <MenuItems
      :isOpenMenu="isOpenMenu"
      :menuList="bottomMenuList"
      :serviceOfBooks="serviceOfBooks"
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
    },
    serviceOfBooks: {
      default: () => {},
      type: Object
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
