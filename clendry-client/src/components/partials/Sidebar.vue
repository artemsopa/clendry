<template>
  <div class="iq-sidebar sidebar-default">
    <div
      class="iq-sidebar-logo d-flex align-items-center justify-content-between"
    >
      <router-link :to="{ name: '' }" class="header-logo">
        <img :src="logo" class="img-fluid rounded-normal" alt="logo" />
      </router-link>
      <div class="iq-menu-bt-sidebar">
        <i class="las la-bars wrapper-menu"></i>
      </div>
    </div>
    <div class="data-scrollbar" data-scroll="1" id="sidebar-scrollbar">
      <div class="new-create select-dropdown input-prepend input-append">
        <div class="btn-group">
          <label class="iq-user-toggle" id="dropdownMenuButton" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
            <div class="search-query selet-caption">
              <i class="las la-plus pr-2"></i>Create New
            </div>
            <span class="search-replace"></span>
            <span class="caret"><!--icon--></span>
          </label>
          <ul class="dropdown-menu" aria-labelledby="dropdownMenuButton">
            <li>
              <div class="item">
                <i class="ri-folder-add-line pr-3"></i>New Folder
              </div>
            </li>
            <li>
              <div class="item">
                <i class="ri-file-upload-line pr-3"></i>Upload Files
              </div>
            </li>
          </ul>
        </div>
      </div>
      <nav class="iq-sidebar-menu">
        <CollapseMenu v-bind:items="sidebarItems" :open="true" />
      </nav>
      <div class="sidebar-bottom">
        <h4 class="mb-3"><i class="las la-cloud mr-2"></i>Storage</h4>
        <p>17.1 / 20 GB Used</p>
        <Progressbar
          :value="67"
          color="primary"
          class="mb-3"
          midclass="iq-progress progress-1"
        />
        <p>75% Full - 3.9 GB Free</p>
        <a href="#" class="btn btn-outline-primary view-more mt-4"
          >Buy Storage</a
        >
      </div>
      <div class="p-3"></div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import CollapseMenu from "@/components/menustyle/CollapseMenu.vue";
import sideBarItem from "@/JsonData/sidebar.json";
import Progressbar from "@/components/progressbar/Progressbar.vue";
import { mapGetters } from "vuex";
import {core} from '@/config/pluginInit.js';

export default defineComponent({
  name: "Sidebar",
  setup() {
  },
  data() {
    return {
      homeurl: "",
      sidebarItems: sideBarItem, //<Array<MenuItem>>JSON.parse(sideBarItem.toString()),
    };
  },
  mounted() {
    core.SmoothScrollbar();
    core.changesidebar();
  },
  destroyed() {
    core.SmoothScrollbar();
    core.changesidebar();
  },
  components: {
    CollapseMenu,
    Progressbar,
  },
  computed: {
    ...mapGetters({
      appName: "appName",
      logo: "logo",
    }),
  },
});
</script>

<style lang="scss">
#dropdownMenuButton:hover {
    cursor: pointer;
}

.dropdown-menu {
    .item:hover {
        cursor: pointer;
    }
}
</style>
