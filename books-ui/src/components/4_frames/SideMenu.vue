<template>
  <div
    class="h-[100vh] z-40 bg-white flex flex-col px-4 pt-8 border-r border-slate-200 transition-all"
    :class="{'w-[300px]': isOpenMenu, 'w-[70px]': !isOpenMenu}"
  >
    <div class="pl-1 flex items-center mb-8 justify-between">
      <svg
        v-if="isOpenMenu"
        width="25"
        height="25"
        viewBox="0 0 100 100"
        fill="none"
        xmlns="http://www.w3.org/2000/svg"
      >
        <path
          fill-rule="evenodd"
          clip-rule="evenodd"
          d="M0 10C0 4.47716 4.47715 0 10 0H90C95.5228 0 100 4.47715 100 10V90C100 95.5228 95.5229 100 90 100H10C4.47716 100 0 95.5229 0 90V10ZM74 50C74 67.1208 63.2548 81 50 81C36.7452 81 26 67.1208 26 50C26 32.8792 36.7452 19 50 19C63.2548 19 74 32.8792 74 50ZM50 59C53.866 59 57 54.9706 57 50C57 45.0294 53.866 41 50 41C46.134 41 43 45.0294 43 50C43 54.9706 46.134 59 50 59Z"
          fill="#9B48DC"
        />
      </svg>
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
        v-for="{title, icon, link, type, name} in menu"
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
    const menu = [
      {
        title: "Главная",
        icon: "home-4-line",
        link: "/",
        type: "link",
        name: "main"
      },
      {
        title: "Воронка",
        icon: "line-chart-line",
        link: "/funnel",
        type: "link",
        name: "funnel"
      },
      {
        title: "Статистика",
        icon: "stack-fill",
        link: "/statistic",
        type: "link",
        name: "statistic"
      },
      {
        title: "Офферы",
        icon: "article-line",
        link: "/offers",
        type: "link",
        name: "offers"
      },
      {
        title: "Профиль",
        icon: "user-line",
        link: "/profile",
        type: "link",
        name: "profile"
      },
      {
        title: "Главная",
        icon: "home-4-line",
        link: "/",
        type: "line",
        name: ""
      },
      {
        title: "Форма",
        icon: "home-4-line",
        link: '/form',
        type: "link",
        name: 'form',
      },
      {
        title: "",
        icon: "home-4-line",
        link: "/",
        type: "line",
        name: "",
        mode: "dev"
      },
      {
        title: "Выход",
        icon: "logout-circle-line",
        link: "/auth",
        type: "link",
        name: ""
      },
    ].filter(({name, mode}) => {
      if (import.meta.env.VITE_APP_MODE === "dev") return true
      return !dev_routes.includes(name) && mode !== "dev"
    })
    return {
      menu
    }
  },
  computed: {
    ...mapGetters('layout', ['isOpenMenu']),
    pageName() {
      return this.$route.name
    }
  },
  methods: {
    ...mapMutations('layout', ['toggleMenu']),
  }
}

</script>
