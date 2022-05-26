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
import { computed, defineComponent } from "vue";
import Loader from '../components/loader/Loader.vue'
import Sidebar from "../components/partials/Sidebar.vue";
import Header from "../components/partials/Header.vue"
import Footer from "../components/partials/Footer.vue"
import { useStore } from "vuex";
import { useRouter } from "vue-router";

export default defineComponent({
  components: {
    Loader,
    Sidebar,
    Header,
    Footer,
  },
  setup() {
    const router = useRouter();
    const store = useStore();
    const auth = computed(() => store.state.authenticated);
    if (!localStorage.auth) {
      router.push('/auth/sign-in');
    }
    return {
      auth,
      router,
    }
  },
  mounted() {
    if (!localStorage.auth) {
      this.router.push('/auth/sign-in');
    }
  }
});
</script>
