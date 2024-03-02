<script>
import {mapGetters, mapMutations} from "vuex";

export default {
  name: 'MenuHead',
  components: {},
  data() {
  },
  computed: {
    ...mapGetters('layout', ['isOpenMenu']),
  },
  methods: {
    ...mapMutations('layout', ['toggleMenu']),
  }
}

</script>

<template>
  <div class="pl-1 flex items-center mb-8 justify-between">
    <div
      v-if="isOpenMenu"
      class="flex items-center"
    >
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
</template>
