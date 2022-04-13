import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import { togglesInit } from './directives/index'

const app = createApp(App)

togglesInit(app)

app.use(store).use(router).mount('#app')
