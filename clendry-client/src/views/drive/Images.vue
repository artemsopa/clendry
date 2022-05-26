<template>
  <div class="container-fluid">
    <div class="row">
      <div class="col-lg-12">
        <div class="
            d-flex
            align-items-center
            justify-content-between
            welcome-content
            mb-3
          ">
          <div class="card card-block card-stretch card-transparent">
            <div class="card-header d-flex justify-content-between pb-0">
              <div class="header-title">
                <h4 v-if="files?.length > 0" class="card-title">Images</h4>
                <h4 v-if="files?.length < 1" class="card-title">No any image...</h4>
              </div>
            </div>
          </div>
          <div v-if="files?.length > 0" class="d-flex align-items-center">
            <div class="list-grid-toggle mr-4" @click="change()">
              <span class="icon icon-grid i-grid" v-if="data"><i class="ri-layout-grid-line font-size-20"></i></span>
              <span class="icon i-list" v-else><i class="ri-list-check font-size-20"></i></span>
              <span class="label label-list">List</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div v-if="data" class="icon icon-grid i-grid">
      <div class="row">
        <div class="col-lg-3 col-md-6 col-sm-6" v-for="(list, index) in files" :key="index">
          <div class="card card-block card-stretch card-height">
            <div class="card-body image-thumb">
              <div style="height:135px" class="mb-4 text-center p-3 rounded iq-thumb">
                <!-- <div class="iq-image-overlay"></div> -->
                <a :href="list.url" target="_blank"><img :src="list.url" class="img-cnt img-fluid" alt="image1"></a>
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
                      <th scope="col">Last Edit</th>
                      <th scope="col">File Size</th>
                      <th scope="col"></th>
                    </tr>
                  </thead>
                  <tbody>
                    <tr v-for="(list, index) in files" :key="index">
                      <td>
                        <div class="d-flex align-items-center">
                          <div class="mr-3">
                            <a :href="list.url" target="_blank"><img :src="list.url" class="img-fluid avatar-30"
                                alt="image1" /></a>
                          </div>
                          <div :title="list.title" v-if="!list.isEdit" class="file-title">
                            {{ list.title }}
                          </div>
                          <div v-if="list.isEdit" class="file-title">
                            <input v-model="input.title" type="text" class="width-inp-file add-inp"
                              :placeholder="getTitle(list.title)">
                          </div>
                        </div>
                      </td>
                      <!-- <td>{{ list.owner }}</td> -->
                      <td>{{ getDate(list.created_at) }}</td>
                      <td>{{ getSize(list.size) }}</td>
                      <td>

                        <div class="opsss card-header-toolbar">
                          <div class="dropdown">
                            <span class="dropdown-toggle" id="dropdownMenuButton5" data-toggle="dropdown">
                              <i class="ri-more-fill"></i>
                            </span>
                            <div class="dropdown-menu dropdown-menu-right">
                              <div v-if="list.isEdit" class="dropdown-item" @click="save(list)">
                                <i class="ri-pencil-fill mr-2"></i>save
                              </div>
                              <div v-if="list.isEdit" class="dropdown-item" @click="cancel(list)"><i
                                  class="ri-delete-bin-6-fill mr-2"></i>cancel</div>
                              <a v-if="!list.isEdit" :href="list.url" target="_blank">
                                <div class="dropdown-item"><i class="ri-eye-fill mr-2"></i>view</div>
                              </a>
                              <a v-if="!list.isEdit" :href="list.url" target="_blank" :download="list.title">
                                <div class="dropdown-item"><i class="ri-download-fill mr-2"></i>download</div>
                              </a>

                              <div v-if="!list.isEdit" class="dropdown-item" @click="edit(list)">
                                <i class="ri-pencil-fill mr-2"></i>edit
                              </div>
                              <hr style="margin-top:0px;margin-bottom:0px" v-if="!list.isEdit">
                              <div @click="addToFav(list.id)" v-if="!list.isEdit && !list.is_fav" class="dropdown-item">
                                <i class="lar la-star mr-2"></i>like
                              </div>
                              <div @click="removeFav(list.id)" v-if="!list.isEdit && list.is_fav" class="dropdown-item">
                                <i class="ri-star-fill mr-2"></i>unlike
                              </div>
                              <div @click="addToTrash(list.id)" v-if="!list.isEdit" class="dropdown-item"><i
                                  class="ri-delete-bin-6-fill mr-2"></i>trash
                              </div>
                            </div>
                          </div>
                        </div>

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
import File from '@/components/lists/File.vue';
import axios from "axios";
import FileServ from "@/models/file";
import { reactive } from "vue";

export default defineComponent({
  name: "Images",
  components: {
    File
  },
  setup() {
    const input = reactive({
      title: "",
      create: "",
    });
    return {
      input
    }
  },
  data() {
    return {
      data: true,
      files: [] as FileServ[],
      ext: "",
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
    getDate(str: string) {
      let date = str.split('-').reverse();
      date[0] = date[0][0] + date[0][1];
      return date[0] + "." + date[1] + "." + date[2];
    },
    removeFile(file: FileServ, index: number) {
      this.files.splice(index, 1);
      return this.files
    },
    downloadFile(file: FileServ) {
      let link = document.createElement('a');
      link.setAttribute('href', file.url);
      link.setAttribute('download', file.title);
      link.click();
    },
    edit(file: FileServ) {
      for (let i = 0; i < this.files.length; i++) {
        this.files[i].isEdit = false;
      }
      this.input.title = this.getTitle(file.title);
      file.isEdit = true;
    },
    async save(file: FileServ) {
      file.title = this.input.title + "." + this.ext;
      this.ext = "";
      file.isEdit = false;
      await axios.put(`/storage/files/title`, {
        id: file.id,
        title: file.title
      }, {
        withCredentials: true
      }).then(() => {
        this.getAllImages();
      })
    },
    cancel(file: FileServ) {
      file.isEdit = false;
    },
    getTitle(str: string) {
      let inp = str.split('.');
      this.ext = inp[inp.length - 1];
      inp.splice(inp.length - 1, 1);
      return inp.join('');
    },
    async getAllImages() {
      await axios.get(`/storage/files/image`, {
        withCredentials: true
      }).then(response => {
        this.files = response!.data;
      })
    },
    async addToFav(id: string) {
      await axios.put(`/storage/files/fav/`, {
        id: id,
      }, {
        withCredentials: true
      });
      await this.getAllImages();
    },
    async removeFav(id: string) {
      await axios.put(`/storage/files/fav/remove`, {
        id: id
      },{
        withCredentials: true
      });
      await this.getAllImages();
    },
    async addToTrash(id: string) {
      await axios.put(`/storage/files/trash/`, {
        id: id,
      }, {
        withCredentials: true
      })
      await this.getAllImages();
    }
  },
  async mounted() {
    await this.getAllImages();
  },
});
</script>

<style lang="scss">
.file-title {
  max-width: 350px;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.img-cnt {
  border-radius: 5px;
}

.width-inp-file {
  width: 400px;
}

#dropdownMenuButton:hover {
  cursor: pointer;
}

.opsss :hover {
  cursor: pointer;
}

.file-title :hover {
  cursor: pointer;
}

.dropdown-menu {
  .item:hover {
    cursor: pointer;
  }
}

a {
  color: #535f6b;
}
</style>