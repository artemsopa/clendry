<template>
  <div class="container-fluid">
    <div class="row">
      <div class="col-lg-12">
        <div
          class="
            d-flex
            align-items-center
            justify-content-between
            welcome-content
            mb-3
          "
        >
          <div class="card card-block card-stretch card-transparent">
            <div class="card-header d-flex justify-content-between pb-0">
              <div class="header-title">
                <h4 v-if="lists?.length>0" class="card-title">{{ title }}</h4>
                <h4 v-if="lists?.length<1" class="card-title">{{ errTitle }}</h4>
              </div>
            </div>
          </div>
          <div v-if="lists?.length>0" class="d-flex align-items-center">
            <div class="list-grid-toggle mr-4" @click="change()">
              <span class="icon icon-grid i-grid" v-if="data"
                ><i class="ri-layout-grid-line font-size-20"></i
              ></span>
              <span class="icon i-list" v-else
                ><i class="ri-list-check font-size-20"></i
              ></span>
              <span class="label label-list">List</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div v-if="data" class="icon icon-grid i-grid">
      <div class="row">
        <div
          class="col-lg-3 col-md-6 col-sm-6"
          v-for="(list, index) in lists"
          :key="index"
        >
          <div class="card card-block card-stretch card-height">
            <div class="card-body image-thumb">
              <div class="mb-4 text-center p-3 rounded iq-thumb">
                <div class=""></div>
                <a
                  href="#"
                  :data-title="list.title"
                  :data-url="list.url"
                  @click="
                    $root.$emit(
                      'bv::show::modal',
                      'viewer-modal',
                      $event.target
                    )
                  "
                  v-b-modal.viewer-modal
                  ><img :src="list.url" class="img-fluid" alt="image1"
                /></a>
              </div>
              <h6 :title="list.title" class="file-title">{{ list.title }}</h6>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div v-else class="icon i-list">
      <div class="row">
        <div class="col-lg-12">
          <div class="card card-block card-stretch card-height">
            <div class="card-body">
              <div class="table-responsive">
                <table class="table mb-0 table-borderless tbl-server-info">
                  <thead>
                    <tr>
                      <th scope="col">Name</th>
                      <th scope="col">Owner</th>
                      <th scope="col">Last Edit</th>
                      <th scope="col">File Size</th>
                      <th scope="col"></th>
                    </tr>
                  </thead>
                  <tbody>
                    <tr v-for="(list, index) in lists" :key="index">
                      <td>
                        <div class="d-flex align-items-center">
                          <div class="mr-3">
                            <a href="#"
                              ><img
                                :src="list.url"
                                class="img-fluid avatar-30"
                                alt="image1"
                            /></a>
                          </div>
                          <template class="file-title" v-if="list.img">
                            {{ list.title }}
                          </template>
                          <template v-else>
                            <a class="file-title"
                              href="#"
                              :data-title="list.title"
                              :data-url="list.url"
                              @click="
                                $root.$emit(
                                  'bv::show::modal',
                                  'viewer-modal',
                                  $event.target
                                )
                              "
                              v-b-modal.viewer-modal
                            >
                              <span style="color: #535f6b">{{
                                list.title
                              }}</span>
                            </a>
                          </template>
                        </div>
                      </td>
                      <td>{{ list.owner }}</td>
                      <td>{{ list.lastedit }}</td>
                      <td>{{ getSize(list.size) }}</td>
                      <td>
                        <b-dropdown
                          id="dropdownMenuButton"
                          right
                          variant="none"
                          data-toggle="dropdown"
                        >
                          <template v-slot:button-content>
                            <i class="ri-more-fill"></i>
                          </template>
                          <b-dropdown-item
                            ><i class="ri-eye-fill mr-2"></i
                            >{{ "view" }}</b-dropdown-item
                          >
                          <b-dropdown-item
                            ><i class="ri-delete-bin-6-fill mr-2"></i
                            >{{ "delete" }}</b-dropdown-item
                          >
                          <b-dropdown-item
                            ><i class="ri-pencil-fill mr-2"></i
                            >{{ "edit" }}</b-dropdown-item
                          >
                          <b-dropdown-item
                            ><i class="ri-printer-fill mr-2"></i
                            >{{ "print" }}</b-dropdown-item
                          >
                          <b-dropdown-item
                            ><i class="ri-file-download-fill mr-2"></i
                            >{{ "download" }}</b-dropdown-item
                          >
                        </b-dropdown>
                      </td>
                    </tr>
                  </tbody>
                </table>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "@vue/runtime-core";

export default defineComponent({
  name: "File",
  props: {
    title: String,
    errTitle: String,
    lists: Array,
  },
  data() {
    return {
      data: true,
    };
  },
  methods: {
    change() {
      this.data = !this.data;
    },
    getSize(size: number): string {
      let spec = size / 1024;
      return size > 1000 * 1024 ?
        size > 1000000 * 1024 ? (spec / 1000000).toFixed(2) + " GB" : (spec / 1000).toFixed(2) + " MB" :
        Math.ceil(spec) + " KB";
    },
  },
});
</script>

<style>
.file-title {
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}
</style>