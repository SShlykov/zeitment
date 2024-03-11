import 'tippy.js/dist/tippy.css'
import './styles/main.css'

import { createApp } from 'vue'
import App from './App.vue'
import ramdaVue from "./ramda-vue.js";
import { store } from '@store'
import VueTippy from 'vue-tippy'
import Router from "@router"

import moment from "moment/dist/moment"
import ru from "moment/dist/locale/ru"

moment.locale('ru', ru)

createApp(App)
  .use(Router)
  .use(ramdaVue)
  .use(store)
  .use(VueTippy)
  .mount('#app')
