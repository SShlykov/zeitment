<script>
import Head from '@frames/Head.vue';
import SideMenu from '@frames/SideMenu/SideMenu.vue';
import {mapGetters} from "vuex";
import AdapterOfBooks from "@adapters/AdapterOfBooks.js";
import ServiceOfBooks from "@services/ServiceOfBooks.js";


export default {
  name: 'App',
  components: {
    Head,
    SideMenu,
  },
  data() {
    return {}
  },
  computed: {
    ...mapGetters('books', ['userBooks']),
    pageName() {
      if (!this.$route) return "pending"
      return this.$route.name
    }
  },
  mounted() {
    const url = import.meta.env.VITE_API_ADDR
    const adapterOfBooks = new AdapterOfBooks(url)
    const store = this.$store

    const serviceOfBooks = new ServiceOfBooks(adapterOfBooks, store)
    serviceOfBooks.fetchUserBooks()

  },
}
</script>

<template>
  <div
    v-if="pageName && pageName !== 'auth'"
    class="flex h-[99.9vh] "
  >
    <SideMenu />
    <div class="flex flex-col flex-grow">
      <Head v-if="false" />
      <div class="flex-grow relative overflow-auto">
        <router-view />
      </div>
    </div>
  </div>
  <div
    v-if="pageName == 'auth'"
    class="relative "
  >
    <router-view />
  </div>
</template>
