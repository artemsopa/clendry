<template>
  <ul :class="`nav nav-${pills ? 'pills' : ''}${tabs ? 'tabs' : ''} ${vertical ? 'flex-column' : ''} ${align ? 'justify-content-'+align : ''}`+' '+extraclass " :id="id" :role="role" :aria-orientation="aria">
    <slot />
  </ul>
</template>

<script lang="ts">
import { defineComponent } from '@vue/runtime-core'
const Tab = require('bootstrap/js/src/tab')

export default defineComponent({
  name: 'tab-nav',
  props: {
    id: { type: String, default: 'myTab' },
    pills: { type: Boolean, default: false },
    vertical: { type: Boolean, default: false },
    tabs: { type: Boolean, default: false },
    align: { type: String, default: '' },
    extraclass:{ type: String, default: '' },
    role:{ type: String, default: '' },
    aria:{ type: String, default: '' },
    
  },
  mounted () {
    var triggerTabList = [].slice.call(document.querySelectorAll(`#${this.id} a`))
    triggerTabList.forEach(function (triggerEl: any) {
      var tabTrigger = new Tab(triggerEl)

      triggerEl.addEventListener('click', function (e: any) {
        e.preventDefault()
        tabTrigger.show()
      })
    })
  }

})
</script>
