<template>
    <div class="container-fluid">
        <div class="row">
            <div class="pricing-custom-tab w-100">
                <div class="pricing-content">
                    <div>
                        <div class="d-flex align-items-center">
                            <div class="col-lg-4 col-sm-6">
                                <div class="card card-block card-stretch card-height blog pricing-details">
                                    <div class="card-body border text-center rounded">
                                        <div class="pricing-header bg-primary text-white">
                                            <div class="icon-data"><i class="ri-star-line"></i>
                                            </div>
                                            <h2 class="mb-4 display-5 font-weight-bolder text-white">Basic</h2>
                                        </div>
                                        <h4 class="mb-3">Basic Usage</h4>
                                        <ul class="list-unstyled mb-0 pricing-list">
                                            <li><i class="lar la-check-circle text-primary mr-2 font-size-20"></i>
                                            10 GB Storage</li>
                                        </ul>
                                        <div v-if="user.memory != 10" @click="updatePlan(10)" class="btn btn-primary mt-5">Try Free</div>
                                        <div v-else class="btn btn-secondary mt-5">Used</div>
                                    </div>
                                </div>
                            </div>
                            <div class="col-lg-4 col-sm-6">
                                <div class="card card-block card-stretch card-height blog pricing-details">
                                    <div class="card-body border text-center rounded">
                                        <div class="pricing-header bg-success text-white">
                                            <div class="icon-data"><i class="ri-star-line"></i>
                                            </div>
                                            <h2 class="mb-4 display-5 font-weight-bolder text-white">Prem</h2>
                                        </div>
                                        <h4 class="mb-3">Premium Usage</h4>
                                        <ul class="list-unstyled mb-0 pricing-list">
                                            <li><i class="lar la-check-circle text-primary mr-2 font-size-20"></i>
                                            100 GB Storage</li>
                                        </ul>
                                        <div v-if="user.memory != 100" @click="updatePlan(100)" class="btn btn-success mt-5">Try Free</div>
                                        <div v-else class="btn btn-secondary mt-5">Used</div>
                                    </div>
                                </div>
                            </div>
                            <div class="col-lg-4 col-sm-6">
                                <div class="card card-block card-stretch card-height blog pricing-details">
                                    <div class="card-body border text-center rounded">
                                        <div class="pricing-header bg-danger text-white">
                                            <div class="icon-data"><i class="ri-star-line"></i>
                                            </div>
                                            <h2 class="mb-4 display-5 font-weight-bolder text-white">Prem +</h2>
                                        </div>
                                        <h4 class="mb-3">Premium + Usage</h4>
                                        <ul class="list-unstyled mb-0 pricing-list">
                                            <li><i class="lar la-check-circle text-primary mr-2 font-size-20"></i>
                                            Unlimited Storage</li>
                                        </ul>
                                        <div v-if="user.memory != 0" @click="updatePlan(0)" class="btn btn-danger mt-5">Try Free</div>
                                        <div v-else class="btn btn-secondary mt-5">Used</div>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import axios from 'axios';
import User from '../../models/user';

export default defineComponent({
    setup() {
        
    },
    data() {
        return {
            user: {} as User
        }
    },
    async mounted() {
        await this.getUser();
    },
    methods: {
        async updatePlan(gb: number) {
            await axios.put(`/profile/memory?memory=${gb}`, { }, {
        withCredentials: true
      });
      location.reload();
        },
        async getUser() {
      await axios.get<User>(`/profile`, {
        withCredentials: true
      }).then(response => {
        this.user = response.data;
      });
    },
    }
})
</script>
