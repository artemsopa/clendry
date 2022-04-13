import { App } from 'vue'

export function navToggleInit(app: App<Element>) {
  app.directive('nav-toggle', {
    beforeMount: function (el, binding) {
      function doClose(e: any) {
        if (!isOpen) return
        if (e === undefined) {
          isOpen = false
          el.querySelector('.dropdown-menu').classList.remove('show')
          document.removeEventListener('click', onClose, false)
        }
      }
      function onClose(e: any) {
        if (e && el.contains(e.target)) return
        doClose(e)
      }
      function onOpen() {
        if (isOpen) return
        isOpen = true
        el.querySelector('.dropdown-menu').classList.add('show')
        document.addEventListener('click', onClose, false)
      }
      function onToggle(e: any) {
        isOpen ? onClose(e) : onOpen()
      }
      let isOpen = false
      const toggle = el.querySelector('.dropdown-toggle')
      const { value } = binding
      const { click = 'toggle' } = value || {}
      if (click === 'toggle') {
        if (toggle !== null) {
          toggle.addEventListener('click', onToggle, false)
        }
      } else if (click === 'open') {
        if (toggle !== null) {
          toggle.addEventListener('click', onOpen, false)
        }
      }
    }
  })
}