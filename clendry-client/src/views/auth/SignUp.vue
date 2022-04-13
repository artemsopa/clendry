<template>
  <section class="login-content">
         <div class="container h-100">
            <div class="row justify-content-center align-items-center height-self-center">
               <div class="col-md-5 col-sm-12 col-12 align-self-center">
                  <div class="sign-user_card">
                     <img src="@/assets/images/logo.png" class="img-fluid rounded-normal light-logo logo" alt="logo">
                     <h3 class="mb-3">Sign Up</h3>
                     <p>Create your account.</p>
                     <form @submit.prevent="onSubmit()">
                        <div class="row">
                           <div class="col-lg-12">
                              <div class="floating-label form-group">
                                 <input class="floating-input form-control" type="text" placeholder=" ">
                                 <label>Login</label>
                              </div>
                           </div>
                           <div class="col-lg-12">
                              <div class="floating-label form-group">
                                 <input class="floating-input form-control" type="email" placeholder=" ">
                                 <label>Email</label>
                              </div>
                           </div>
                           <div class="col-lg-6">
                              <div class="floating-label form-group">
                                 <input class="floating-input form-control" type="password" placeholder=" ">
                                 <label>Password</label>
                              </div>
                           </div>
                           <div class="col-lg-6">
                              <div class="floating-label form-group">
                                 <input class="floating-input form-control" type="password" placeholder=" ">
                                 <label>Confirm Password</label>
                              </div>
                           </div>
                           <div class="col-lg-12">
                              <div class="custom-control custom-checkbox mb-3 text-left">
                                 <input type="checkbox" class="custom-control-input" id="customCheck1">
                                 <label class="custom-control-label" for="customCheck1">I agree with the terms of use</label>
                              </div>
                           </div>
                        </div>
                        <button type="submit" class="btn btn-primary">Sign Up</button>
                        <p class="mt-3">
                           Already have an Account <router-link :to="{name: 'auth.login'}" class="text-primary">Sign In</router-link>
                        </p>
                     </form>
                  </div>
               </div>
            </div>
         </div>
      </section>
</template>

<script lang="ts">
import { defineComponent } from "@vue/runtime-core"

export default defineComponent({
    name:'SignUp',
    methods: {
    onSubmit () {
      this.$router.push({name: 'auth.login'})
    }
  }
})
</script>