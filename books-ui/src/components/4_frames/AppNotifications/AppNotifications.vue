<script>

import AppNotificationsItem from "@frames/AppNotifications/AppNotificationsItem.vue";

export default {
  name: 'AppNotifications',
  components: {AppNotificationsItem},
  props: {
    notifications: {
      type: Array,
      default: () => []
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
  <div class="fixed flex flex-col gap-2 left-[50vw] translate-x-[-50%] top-[20px] z-40">
    <AppNotificationsItem
      v-for="{message, type} in notifications"
      :key="`${message}${type}`"
      :message="message"
      :type="type"
    />
  </div>
</template>
