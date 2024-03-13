<script>
import InteractivePopup from "@molecules/InteractivePopup.vue";

export default {
  name: 'ItemBook',
  components: {InteractivePopup},
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
    },
    serviceOfBooks: {
      type: Object,
      default: () => {}
    }
  },
  data() {
    return {
      isRemoving: false
    }
  },
  computed: {
    id() {
      return this.name
    },
    pageName() {
      if (!this.$route) return ""
      return this.$route.params.id
    },
  },
  methods: {
    removeBook(id)  {
      this.serviceOfBooks.removeBook(id)
    },
    toggleRemoving() {
      this.isRemoving = !this.isRemoving
    }
  }
}

</script>


<template>
  <div
    class="w-full h-full relative flex p-2 rounded-md transition-all hover:bg-gray-200 items-center group"
    :class="{'bg-gray-200': pageName === name}"
  >
    <router-link
      class="w-full"
      :to="link"
    >
      <i :class="`ri-${icon} text-slate-600 mr-2 `" />
      <span
        v-if="isOpenMenu"
        class="text-slate-700 truncate"
      >{{ title }}</span>
    </router-link>
    <InteractivePopup>
      <template #target>
        <div
          class="absolute z-10 right-[10px] top-[50%] -translate-y-1/2 transition-all opacity-0 group-hover:opacity-100 hover:bg-gray-300 cursor-pointer h-6 w-6 rounded-md flex items-center justify-center"
        >
          <i class="ri-more-2-line transition-all" />
        </div>
      </template>
      <template #popup>
        <div class="flex z-20 flex-col absolute top-[100%] right-[-30%] bg-white p-2 rounded-md shadow-md border border-gray-100">
          <div class="flex items-center">
            <div
              v-if="!isRemoving"
              class="transition-all flex text-sky-800 text-sm p-2 cursor-pointer rounded-md hover:bg-slate-200"
              @click="toggleRemoving"
            >
              <i class="ri-delete-bin-line mr-2" />
              Удалить
            </div>
            <div
              v-if="isRemoving"
              class="transition-all flex text-sky-800 text-sm p-2 cursor-pointer rounded-md text-white bg-red-700 hover:bg-red-800 mr-2"
              @click="removeBook(id)"
            >
              Удалить
            </div>
            <div
              v-if="isRemoving"
              class="transition-all flex text-sky-800 text-sm p-2 cursor-pointer rounded-md text-white bg-gray-700 hover:bg-gray-800"
              @click="toggleRemoving"
            >
              Отмена
            </div>
          </div>
        </div>
      </template>
    </InteractivePopup>
  </div>
</template>

