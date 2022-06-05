<template>
  <div class="iq-sidebar sidebar-default">
    <div class="iq-sidebar-logo d-flex align-items-center justify-content-between">
      <router-link :to="{ name: '' }" class="header-logo">
        <img :src="logo" class="img-fluid rounded-normal" alt="logo" />
      </router-link>
      <div class="iq-menu-bt-sidebar">
        <i class="las la-bars wrapper-menu"></i>
      </div>
    </div>
    <div class="data-scrollbar" data-scroll="1" id="sidebar-scrollbar">
      <div class="new-create select-dropdown input-prepend input-append">
        <div class="btn-group">
          <label class="iq-user-toggle" id="dropdownMenuButton" data-toggle="dropdown" aria-haspopup="true"
            aria-expanded="false">
            <div class="search-query selet-caption">
              <i class="las la-plus pr-2"></i>Create New
            </div>
            <span class="search-replace"></span>
            <span class="caret">
              <!--icon-->
            </span>
          </label>
          <ul class="dropdown-menu" aria-labelledby="dropdownMenuButton">
            <li>
              <div class="item add-it">
                <div class="add-blck">
                  <i @click="add()" class="ri-folder-add-line pr-3"></i>
                  <input v-model="input.create" maxlength="25" class="width-inp-create add-inp" type="text"
                    placeholder="New Folder">
                </div>
              </div>
            </li>
            <li>
              <div class="item">
                <router-link to="/upload-file">
                  <i class="ri-file-upload-line pr-3"></i>Upload Files
                </router-link>
              </div>
            </li>
          </ul>
        </div>
      </div>
      <nav class="iq-sidebar-menu">

        <ul id="sidebar" class="iq-menu collapse show" data-bs-parent="sidebar" style="">
          <li class="">
            <router-link to="/" class="router-link-active iq-waves-effect" data-toggle="collapse" data-target="#home"
              aria-expanded="false" aria-controls="home"><i class="las la-home iq-arrow-left"></i><span>Dashboard</span>
              <small class=""></small>
            </router-link>
          </li>
          <li class=""><a aria-current="page" href="upload-file"
              class="router-link-active router-link-exact-active iq-waves-effect" data-toggle="collapse"
              data-target="#folders" aria-expanded="false" aria-controls="folders"><i
                class="las la-boxes"></i><span>Folders</span><i
                class="las la-angle-down iq-arrow-right arrow-hover"></i><i
                class="las la-angle-right iq-arrow-right arrow-active"></i><small class=""></small></a>
            <ul id="folders" class="collapse iq-menu iq-submenu collapse" data-bs-parent="sidebar">
              <div v-if="folders.length == 0">
                <li class="men-itm">
                  <div>
                    <router-link to="/" class="iq-waves-effect" data-toggle="collapse" data-target="#folder-1"
                      aria-expanded="false" aria-controls="folder-1"><i class="las la-object-group"></i><span>No any
                        folder...</span>
                      <small class=""></small>
                    </router-link>
                  </div>
                </li>
              </div>
              <div v-if="folders.length > 0">
                <div v-for="(folder, index) in folders" :key="index">
                  <li class="men-itm">
                    <div v-if="!folder.isEdit">
                      <router-link :to="{ name: 'layout.folder', params: { id: folder.id } }" class="iq-waves-effect" data-toggle="collapse" data-target="#folder-1"
                        aria-expanded="false" aria-controls="folder-1"><i class="las la-object-group"></i><span>{{
                            folder.title
                        }}</span>
                        <small class=""></small>
                      </router-link>
                    </div>
                    <div v-if="folder.isEdit">
                      <router-link :to="{ name: 'layout.folder', params: { id: folder.id } }" class="iq-waves-effect" data-toggle="collapse" data-target="#folder-1"
                        aria-expanded="false" aria-controls="folder-1"><i class="las la-object-group"></i><span><input
                            v-model="input.title" type="text" maxlength="25" class="width-inp-add add-inp"
                            :placeholder="folder.title"></span>
                        <small class=""></small>
                      </router-link>
                    </div>
                    <div class="opsss card-header-toolbar">
                      <div class="dropdown">
                        <span class="dropdown-toggle" id="dropdownMenuButton5" data-toggle="dropdown">
                          <i class="ri-more-2-fill"></i>
                        </span>
                        <div class="dropdown-menu dropdown-menu-right" aria-labelledby="dropdownMenuButton5">
                          <ul v-if="!folder.isEdit">
                            <li class="dropdown-item" @click="edit(folder)"><i class="ri-pencil-fill mr-2"></i>Edit</li>
                            <li class="dropdown-item" @click="deleteFolder(folder)"><i
                                class="ri-delete-bin-6-fill mr-2"></i>Delete</li>
                          </ul>
                          <ul v-if="folder.isEdit">
                            <li class="dropdown-item" @click="save(folder)"><i class="ri-pencil-fill mr-2"></i>Save</li>
                            <li class="dropdown-item" @click="cancel(folder)"><i
                                class="ri-delete-bin-6-fill mr-2"></i>Cancel</li>
                          </ul>
                        </div>
                      </div>
                    </div>
                  </li>
                </div>
              </div>
            </ul>
          </li>
          <li class="">
            <router-link to="/images" class="iq-waves-effect" data-toggle="collapse" data-target="#images"
              aria-expanded="false" aria-controls="images"><i class="las la-image"></i><span>Images</span>
              <small class=""></small>
            </router-link>
          </li>
          <li class="">
            <router-link to="/media" class="iq-waves-effect" data-toggle="collapse" data-target="#media"
              aria-expanded="false" aria-controls="media"><i class="las la-photo-video"></i><span>Media</span>
              <small class=""></small>
            </router-link>
          </li>
          <li class="">
            <router-link to="/files" class="iq-waves-effect" data-toggle="collapse" data-target="#files"
              aria-expanded="false" aria-controls="files"><i
                class="lar la-file-alt iq-arrow-left"></i><span>Files</span>
              <small class=""></small>
            </router-link>
          </li>
          <li class="">
            <router-link to="/favourites" class="iq-waves-effect" data-toggle="collapse" data-target="#favourites"
              aria-expanded="false" aria-controls="favourites"><i class="lar la-star"></i><span>Favourites</span>
              <small class=""></small>
            </router-link>
          </li>
          <li class="">
            <router-link to="/trash" class="iq-waves-effect" data-toggle="collapse" data-target="#trash"
              aria-expanded="false" aria-controls="trash"><i
                class="las la-trash-alt iq-arrow-left"></i><span>Trash</span>
              <small class=""></small>
            </router-link>
          </li>
        </ul>

      </nav>
      <div class="sidebar-bottom">
        <h4 class="mb-3"><i class="las la-cloud mr-2"></i>Storage</h4>
        <p>{{usedGB[0]}} / 20 GB Used</p>
        <Progressbar :value="usedGB[1]" color="primary" class="mb-3" midclass="iq-progress progress-1" />
        <p>{{ usedGB[1] }}% Full - {{20 - usedGB[0]}} GB Free</p>
        <router-link to="/plan" class="btn btn-outline-primary view-more mt-4">Try Another Plan</router-link>
      </div>
      <div class="p-3"></div>
    </div>
  </div>
