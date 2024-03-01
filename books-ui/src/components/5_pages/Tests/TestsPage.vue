<script>
import { mapGetters, mapActions } from 'vuex'

export default {
  name: 'TestsPage',
  components: {},
  data() {},
  computed: {
    ...mapGetters('test', ['integrationTestLog'])
  },
  mounted() {},
  methods: {
    ...mapActions('test', ['startIntegrationTest'])
  }
}

</script>

<template>
  <div class="p-4">
    <div>
      <button
        type="button"
        class="text-white bg-purple-700 hover:bg-purple-800 focus:outline-none focus:ring-4 focus:ring-purple-300 font-medium rounded-full text-xs px-3 py-2 text-center mb-2 dark:bg-purple-600 dark:hover:bg-purple-700 dark:focus:ring-purple-900"
        @click="startIntegrationTest"
      >
        Запустить тесты интеграции
      </button>
    </div>
    <pre>
<code>{{ integrationTestLog }}
      </code>
    </pre>
  </div>
</template>
