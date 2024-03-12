<script>
import { directive } from 'vue-tippy'
import InteractivePopup from "@molecules/InteractivePopup.vue";

export default {
  name: 'ItemBook',
  components: {InteractivePopup},
  directives: {
    tippy: directive
  },
  props: {
    icon: {
      default: "",
      type: String
    },
    title: {
      default: "",
      type: String
    },
    link: {
      default: "",
      type: String
    },
    name: {
      default: "",
      type: String
    },
    isOpenMenu: {
      type: Boolean,
      default: false
    }
  },
  computed: {
    pageName() {
      if (!this.$route) return ""
      return this.$route.name
    },
    bookAction() {
      const layout = `

      `
      return { content: layout, allowHTML: true, interactive: true, trigger: 'click', placement: 'bottom-start', arrow: false}
    }
  },
  methods: {
    test: () => {
      console.log(1234)
    }
  }
}

</script>


<template>
  <div class="w-full h-full relative flex p-2 rounded-md transition-all hover:bg-gray-200 items-center group">
    <router-link
      class="w-full"
      :class="{'bg-gray-200': pageName === name}"
      :to="link"
    >
      <i :class="`ri-${icon} text-slate-600 mr-2 `" />
      <span
        v-if="isOpenMenu"
        class="text-slate-700 truncate"
      >{{ title }}</span>
    </router-link>
    <InteractivePopup class="z-20">
      <template #target>
        <div
          class="absolute right-[10px] top-[50%] -translate-y-1/2 transition-all opacity-0 group-hover:opacity-100 hover:bg-gray-300 cursor-pointer h-6 w-6 rounded-md flex items-center justify-center"
        >
          <i class="ri-more-2-line transition-all" />
        </div>
      </template>
      <template #popup>
        <div class="flex z-10 flex-col absolute top-[100%] right-[-30%] bg-white p-2 rounded-md shadow-md border border-gray-100">
          <div
            class="transition-all flex text-sky-800 text-sm p-2 cursor-pointer rounded-md hover:bg-slate-200"
            @click="test"
          >
            <i class="ri-delete-bin-line mr-2" />
            Удалить
          </div>
        </div>
      </template>
    </InteractivePopup>
  </div>
</template>

