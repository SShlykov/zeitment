<script>
import AdapterOfBooks from '@/adapters/AdapterOfBooks'
import AdapterOfChapters from '@/adapters/AdapterOfChapters'
import AdapterOfPages from '@/adapters/AdapterOfPages'
import AdapterOfParagraphs from '@/adapters/AdapterOfParagraphs'

const startIntegrationTests = (adapter, logFunction) => {
  adapter.integrationTests(logFunction)
}

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
    logFunction(data) {
      this.testData = [
        ...this.testData,
        data
      ]
    },
    testBook() {
      const booksAdapter = new AdapterOfBooks(this.url)
      this.testData = []
      return startIntegrationTests(booksAdapter, this.logFunction)
    },
    testChapter() {
      const chaptersAdapter = new AdapterOfChapters(this.url)
      this.testData = []
      return startIntegrationTests(chaptersAdapter, this.logFunction)
    },
    testPage() {
      const chaptersAdapter = new AdapterOfPages(this.url)
      this.testData = []
      return startIntegrationTests(chaptersAdapter, this.logFunction)
    },
    testParagraph() {
      const chaptersAdapter = new AdapterOfParagraphs(this.url)
      this.testData = []
      return startIntegrationTests(chaptersAdapter, this.logFunction)
    }
  }
}

</script>

<template>
  <div class="p-4">
    <div class="mb-3">
      <span class="tx-gray-700 font-w-600 ">
        Запустить тесты интеграции
      </span>
    </div>
    <div class="flex items-center gap-2">
      <button
        type="button"
        class="text-white bg-purple-700 hover:bg-purple-800 focus:outline-none focus:ring-4 focus:ring-purple-300 font-medium rounded-full text-xs px-3 py-2 text-center mb-2 dark:bg-purple-600 dark:hover:bg-purple-700 dark:focus:ring-purple-900"
        @click="testBook"
      >
        Тесты книги
      </button>
      <button
        type="button"
        class="text-white bg-purple-700 hover:bg-purple-800 focus:outline-none focus:ring-4 focus:ring-purple-300 font-medium rounded-full text-xs px-3 py-2 text-center mb-2 dark:bg-purple-600 dark:hover:bg-purple-700 dark:focus:ring-purple-900"
        @click="testChapter"
      >
        Тесты глав
      </button>
      <button
        type="button"
        class="text-white bg-purple-700 hover:bg-purple-800 focus:outline-none focus:ring-4 focus:ring-purple-300 font-medium rounded-full text-xs px-3 py-2 text-center mb-2 dark:bg-purple-600 dark:hover:bg-purple-700 dark:focus:ring-purple-900"
        @click="testPage"
      >
        Тесты страниц
      </button>
      <button
        type="button"
        class="text-white bg-purple-700 hover:bg-purple-800 focus:outline-none focus:ring-4 focus:ring-purple-300 font-medium rounded-full text-xs px-3 py-2 text-center mb-2 dark:bg-purple-600 dark:hover:bg-purple-700 dark:focus:ring-purple-900"
        @click="testParagraph"
      >
        Тесты параграфов
      </button>
    </div>
    <pre>
<code>{{ testData }}
      </code>
    </pre>
  </div>
</template>
