<script>
import VSpinner from "@atoms/VSpinner.vue";

export default {
  name: 'ContentLoader',
  components: {VSpinner},
  props: {
    loaderSize: {
      type: Number,
      default: 20
    }
  },
  computed: {},
  mounted() {
  },

  methods: {
  }
}

</script>

<template>
  <div class="absolute flex justify-center items-center inset-0">
    <VSpinner
      class="mr-3"
      :size="loaderSize"
    />
    <slot />
  </div>
</template>
