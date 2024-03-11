<script>
export default {
  name: 'BooksViewerContainer',
  components: {},
  data() {},
  computed: {},
  mounted() {},
  methods: {}
}

</script>

<template>
  <div class="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-4 2xl:grid-cols-6 3xl:grid-cols-7 gap-3">
    <slot />
  </div>
</template>
