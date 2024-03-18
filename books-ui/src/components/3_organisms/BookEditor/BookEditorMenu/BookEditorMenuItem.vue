<script>
export default {
  name: 'BookEditorMenuItem',
  components: {},
  props: {
    title: {
      type: String,
      default: '123'
    }
  },
  computed: {},
  mounted() {},
  methods: {}
}

</script>

<template>
  <div class="w-[100px] transition-all w-full cursor-pointer rounded-md p-1 pl-2 hover:bg-gray-100">
    {{ title }}
  </div>
</template>
