<script>
import BooksViewer from '@organisms/BooksViewer/BooksViewer.vue';
import AdapterOfBooks from "@adapters/AdapterOfBooks/AdapterOfBooks.js";
import ServiceOfBooks from "@services/ServiceOfBooks.js";
import {useStore} from "@store";

export default {
  name: 'UserBooks',
  components: {BooksViewer},
  setup() {
    const url = import.meta.env.VITE_API_ADDR
    const adapterOfBooks = new AdapterOfBooks(url)
    const store = useStore()

    const serviceOfBooks = new ServiceOfBooks(adapterOfBooks, store)
    return {
      serviceOfBooks
    }
  },
  computed: {},
  mounted() {},
  methods: {}
}

</script>

<template>
  <div class="p-6">
    <BooksViewer :serviceOfBooks="serviceOfBooks" />
  </div>
</template>
