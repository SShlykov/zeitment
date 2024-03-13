<script>

export default {
  name: 'AppNotificationsItem',
  components: {},
  props: {
    message: {
      type: String,
      default: ''
    },
    type: {
      type: String,
      default: 'info' // 'info' | 'success' | 'warning' | 'error'
    }
  },
  computed: {
    icon() {
      let iconClass
      switch (this.type) {
      case 'info':
        iconClass = 'text-blue-500 ri-information-2-line'
        break;
      case 'success':
        iconClass = 'text-green-500 ri-checkbox-circle-line'
        break;
      case 'warning':
        iconClass = 'text-yellow-500 ri-alert-line'
        break;
      case 'error':
        iconClass = 'text-red-500 ri-spam-2-line'
        break;
      }
      return iconClass
    }
  },
  mounted() {
  },
  methods: {
  }
}

</script>

<template>
  <div class="w-[70vw] flex items-center max-w-[400px] min-w-[300px] bg-white shadow-md border border-gray-200 p-2 rounded-md appearFromTop">
    <i
      class="mr-2"
      :class="icon"
    />
    <span class="text-gray-500">{{ message }}</span>
  </div>
</template>
