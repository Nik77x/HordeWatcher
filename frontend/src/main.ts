import { createApp } from 'vue'
import App from './App.vue'

import './assets/main.scss'

/* import the fontawesome core */
import { library } from '@fortawesome/fontawesome-svg-core'

/* import font awesome icon component */
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

/* import specific icons */
import { faChevronDown, faGear, faGears, faUserSecret } from '@fortawesome/free-solid-svg-icons'

/* add icons to the library */
library.add(faUserSecret, faGear, faGears, faChevronDown)

createApp(App).component('font-awesome-icon', FontAwesomeIcon).mount('#app')
