import 'tippy.js/dist/tippy.css'
import './styles/main.css'

import { createApp } from 'vue'
import App from './App.vue'
import ramdaVue from "./ramda-vue.js";
import { store } from '@/store'
import VueTippy from 'vue-tippy'

createApp(App)
  .use(ramdaVue)
  .use(store)
  .use(VueTippy)
  .mount('#app')
