<template>
   <section class="login-content">
      <div class="container h-100">
         <div class="row justify-content-center align-items-center height-self-center">
            <div class="col-md-5 col-sm-12 col-12 align-self-center">
               <div class="sign-user_card">
                  <img src="@/assets/images/logo.png" class="img-fluid rounded-normal light-logo logo" alt="logo">

                  <h3 class="mb-3">Sign In</h3>
                  <p>Login to stay connected.</p>
                  <form  @submit.prevent="submit">
                     <div class="row">
                        <div class="col-lg-12">
                           <div class="floating-label form-group">
                              <input v-model="input.login" class="floating-input form-control" type="text"
                                 placeholder=" ">
                              <label>Login</label>
                           </div>
                        </div>
                        <div class="col-lg-12">
                           <div class="floating-label form-group">
                              <input v-model="input.password" class="floating-input form-control" type="password"
                                 placeholder=" ">
                              <label>Password</label>
                           </div>
                        </div>
                        <div class="col-lg-6">
                           <div class="custom-control custom-checkbox mb-3 text-left">
                              <input type="checkbox" class="custom-control-input" id="customCheck1">
                              <label class="custom-control-label" for="customCheck1">Remember Me</label>
                           </div>
                        </div>
                     </div>
                     <button type="submit" class="btn btn-primary">Sign In</button>
                     <p class="mt-3">
                        Create an Account <router-link :to="{ name: 'auth.register' }" class="text-primary">Sign Up
                        </router-link>
                     </p>
                  </form>
               </div>
            </div>
         </div>
      </div>
   </section>
</template>

<script lang="ts">
import axios from "axios";
import { computed, defineComponent, reactive } from "vue";
import { useRouter } from "vue-router";
import { useStore } from "vuex";

export default defineComponent({
   name: 'SignIn',
   setup() {
      const store = useStore();
      const auth = computed(() => store.state.authenticated);
      const router = useRouter();

      const input = reactive({
         login: "",
         password: "",
      });
      const submit = async () => {
         try {
            await axios.post("auth/sign-in", input, {
               withCredentials: true,
            });
            await store.dispatch("setAuth", true);
            localStorage.auth = true;
            await router.push("/");
         } catch (e) {
            console.log(e);
         }
      };
      return {
         input,
         submit,
      };
   },
   methods: {
      onSubmit() {
         this.$router.push({ name: 'layout.dashboard' })
      }
   }
})
</script>