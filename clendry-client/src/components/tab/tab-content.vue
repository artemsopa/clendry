<template>
  <div :class="'tab-content' +' '+extraclass " :id="id"  >
    <slot />
  </div>
</template>

<script lang="ts">
import { defineComponent } from "@vue/runtime-core"


export default defineComponent({
  name: 'tab-content',
  props: {
    id: { type: String, default: 'myTabContent' },
    extraclass:{ type:String,default:"null"}
  }
})
</script>
