<script>
import AdapterOfBooks from '@/adapters/AdapterOfBooks'

export default {
  name: 'TestsPage',
  components: {},
  data() {
    return {
      testData: [],
      url: import.meta.env.VITE_API_ADDR
    }
  },
  computed: {
  },
  mounted() {},
  methods: {
    startIntegrationTest() {
      const adapterOfBooks = new AdapterOfBooks(this.url)
      adapterOfBooks.integrationTests(
        (data) => {
          this.testData = [
            ...this.testData,
            data
          ]
        }
      )
    }
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
<code>{{ testData }}
      </code>
    </pre>
  </div>
</template>
