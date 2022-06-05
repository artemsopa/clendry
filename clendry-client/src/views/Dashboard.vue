<template>
    <div class="container-fluid">
        <div class="row">
            <div class="col-lg-12">
                <div class="card-transparent card-block card-stretch card-height mb-3">
                    <div class="d-flex justify-content-between">
                        <div class="header-title">
                            <h4 class="card-title">My Drive</h4>
                        </div>
                    </div>
                </div>
            </div>
            <div class="col-lg-8">
                <div class="card card-block card-stretch card-height iq-welcome"
                    :style="`background: url(${require('@/assets/images/layouts/mydrive/background.png')}) no-repeat scroll right center; background-color: #ffffff; background-size: contain;`">
                    <div class="card-body property2-content">
                        <div class="d-flex flex-wrap align-items-center">
                            <div class="col-lg-6 col-sm-6 p-0">
                                <h3 class="mb-3">Welcome @{{user.nick}}!</h3>
                                <p class="mb-5">We are glad to see you in our service!</p>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
            <div class="col-lg-4">
                <div class="card card-block card-stretch card-height  plan-bg">
                    <div class="card-body">
                        <h4 class="mb-3 text-white">Choose Your Plan</h4>
                        <p>Expanded Storage, Access To<br> More Features On Clendry</p>
                        <div class="row align-items-center justify-content-between">
                            <div class="col-6 go-white ">
                                <router-link to="/plan" class="btn d-inline-block mt-5">Try Free</router-link>
                            </div>
                            <div class="col-6">
                                <img src="../assets/images/layouts/mydrive/lock-bg.png" class="img-fluid" alt="image1">
                            </div>
                        </div>
                    </div>
                </div>
            </div>


            <div class="col-lg-12"></div>

            <div class="col-lg-4">
                <div class="card card-block card-stretch card-height ">
                    <div class="card-header d-flex justify-content-between">
                        <div class="header-title">
                            <h4 class="card-title">Statistic</h4>
                        </div>
                    </div>
                    <div class="card-body">
                        <AmChart element="layout-1-chart" type="dashboard1chart" />
                        <div class="row mt-4">
                            <div class="col-lg-6 col-md-6 col-6">
                                <div class="media align-items-center">
                                    <div class="icon iq-icon-box bg-primary rounded icon-statistic">
                                        <i class="las la-long-arrow-alt-down"></i>
                                    </div>
                                    <div class="media-body ml-3">
                                        <p class="mb-0">Downloads</p>
                                        <h5>2</h5>
                                    </div>
                                </div>
                            </div>
                            <div class="col-lg-6 col-md-6 col-6">
                                <div class="media align-items-center">
                                    <div class="icon iq-icon-box bg-light rounded icon-statistic">
                                        <i class="las la-long-arrow-alt-up"></i>
                                    </div>
                                    <div class="media-body ml-3">
                                        <p class="mb-0">Uploads</p>
                                        <h5>23</h5>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
            <div class="col-lg-8">
                <div class="card card-block card-stretch card-height">
                    <div class="card-header d-flex justify-content-between pb-0">
                        <div class="header-title">
                            <h4 class="card-title">Storage</h4>
                        </div>
                    </div>
                    <div class="card-header d-flex justify-content-between pb-0">
                        <div class="header-title">
                            <p class="card-title">Get your uploading statistics by months!</p>
                        </div>
                    </div>
                    <div style="margin-top:50px" class="card-body pt-0">
                        <ApexChart element="layout-1-chart2" :chartOption="chart1" />
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
import { defineComponent } from "@vue/runtime-core"
import axios from "axios"
import AmChart from '../components/charts/AmChart.vue'
import ApexChart from '../components/charts/ApexChart.vue'
import Folder from "../models/folder"
import User from "../models/user"

export default defineComponent({
    name: 'Dashbord',
    components: {
        ApexChart,
        AmChart
    },
    data() {
        return {
            user: {} as User,
            chart1: {
                series: [{
                    name: "Uploadings",
                    data: [0, 0, 0, 0, 0, 0, 0, 23, 23]
                }],
                colors: ["#8f93f6"],
                chart: {
                    toolbar: {
                        show: false
                    },
                    height: 225,
                    type: 'line',
                    zoom: {
                        enabled: false
                    }
                },
                dataLabels: {
                    enabled: false
                },
                stroke: {
                    curve: 'straight'
                },
                title: {
                    text: '',
                    align: 'left'
                },
                grid: {
                    row: {
                        colors: ['#f3f3f3', 'transparent'],
                        opacity: 0.5
                    },
                },
                xaxis: {
                    categories: ['Nov', 'Dec' ,'Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', '*Curr'],
                    labels: {
                        minHeight: 20,
                        maxHeight: 20
                    }
                },
                yaxis: {
                    show: true,
                    labels: {
                        minWidth: 15,
                        maxWidth: 15
                    }
                }
            },

        }
    },
    methods: {
        async getUser() {
            await axios.get<User>(`/profile`, {
                withCredentials: true
            }).then(response => {
                this.user = response.data;
            });
        },
    },
    async mounted() {
        await this.getUser()
    }
})
</script>
