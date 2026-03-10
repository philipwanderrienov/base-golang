import './assets/main.css'

import { createApp } from 'vue'
import App from './App.vue'
import router from './router'

import PrimeVue from 'primevue/config'
// import Aura from '@primevue/themes/aura';
import Aura from '@primeuix/themes/aura'
import './assets/main.css'

const app = createApp(App)

app.use(router)
app.use(PrimeVue, {
  theme: {
    preset: Aura,
    options: {
      // darkModeSelector: '.my-app-dark', // Easier dark mode handling in v4
      cssLayer: {
        name: 'primevue',
        order: 'base, primevue, utilities',
      },
    },
  },
})

app.mount('#app')
