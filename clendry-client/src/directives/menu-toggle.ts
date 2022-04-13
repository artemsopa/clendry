import { App } from 'vue'

export function menuToggleInit(app: App<Element>) {
  app.directive('menu-toggle', {
    beforeMount: function (el, binding) {
      function doClose(e: any) {
        if (!isOpen) return
        if (e === undefined) {
          isOpen = false
          el.querySelector('.iq-dropdown').classList.remove('show')
          el.querySelector('.iq-toggle').classList.add('collapsed')
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
        el.querySelector('.iq-dropdown').classList.add('show')
        el.querySelector('.iq-toggle').classList.remove('collapsed')
        document.addEventListener('click', onClose, false)
      }
      function onToggle(e: any) {
        isOpen ? onClose(e) : onOpen()
      }
      let isOpen = false
      const toggle = el.querySelector('.iq-toggle')
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