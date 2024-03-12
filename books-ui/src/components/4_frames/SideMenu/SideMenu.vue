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

export default {
  name: 'SideMenu',
  components: {MenuHead, MenuList},
  data() {
  },
  computed: {
    ...mapGetters('layout', ['isOpenMenu', 'menuList']),
    ...mapGetters('books', ['userBooks']),
    flattenMenuList() {
      const newBook = {
        "title": "Создать книгу",
        "icon": "file-add-line",
        "itemFunction": async () => {
          const url = import.meta.env.VITE_API_ADDR
          const adapterOfBooks = new AdapterOfBooks(url)
          const store = this.$store

          const serviceOfBooks = new ServiceOfBooks(adapterOfBooks, store)
          const book = await serviceOfBooks.createBook()

          this.$router.push(`/book/${book.id}`)
        },
        "type": "button",
        "name": "new_book"
      }
      const booksMenu = booksListToMenuList(this.userBooks)
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
