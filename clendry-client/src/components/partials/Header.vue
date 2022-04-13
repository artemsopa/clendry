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
                  
                        <div class="btn-group" >
                           
                            <label class="dropdown-toggle searchbox" data-toggle="collapse" href="#collapseExample" role="button" aria-expanded="false" aria-controls="collapseExample">
                            <input class="dropdown-toggle search-query text search-input " type="text" v-model="selectedInput" placeholder="Type here to search..." >
                            <span class="search-replace"></span>
                            <a class="search-link" href="#"><i class="ri-search-line"></i></a>
                            <!-- <span class="caret">icon</span> -->
                            </label>
                              
                            <ul class="dropdown-menu collapse" id="collapseExample" >
                                <li @click="selectedInput = 'PDFs'"><a href="#"><div class="item"><i class="far fa-file-pdf bg-info"></i>PDFs</div></a></li>
                                <li @click="selectedInput = 'Documents'"><a href="#"><div class="item"><i class="far fa-file-alt bg-primary"></i>Documents</div></a></li>
                                <li @click="selectedInput = 'Spreadsheet'"><a href="#"><div class="item"><i class="far fa-file-excel bg-success"></i>Spreadsheet</div></a></li>
                                <li @click="selectedInput = 'Presentation'"><a href="#"><div class="item"><i class="far fa-file-powerpoint bg-danger"></i>Presentation</div></a></li>
                                <li @click="selectedInput = 'Photos & Images'"><a href="#"><div class="item"><i class="far fa-file-image bg-warning"></i>Photos & Images</div></a></li>
                                <li @click="selectedInput = 'Videos'"><a href="#"><div class="item"><i class="far fa-file-video bg-info"></i>Videos</div></a></li>
                            </ul>
                              
                        </div>
                        
                    </div>
                </form>
            </div>

            <div class="d-flex align-items-center">
                <b-navbar-toggle target="nav-collapse" class="mr-2 text-primary">
                <i class="ri-menu-3-line"></i>
                </b-navbar-toggle>
                <b-collapse id="nav-collapse" is-nav>
                    <ul class="navbar-nav ml-auto navbar-list align-items-center">
                    <li class="nav-item nav-icon search-content">
                        <a href="#" class="search-toggle rounded dropdown-toggle" id="dropdownSearch" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                            <i class="ri-search-line"></i>
                        </a>
                        <div class="iq-search-bar iq-sub-dropdown dropdown-menu" aria-labelledby="dropdownSearch">
                            <form action="#" class="searchbox p-2">
                                <div class="form-group mb-0 position-relative">
                                <input type="text" class="text search-input font-size-12" placeholder="type here to search...">
                                <a href="#" class="search-link"><i class="las la-search"></i></a> 
                                </div>
                            </form>
                        </div>
                    </li>                
                     <li class="nav-item nav-icon dropdown" >
                        <a href="#" class="search-toggle dropdown-toggle                                                                                                                                                                                                                                                                                                                                                                                          " id="dropdownMenuButton01" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                        <i class="ri-question-line"></i>
                        </a>
                        <div class="iq-sub-dropdown dropdown-menu" aria-labelledby="dropdownMenuButton01">
                            <div class="card shadow-none m-0">
                                <div class="card-body p-0 ">
                                    <div class="p-3">
                                        <a href="#" class="iq-sub-card pt-0"><i class="ri-questionnaire-line"></i>Help</a>
                                        <a href="#" class="iq-sub-card"><i class="ri-recycle-line"></i>Training</a>
                                        <a href="#" class="iq-sub-card"><i class="ri-refresh-line"></i>Updates</a>
                                        <a href="#" class="iq-sub-card"><i class="ri-service-line"></i>Terms and Policy</a>
                                        <a href="#" class="iq-sub-card"><i class="ri-feedback-line"></i>Send Feedback</a>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </li>
                    <li class="nav-item nav-icon dropdown"  > 
                        <a href="#" class="search-toggle dropdown-toggle  dropdown-toggle" id="dropdownMenuButton02" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false" >
                        <i class="ri-settings-3-line"></i>
                        </a>
                        <div class="iq-sub-dropdown dropdown-menu" aria-labelledby="dropdownMenuButton02">
                            <div class="card shadow-none m-0">
                                <div class="card-body p-0 ">
                                    <div class="p-3">
                                        <a href="#" class="iq-sub-card pt-0"><i class="ri-settings-3-line"></i> Settings</a>
                                        <a href="#" class="iq-sub-card"><i class="ri-hard-drive-line"></i> Get Drive for desktop</a>
                                        <a href="#" class="iq-sub-card"><i class="ri-keyboard-line"></i> Keyboard Shortcuts</a>
                                    </div>                                
                                </div>
                            </div>
                        </div>
                    </li>
                     <li class="nav-item nav-icon dropdown caption-content"  >
                        <a href="#" class="search-toggle dropdown-toggle" id="dropdownMenuButton03" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false" >
                            <div class="caption bg-primary line-height">P</div>
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
                                                P                                                    
                                                <a href="">
                                                    
                                                </a>
                                            </div>
                                            <div class="profile-detail mt-3">
                                            <h5><a href="@/app/user-profile-edit.html">Panny Marco</a></h5>
                                            <p>pannymarco@gmail.com</p>
                                            </div>
                                            <router-link :to="{ name: 'auth.login'}" class="btn btn-primary">Sign Out</router-link>
                                        </div>
                                        <div class="profile-details mt-4 pt-4 border-top">
                                            <div class="media align-items-center mb-3">
                                                <div class="rounded-circle iq-card-icon-small bg-primary">
                                                    A
                                                </div>
                                                <div class="media-body ml-3">
                                                    <div class="media justify-content-between">
                                                        <h6 class="mb-0">Anna Mull</h6>
                                                        <p class="mb-0 font-size-12"><i>Signed Out</i></p>
                                                    </div>
                                                    <p class="mb-0 font-size-12">annamull@gmail.com</p>
                                                </div>                                                 
                                            </div>
                                            <div class="media align-items-center mb-3">
                                                <div class="rounded-circle iq-card-icon-small bg-success">
                                                    K
                                                </div>
                                                <div class="media-body ml-3">
                                                    <div class="media justify-content-between">
                                                        <h6 class="mb-0">King Poilin</h6>
                                                        <p class="mb-0 font-size-12"><i>Signed Out</i></p>
                                                    </div>
                                                    <p class="mb-0 font-size-12">kingpoilin@gmail.com</p>
                                                </div>
                                            </div>
                                            <div class="media align-items-center">
                                                <div class="rounded-circle iq-card-icon-small bg-danger">
                                                    D
                                                </div>
                                                <div class="media-body ml-3">
                                                    <div class="media justify-content-between">
                                                        <h6 class="mb-0">Devid Worner</h6>
                                                        <p class="mb-0 font-size-12"><i>Signed Out</i></p>
                                                    </div>
                                                    <p class="mb-0 font-size-12">devidworner@gmail.com</p>
                                                </div>
                                            </div>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </li>
                    </ul>                     
                </b-collapse> 
            </div>
        </nav>
    </div>
</div>

</template>
<script>

import { defineComponent } from '@vue/runtime-core'
import { mapGetters } from 'vuex'

export default defineComponent({
    name:"Header",
    data(){
        return{
            selectedInput: ''
        }
    },
    computed : {
        ...mapGetters({
            appName: 'appName',
            logo:'logo',
           
            name:'name',
            image1:'image1'
        })
    },
  
    
})
</script>
