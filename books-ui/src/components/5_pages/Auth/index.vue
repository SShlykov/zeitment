<script>
import Login from './Login.vue'

export default {
  name: 'Auth',
  components: {
    Login
  },
  data() {
    return {}
  },
  computed: {
  },
  methods: {
  }
}

</script>

<template>
  <section class="bg-gray-50 dark:bg-gray-900">
    <div class="flex flex-col items-center justify-center px-6 py-8 mx-auto md:h-screen lg:py-0">
        <a href="#" class="flex items-center mb-6 text-2xl font-semibold text-gray-900 dark:text-white">
            <img class="w-8 h-8 mr-2" src="/icons/logo.svg" alt="logo">
            Книжный 
        </a>
        <div class="w-full bg-white p-6 rounded-lg shadow dark:border md:mt-0 sm:max-w-md  dark:bg-gray-800 dark:border-gray-700">
            <Login />
        </div>
    </div>
  </section>
</template>
