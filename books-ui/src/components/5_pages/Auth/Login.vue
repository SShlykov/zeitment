<script>
import {mapGetters, mapActions} from "vuex";

export default {
  name: 'LogIn',
  components: {
  },
  data() {
    return {
      login: "",
      password: ""
    }
  },
  computed: {
    ...mapGetters('auth', ['errorMsg'])
  },
  mounted() {
    localStorage.removeItem('token')
  },
  unmounted() {
    this.login = ""
    this.password = ""
  },
  methods: {
    ...mapActions('auth', ['signin']),
  },
}

</script>

<template>
  <div class="space-y-4 md:space-y-6">
    <div>
      <label
        for="login"
        class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
      >Логин</label>
      <input
        id="login"
        v-model="login"
        type="login"
        name="login"
        class="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-blue-600 focus:border-blue-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
        placeholder="login"
        required=""
      >
    </div>
    <div>
      <label
        for="password"
        class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
      >Пароль</label>
      <input
        id="password"
        v-model="password"
        type="password"
        name="password"
        placeholder="••••••••"
        class="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-blue-600 focus:border-blue-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
        required=""
      >
    </div>
    <div
      v-if="errorMsg"
      class="text-red-600 text-sm"
    >
      {{ errorMsg }}
    </div>
    <button
      type="submit"
      class="w-full text-white bg-blue-600 hov er:bg-blue-700 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
      @click="signin({login, password})"
    >
      Войти
    </button>
  </div>
</template>
