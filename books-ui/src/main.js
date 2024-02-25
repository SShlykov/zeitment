import 'tippy.js/dist/tippy.css'
import './styles/main.css'

import { createApp } from 'vue'
import App from './App.vue'
import ramdaVue from "./ramda-vue.js";
import { store } from '@store'
import VueTippy from 'vue-tippy'
import Router from "@router"

createApp(App)
  .use(Router)
  .use(ramdaVue)
  .use(store)
  .use(VueTippy)
  .mount('#app')
