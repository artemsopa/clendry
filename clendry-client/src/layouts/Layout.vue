<template>
  <div>
    <Loader />
    <div class="wrapper">
      <Sidebar />
      <Header />
      <div class="content-page">
        <transition name="router-anim">
          <router-view />
        </transition>
      </div>
    </div>
    <Footer />
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import Loader from '@/components/loader/Loader.vue'
import Sidebar from "@/components/partials/Sidebar.vue";
import Header from "@/components/partials/Header.vue"
import Footer from "@/components/partials/Footer.vue"

export default defineComponent({
  components: {
    Loader,
    Sidebar,
    Header,
    Footer,
  },
  setup() {},
});
</script>
