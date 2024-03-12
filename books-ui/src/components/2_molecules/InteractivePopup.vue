<script>


export default {
  name: 'InteractivePopup',
  components: {},
  data() {
    return {
      isOpen: false
    }
  },
  computed: {},
  mounted() {
    // Listen for clicks outside the dropdown
    document.addEventListener("click", this.handleClickOutside);
  },
  beforeUnmount() {
    // Remove click outside listener when component is unmounted
    document.removeEventListener("click", this.handleClickOutside);
  },
  methods: {
    toggle() {
      this.isOpen = !this.isOpen
    },
    close() {
      this.isOpen = false
    },
    handleStopPropagation(event) {
      // Prevent clicking inside the dropdown from closing it
      event.stopPropagation();
    },
    handleClickOutside(event) {
      const dropdown = this.$refs.popup;
      if (dropdown && !dropdown.contains(event.target)) {
        this.isOpen = false;
      }
    }
  }
}

</script>

<template>
  <div @click="handleStopPropagation">
    <div
      @click="toggle"
    >
      <slot name="target" />
    </div>
    <div
      v-if="isOpen"
      ref="popup"
    >
      <slot name="popup" />
    </div>
  </div>
</template>
