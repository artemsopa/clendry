
import { App } from 'vue'

import { headerToggleInit } from './header-toggle'
import { menuToggleInit } from './menu-toggle'
import { navToggleInit } from './nav-toggle'
import { sidebarToggleInit } from './sidebar-toggle'

export function togglesInit(app: App<Element>) {
    headerToggleInit(app)
    menuToggleInit(app)
    navToggleInit(app)
    sidebarToggleInit(app)
}
