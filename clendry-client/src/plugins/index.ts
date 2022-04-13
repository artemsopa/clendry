import { App } from 'vue'

import { globalInit } from './global'
import { flatPickrInit } from './flatpicker'
import { bootstrapVue3Init } from './bootstrap-vue'
import "./jvector"
import "./masonry"
import "./magnific-popup"
import "./datatable"
import "./select2"

export function pluginsInit(app: App<Element>) {
    globalInit(app)
    flatPickrInit(app)
    bootstrapVue3Init(app)
}
