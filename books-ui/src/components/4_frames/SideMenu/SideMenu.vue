<template>
  <div
    class="h-[100vh] z-40  flex flex-col px-4 pt-8 border-r border-slate-200 transition-all bg-gray-100"
    :class="{'w-[300px]': isOpenMenu, 'w-[70px]': !isOpenMenu}"
  >
    <MenuHead
      :isOpenMenu="isOpenMenu"
      :toggleMenu="toggleMenu"
    />
    <MenuList :menuList="flattenMenuList" />
  </div>
</template>

<script>
import {mapGetters, mapMutations} from "vuex";
import MenuHead from './SideMenuHead.vue'
import MenuList from './SideMenuList.vue'
import {booksListToMenuList} from '@helpers/menuFuncs'

export default {
  name: 'SideMenu',
  components: {MenuHead, MenuList},
  data() {
  },
  computed: {
    ...mapGetters('layout', ['isOpenMenu', 'menuList']),
    ...mapGetters('userBooks', ['booksList']),
    flattenMenuList() {
      const newBook = {
        "title": "Создать книгу",
        "icon": "file-add-line",
        "link": "/new_book",
        "type": "link",
        "name": "new_book"
      }
      const booksMenu = booksListToMenuList(this.booksList)
      return [newBook, ...booksMenu, ...this.menuList]
    },
    pageName() {
      return this.$route.name
    }
  },
  methods: {
    ...mapMutations('layout', ['toggleMenu']),
  }
}

</script>
