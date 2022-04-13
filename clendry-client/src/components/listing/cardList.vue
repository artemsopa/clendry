<template>
     <card class=" card-block card-stretch card-height"  bodyClass="image-thumb">
         <template v-slot:body>
            <div class="mb-4 text-center p-3 rounded iq-thumb">
                <a class="image-popup-vertical-fit" :href="data.src">
                    <img :src="data.src" class="img-fluid" alt="images" />  
                    <div class="iq-image-overlay"></div>
                </a>           
            </div>
            <div class="d-flex justify-content-between"> 
            <h6>{{data.name}}</h6> 
            <slot name="dropdown"/>   
            </div>      
         </template>
    </card>
</template>

<script lang="ts">
 import {core} from '@/config/pluginInit'
import { defineComponent } from '@vue/runtime-core'

export default defineComponent({
    name:'cardList',
   mounted(){
       core.verticalImagePopup()
    },
    props:{
        data:{
            type:Object,
            require:true
        }
    }
})
</script>