</template>

<script lang="ts">
import axios from "axios";
import { defineComponent, reactive } from "vue";
import Progressbar from "../progressbar/Progressbar.vue";
import { mapGetters } from "vuex";
import { core } from "../../config/pluginInit";
import Folder from "../../models/folder";

export default defineComponent({
  name: "Sidebar",
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
      homeurl: "",
      folders: [] as Folder[],
      usedGB: [] as number[]
    };
  },
  async mounted() {
    core.SmoothScrollbar();
    core.changesidebar();
    await this.getAllFolders();
    let kb = await this.getKB();
    this.usedGB[0] = this.getSize(kb);
    this.usedGB[1] = Math.round((this.usedGB[0] / 20) * 100);
    this.usedGB = this.usedGB.slice();
  },
  destroyed() {
    core.SmoothScrollbar();
    core.changesidebar();
  },
  components: {
    Progressbar,
  },
  methods: {
    async getKB() {
      const response = await axios.get(`/storage/kb`, {
        withCredentials: true
      });
      return response.data;
    },
    getSize(size: number): number {
      return Number((size / (1024 * 1000000)).toFixed(2));
    },
    edit(folder: Folder) {
      for (let i = 0; i < this.folders.length; i++) {
        this.folders[i].isEdit = false;
      }
      this.input.title = folder.title;
      folder.isEdit = true;
    },
    async save(folder: Folder) {
      folder.title = this.input.title;
      await axios.put(`/storage/folders/`, {
        id: folder.id,
        title: folder.title
      }, {
        withCredentials: true
      }).then(() => {
        this.getAllFolders();
      });
      location.reload();
    },
    async add() {
      if (this.input.create.length > 5) {
        let folder = new Folder();
        folder.title = this.input.create;
        folder.isEdit = false;
        await axios.post(`/storage/folders/`, {
          title: folder.title
        }, {
          withCredentials: true
        }).then(() => {
        this.getAllFolders();
        })
      }
      this.input.create = "";
      location.reload();
    },
    cancel(folder: Folder) {
      folder.isEdit = false;
    },
    async deleteFolder(folder: Folder) {
      await axios.delete(`/storage/folders/${folder.id}`, {
          withCredentials: true
        }).then(() => {
        this.getAllFolders();
        })
      location.reload();
    },
    async getAllFolders() {
      await axios.get<Folder[]>(`/storage/folders/`, {
        withCredentials: true
      }).then(response => {
        this.folders = response.data;
      })
    }
  },
  computed: {
    ...mapGetters({
      appName: "appName",
      logo: "logo",
    }),
  },
});
</script>

<style lang="scss">
.width-inp-add {
  width: 130px;
}

.width-inp-create {
  width: 150px;
}

.add-blck {
  display: flex;
}

.add-it :hover {
  color: #8f93f6;
}

.men-itm {
  display: flex;
  justify-content: space-between;
}

.add-blck i {
  padding-top: 7px;
  margin-right: 10px;
  width: 20px;
}

.add-inp {
  border: 0px;
  background-color: rgba(240, 248, 255, 0);
  font-size: 16px;
}

.add-inp:hover :active :focus {
  background-color: rgba(240, 248, 255, 0);
}

.opsss {
  padding-top: 10px;
  padding-right: 20px;
}

.opsss :hover {
  cursor: pointer;
}

#dropdownMenuButton:hover {
  cursor: pointer;
}

.dropdown-menu {
  .item:hover {
    cursor: pointer;
  }
}
</style>
