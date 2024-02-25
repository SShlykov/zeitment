<template>
  <div
    class="h-[100vh] z-40 bg-white flex flex-col px-4 pt-8 border-r border-slate-200 transition-all"
    :class="{'w-[300px]': isOpenMenu, 'w-[70px]': !isOpenMenu}"
  >
    <div class="pl-1 flex items-center mb-8 justify-between">
      <div class="flex items-center" v-if="isOpenMenu">
        <img
          class="mr-2"
          height="25"
          width="25"
          src="/icon.svg"
          alt="logo"
        >
        <div>
          Буглов В. А.
        </div>
      </div>
      <div
        class="p-1 px-2 rounded-md transition-all cursor-pointer hover:bg-gray-200"
        :class="{'rotate-90': !isOpenMenu, '-rotate-90': isOpenMenu}"
        @click="toggleMenu"
      >
        <i class="ri-upload-line" />
      </div>
    </div>
    <div class="flex flex-col">
      <div
        v-for="{title, icon, link, type, name} in menuList"
        :key="`${title}${name}`"
      >
        <router-link
          v-if="type === 'link'"
          class="flex p-2 rounded-md hover:bg-gray-100 flex items-center"
          :class="{'bg-gray-200': pageName === name}"
          :to="link"
        >
          <i :class="`ri-${icon} text-slate-600 mr-2 `" />
          <span
            v-if="isOpenMenu"
            class="text-slate-700"
          >{{ title }}</span>
        </router-link>
        <div
          v-if="type === 'line'"
          class="border-b my-4 border-slate-200"
        />
      </div>
    </div>
  </div>
</template>

<script>
import {mapGetters, mapMutations} from "vuex";
import {dev_routes} from "@router"

export default {
  name: 'App',
  components: {},
  data() {
  },
  computed: {
    ...mapGetters('layout', ['isOpenMenu', 'menuList']),
    pageName() {
      return this.$route.name
    }
  },
  methods: {
    ...mapMutations('layout', ['toggleMenu']),
  }
}

</script>
