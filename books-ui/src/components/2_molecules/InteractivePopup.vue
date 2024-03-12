<script>
import {ClickOutside} from '@directives/ClickOutside';

export default {
  name: 'InteractivePopup',
  directives: {
    'click-outside': ClickOutside
  },
  components: {},
  data() {
    return {
      isOpen: false
    }
  },
  computed: {},
  mounted() {

  },
  methods: {
    toggle() {
      this.isOpen = !this.isOpen
    },
    close() {
      console.log(123)
      this.isOpen = false
    }
  }
}

</script>

<template>
  <div
    v-click-outside="close"
  >
    <div
      @click="toggle"
    >
      <slot name="target" />
    </div>
    <div
      v-if="isOpen"

    >
      <slot name="popup" />
    </div>
  </div>
</template>
