import { App } from 'vue'

import BootstrapVue3 from 'bootstrap-vue-3'

export function bootstrapVue3Init(app: App<Element>) {
    app.use(BootstrapVue3)
}
