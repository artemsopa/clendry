<template>
  <b-collapse
    tag="ul"
    :class="className"
    :visible="open"
    :id="idName"
    :accordion="accordianName"
  >
    <li
      v-for="(item, index) in items"
      :key="index"
      :class="
        activeLink(item) && item.children
          ? 'active'
          : activeLink(item)
          ? 'active'
          : ''
      "
    >
      <router-link
        :to="item.link.name"
        v-if="!item.is_heading"
        :class="`iq-waves-effect ${
          activeLink(item) && item.children
            ? 'active'
            : activeLink(item)
            ? 'active'
            : ''
        }`"
        data-toggle="collapse"
        :data-target="`#`+item.name"
        aria-expanded="false"
        :aria-controls="item.name"
      >
        <i
          v-if="item.name && hideListMenuTitle"
          :class="item.icon_class ? item.icon_class : ''"
        ></i>
        <template v-else v-html="item.icon"> </template>
        <span>{{ item.title }}</span>
        <i v-if="item.children" :class="item.expanded_icon_class" />
        <i v-if="item.children" :class="item.collapse_icon_class" />
        <small
          v-html="item.append"
          v-if="hideListMenuTitle"
          :class="item.append_class"
        />
      </router-link>
      <List
        :id="item.name"
        v-if="item.children"
        :items="item.children"
        class="iq-submenu collapse"
      />
    </li>
  </b-collapse>
</template>
<script lang="ts">
import { defineComponent } from "@vue/runtime-core";
import {core} from '@/config/pluginInit.js';

export default defineComponent({
  name: "List",
  props: {
    items: Array,
    className: { type: String, default: "iq-menu" },
    open: { type: Boolean, default: false },
    idName: { type: String, default: "sidebar" },
    accordianName: { type: String, default: "sidebar" },
    sidebarGroupTitle: { type: Boolean, default: true },
  },
  computed: {
    hideListMenuTitle() {
      return this.sidebarGroupTitle;
    },
  },
  methods: {
    activeLink(item: any) {
      return core.getActiveLink(item, this.$route.name);
    },
  },
});
</script>

<style lang="scss">
</style>
