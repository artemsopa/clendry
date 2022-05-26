<template>
    <div class="iq-top-navbar">
        <div class="iq-navbar-custom">
            <nav class="navbar navbar-expand-lg navbar-light p-0">
                <div class="iq-navbar-logo d-flex align-items-center justify-content-between">
                    <i class="ri-menu-line wrapper-menu"></i>
                    <router-link to="/" class="header-logo">
                        <img :src="logo" class="img-fluid rounded-normal" alt="logo">
                    </router-link>
                </div>
                <div class="iq-search-bar device-search">

                    <form>
                        <div class="input-prepend input-append">

                            <div class="btn-group">

                                <label class="dropdown-toggle searchbox" data-toggle="collapse" href="#collapseExample"
                                    role="button" aria-expanded="false" aria-controls="collapseExample">
                                    <input class="dropdown-toggle search-query text search-input " type="text"
                                        v-model="selectedInput" placeholder="Type here to search...">
                                    <span class="search-replace"></span>
                                    <a class="search-link" href="#"><i class="ri-search-line"></i></a>
                                    <!-- <span class="caret">icon</span> -->
                                </label>

                                <ul class="dropdown-menu collapse" id="collapseExample">
                                    <li @click="selectedInput = 'PDFs'"><a href="#">
                                            <div class="item"><i class="far fa-file-pdf bg-info"></i>PDFs</div>
                                        </a></li>
                                </ul>

                            </div>

                        </div>
                    </form>
                </div>

                <div class="d-flex align-items-center">
                    <ModeSwitch />
                    <div id="nav-collapse" is-nav>
                        <ul class="navbar-nav ml-auto navbar-list align-items-center">
                            <li class="nav-item nav-icon search-content">
                                <a href="#" class="search-toggle rounded dropdown-toggle" id="dropdownSearch"
                                    data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                                    <i class="ri-search-line"></i>
                                </a>
                                <div class="iq-search-bar iq-sub-dropdown dropdown-menu"
                                    aria-labelledby="dropdownSearch">
                                    <form action="#" class="searchbox p-2">
                                        <div class="form-group mb-0 position-relative">
                                            <input type="text" class="text search-input font-size-12"
                                                placeholder="type here to search...">
                                            <a href="#" class="search-link"><i class="las la-search"></i></a>
                                        </div>
                                    </form>
                                </div>
                            </li>
                            <li class="nav-item nav-icon dropdown">
                                <a href="#"
                                    class="search-toggle dropdown-toggle                                                                                                                                                                                                                                                                                                                                                                                          "
                                    id="dropdownMenuButton01" data-toggle="dropdown" aria-haspopup="true"
                                    aria-expanded="false">
                                    <i class="ri-question-line"></i>
                                </a>
                                <div class="iq-sub-dropdown dropdown-menu" aria-labelledby="dropdownMenuButton01">
                                    <div class="card shadow-none m-0">
                                        <div class="card-body p-0 ">
                                            <div class="p-3">
                                                <a href="#" class="iq-sub-card pt-0"><i
                                                        class="ri-questionnaire-line"></i>Help</a>
                                                <a href="#" class="iq-sub-card"><i
                                                        class="ri-recycle-line"></i>Training</a>
                                                <a href="#" class="iq-sub-card"><i
                                                        class="ri-refresh-line"></i>Updates</a>
                                                <a href="#" class="iq-sub-card"><i class="ri-service-line"></i>Terms and
                                                    Policy</a>
                                                <a href="#" class="iq-sub-card"><i class="ri-feedback-line"></i>Send
                                                    Feedback</a>
                                            </div>
                                        </div>
                                    </div>
                                </div>
                            </li>
                            <li class="nav-item nav-icon dropdown">
                                <a href="#" class="search-toggle dropdown-toggle  dropdown-toggle"
                                    id="dropdownMenuButton02" data-toggle="dropdown" aria-haspopup="true"
                                    aria-expanded="false">
                                    <i class="ri-settings-3-line"></i>
                                </a>
                                <div class="iq-sub-dropdown dropdown-menu" aria-labelledby="dropdownMenuButton02">
                                    <div class="card shadow-none m-0">
                                        <div class="card-body p-0 ">
                                            <div class="p-3">
                                                <a href="#" class="iq-sub-card pt-0"><i class="ri-settings-3-line"></i>
                                                    Settings</a>
                                                <a href="#" class="iq-sub-card"><i class="ri-hard-drive-line"></i> Get
                                                    Drive for desktop</a>
                                                <a href="#" class="iq-sub-card"><i class="ri-keyboard-line"></i>
                                                    Keyboard Shortcuts</a>
                                            </div>
                                        </div>
                                    </div>
                                </div>
                            </li>
                            <li class="nav-item nav-icon dropdown caption-content">
                                <a href="#" class="search-toggle dropdown-toggle" id="dropdownMenuButton03"
                                    data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                                    <div class="caption bg-primary line-height">
                                        {{ avatar }}
                                    </div>
                                </a>
                                <div class="iq-sub-dropdown dropdown-menu" aria-labelledby="dropdownMenuButton03">
                                    <div class="card mb-0">
                                        <div class="card-header d-flex justify-content-between align-items-center mb-0">
                                            <div class="header-title">
                                                <h4 class="card-title mb-0">Profile</h4>
                                            </div>
                                        </div>
                                        <div class="card-body">
                                            <div class="profile-header">
                                                <div class="cover-container text-center">
                                                    <div class="rounded-circle profile-icon bg-primary mx-auto d-block">
                                                        {{ avatar }}
                                                    </div>
                                                    <div class="profile-detail mt-3">
                                                        <h5>{{ user ? user.nick : "Hi" }}</h5>
                                                        <p>{{ user ? user.email : "Hi" }}</p>
                                                        <hr>
                                                    </div>
                                                    <router-link :to="{ name: 'auth.login' }" @click="logout"
                                                        class="btn btn-primary">
                                                        Sign Out</router-link>
                                                </div>
                                            </div>
                                        </div>
                                    </div>
                                </div>
                            </li>
                        </ul>
                    </div>
                </div>
            </nav>
        </div>
    </div>

</template>
<script lang="ts">


import axios from "axios";
import { computed, defineComponent } from "vue";
import { useRouter } from "vue-router";
import { useStore } from "vuex";
import { mapGetters } from 'vuex';
import User from '../../models/user'

export default defineComponent({
    name: "Header",
    data() {
        return {
            selectedInput: '',
            user: {} as User,
            avatar: ''
        }
    },
    setup() {
        const store = useStore();
        const auth = computed(() => store.state.authenticated);
        const router = useRouter();
        const logout = async () => {
            await axios.post(
                "auth/logout",
                {},
                {
                    withCredentials: true,
                }
            );
            await store.dispatch("setAuth", false);
            await router.push("/auth/sign-in");
        };
        return {
            auth,
            router,
            logout,
        };
    },
    methods: {
        async getUser() {
            await axios.get<User>(`/profile`, {
                withCredentials: true
            }).then(response => {
                this.user = response.data;
                this.avatar = this.user.nick.split('')[0].toUpperCase()
            }).catch(() => this.router.push("/auth/sign-in"));
        }
    },
    computed: {
        ...mapGetters({
            appName: 'appName',
            logo: 'logo',

            name: 'name',
            image1: 'image1'
        })
    },
    mounted() {
        this.getUser();
    }
})
</script>
