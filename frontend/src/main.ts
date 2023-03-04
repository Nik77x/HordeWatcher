import { createApp } from 'vue'
import App from './App.vue'
import './assets/main.scss'
import { library } from '@fortawesome/fontawesome-svg-core'

import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

import {
  faBarsStaggered,
  faChevronDown,
  faGear,
  faGears,
  faUserSecret
} from '@fortawesome/free-solid-svg-icons'
import { createPinia } from 'pinia'

library.add(faUserSecret, faGear, faGears, faChevronDown, faBarsStaggered)

const pinia = createPinia()

const app = createApp(App).component('font-awesome-icon', FontAwesomeIcon).use(pinia).mount('#app')
