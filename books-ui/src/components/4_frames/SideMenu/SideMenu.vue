<template>
  <div
    class="h-[100vh] z-40  flex flex-col px-4 pt-8 border-r border-slate-200 transition-all bg-gray-100"
    :class="{'w-[300px]': isOpenMenu, 'w-[70px]': !isOpenMenu}"
  >
    <MenuHead
      :isOpenMenu="isOpenMenu"
      :toggleMenu="toggleMenu"
    />
    <MenuList
      :menuList="flattenMenuList"
      :isOpenMenu="isOpenMenu"
      :serviceOfBooks="serviceOfBooks"
    />
  </div>
</template>

<script>
import {mapGetters, mapMutations} from "vuex";
import MenuHead from './SideMenuHead.vue'
import MenuList from './SideMenuList.vue'
import {booksListToMenuList} from '@helpers/menuFuncs'
import AdapterOfBooks from "@adapters/AdapterOfBooks.js";
import ServiceOfBooks from "@services/ServiceOfBooks.js";
import {useStore} from "@store";

export default {
  name: 'SideMenu',
  components: {MenuHead, MenuList},
  setup() {
    const url = import.meta.env.VITE_API_ADDR
    const adapterOfBooks = new AdapterOfBooks(url)
    const store = useStore()

    const serviceOfBooks = new ServiceOfBooks(adapterOfBooks, store)
    return {
      serviceOfBooks
    }
  },
  computed: {
    ...mapGetters('layout', ['isOpenMenu', 'menuList']),
    ...mapGetters('books', ['userBooks']),
    flattenMenuList() {
      const newBook = {
        "title": "Создать книгу",
        "icon": "file-add-line",
        "type": "button",
        "name": "new_book"
      }
      const booksMenu = booksListToMenuList(this.userBooks)
      return [newBook, ...booksMenu, ...this.menuList]
    },
  },
  methods: {
    ...mapMutations('layout', ['toggleMenu']),
  }
}

</script>
