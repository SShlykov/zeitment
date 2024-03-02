<script>
import Head from '@frames/Head.vue';
import SideMenu from '@frames/SideMenu/SideMenu.vue';
import {mapActions} from "vuex";



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
    pageName() {
      if (!this.$route) return "pending"
      return this.$route.name
    }
  },
  mounted()     {
    this.initScreenSizeRecalc()
  },
  methods: {
    ...mapActions('layout', ['initScreenSizeRecalc'])
  }
}
</script>


<template>
  <div
    v-if="pageName && pageName !== 'auth'"
    class="flex h-[99.9vh]"
  >
    <SideMenu />
    <div class="flex flex-col flex-grow">
      <Head v-if="false" />
      <div class="flex-grow relative">
        <router-view />
      </div>
    </div>
  </div>
  <div
    v-if="pageName == 'auth'"
    class="relative"
  >
    <router-view />
  </div>
</template>
