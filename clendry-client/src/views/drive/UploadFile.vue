<template>
  <div class="container-fluid" style="cursor: default">
    <div class="row">
      <div class="col-lg-12">
        <div class="card-transparent card-block card-stretch card-height mb-3">
          <div class="d-flex justify-content-between">
            <div class="select-dropdown input-prepend input-append">
              <div class="btn-group">
                <label data-toggle="dropdown">
                  <div class="dropdown-toggle search-query">Upload Files</div>
                  <span class="caret">
                    (100 MB MAX)
                  </span>
                </label>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="col-lg-8">
        <file-selector v-model="files">
          <dropzone>
            <div
              style="
                cursor: move;
                height: 140px;
                background: linear-gradient(90deg, rgba(125,136,255,1) 12%, rgba(85,248,255,1) 57%, rgba(234,112,255,1) 93%);"
              class="card card-block card-stretch card-height">
              <div class="card-body property2-content">
                <div class="d-flex flex-wrap align-items-center">
                  <div class="col-lg-12 col-sm-6 p-0">
                    <h3 class="mb-3">Drag & Drop</h3>
                    <h6>Move any file for uploading</h6>
                  </div>
                </div>
              </div>
            </div>
          </dropzone>
        </file-selector>
      </div>
      <div class="col-lg-4">
        <div class="card card-block card-stretch card-height"
          style="
            cursor: pointer;
            height: 140px;
            background: linear-gradient(90deg,rgba(198, 141, 236, 1) 0%,rgba(255, 115, 115, 1) 50%,rgba(255, 207, 140, 1) 100%);">
          <div class="card-header d-flex justify-content-between">
            <div class="col-lg-12 col-sm-6 p-0">
              <h3 class="card-title">Select Files</h3>
              <h6>Click to select files</h6>
              <input multiple type="file" @change="onChangeFiles" style="
              opacity: 0;
              position: absolute;
              top: 0;
              left: 0;
              bottom: 0;
              right: 0;
              width: 100%;
              height: 100%;
              cursor: pointer;
            " />
            </div>
          </div>
          <div class="card-body">
            <ul class="list-inline p-0 mb-0 row align-items-center">
              <li class="col-lg-12 col-sm-6 mb-3 mb-sm-0">
                <div data-load-file="file" data-load-target="#resolte-contaniner" class="p-2 text-center rounded"></div>
              </li>
            </ul>
          </div>
        </div>
      </div>


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
                    <h4 v-if="files?.length > 0" class="card-title">Selected Files</h4>
                    <h4 v-if="files?.length < 1" class="card-title">No any selected file...</h4>
                  </div>
                </div>
              </div>
              <div v-if="files?.length > 0" class="d-flex align-items-center">
                <div class="list-grid-toggle mr-4">
                  <div class="header-title">
                    <button @click="uploadFiles()" type="button" class="btn btn-primary" data-toggle="button"
                      aria-pressed="false">
                      Upload Files
                    </button>
                  </div>
                </div>
                <div class="list-grid-toggle mr-4">
                  <div class="header-title">
                    <div class=""><a aria-current="page" href="upload-file"
                        class="router-link-active router-link-exact-active iq-waves-effect" data-toggle="collapse"
                        data-target="#choose" aria-expanded="false" aria-controls="choose">
                        <span>
                          <div type="button" class="btn btn-success" data-toggle="button" aria-pressed="false">
                            Selected Folders - {{ selected.length }}
                          </div>
                        </span><small class=""></small></a>
                    </div>
                  </div>

                </div>
                <div class="list-grid-toggle mr-4">
                  <div class="header-title">
                    <button @click="clearFiles()" type="button" class="btn btn-warning" data-toggle="button"
                      aria-pressed="false">
                      Clear All
                    </button>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <div v-if="files?.length > 0" id="choose" class="collapse iq-menu iq-submenu" data-bs-parent="sidebar">

          <div class="jstfc">
            <div class="fold" @click="chooseFolder(folder)" v-for="(folder, index) in folders" :key="index">
              <div :class="'folder-item ' + getStyle(folder)">
                <div class="icn">
                  <i class="ri-file-copy-line"></i>
                </div>
                <div>
                  {{ folder.title }}
                </div>
              </div>
            </div>
          </div>
        </div>

        <div v-if="files?.length > 0" class="icon i-list">
          <div class="row">
            <div class="col-lg-12">
              <div class="card card-block card-stretch card-height">
                <div class="card-body">
                  <div class="table-responsive">
                    <table class="table mb-0 table-borderless tbl-server-info">
                      <thead>
                        <tr>
                          <th scope="col">Name</th>
                          <th style="text-align: center;" class="align-center" scope="col">Type</th>
                          <th style="text-align: center;" class="text-align-center" scope="col">Size</th>
                          <th style="text-align: center;" class="text-align-center" scope="col">Progress</th>
                          <th scope="col"></th>
                        </tr>
                      </thead>
                      <tbody>
                        <tr v-for="(list, index) in files" :key="index">
                          <td>
                            <div class="d-flex align-items-center">
                              <div :class="'icon-small rounded mr-3 ' + getColor(list)">
                                <i :class="getIcon(list)"></i>
                              </div>
                              <div :title="list.name" class="file-title">
                                {{ list.name }}
                              </div>
                            </div>
                          </td>
                          <td style="text-align: center;" :title="list.type ? list.type : `unknown type...`"
                            class="file-type" :style="list.type ? `background-colot:red` : ``">
                            {{ getExtension(list) }}
                          </td>
                          <td style="text-align: center;">{{ getSize(parseInt(list.size)) }}</td>
                          <td style="width: 180px, text-align: center;">
                            <div style="margin-top:17px" class="progress mb-3">
                              <div class="progress-bar progress-bar-striped progress-bar-animated" role="progressbar"
                                aria-valuemin="0" aria-valuemax="100" :style="{ width: list.state + '%' }">
                              </div>
                            </div>
                          </td>
                          <td class="width-btn-dang" v-if="list.state == 0" @click="removeFile(index)">
                            <button type="button" class="btn btn-danger"><i
                                class="ri-delete-bin-2-fill pr-0"></i>Remove</button>
                          </td>
                          <td class="width-btn-dang" v-if="list.state > 0" :title="list.status">
                            <button type="button" class="btn">{{ list.status }}</button>
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


    </div>
  </div>
</template>
<script lang="ts">
import Folder from "@/models/folder";
import { defineComponent } from "@vue/runtime-core";
import axios from "axios";
import { ref } from "vue";
import { FileSelector, Dropzone, DialogButton } from "vue3-file-selector";
import FileItem from "../../models/fileItem"

export default defineComponent({
  name: "UploadFile",
  components: {
    FileSelector,
    Dropzone,
    DialogButton,
  },
  setup() {
    const files = ref<FileItem[]>([]);
    let selected: any[] = [];
    return {
      files,
      selected,
    };
  },
  data() {
    return {
      folders: [] as Folder[],
      ids: [] as any[],
    }
  },
  methods: {
    chooseFolder(folder: any) {
      folder.iscurr = !folder.iscurr;
      this.folders = this.folders.slice();
      this.getAllSelected();
    },
    getAllSelected() {
      this.selected = [];
      for (let i = 0; i < this.folders.length; i++) {
        if (this.folders[i].iscurr) {
          this.selected.push(this.folders[i]);
        }
      }
    },
    getStyle(folder: any) {
      if (folder.iscurr) {
        return "bg-success rounded mb-4";
      }
      return folder.color;
    },
    getRandomColor() {
      let arr = ["bg-danger rounded mb-4", "bg-primary rounded mb-4 ", "bg-info rounded mb-4"];
      let count = 0;
      for (let i = 0; i < this.folders.length; i++) {
        this.folders[i].color = arr[count];
        count++;
        if (count == 3) {
          count = 0;
        }
      }
    },
    removeFile(index: number) {
      this.files.splice(index, 1);
      return this.files
    },
    getSize(size: number): string {
      let spec = size / 1024;
      return size > 1000 * 1024 ?
        size > 1000000 * 1024 ? (spec / 1000000).toFixed(2) + " GB" : (spec / 1000).toFixed(2) + " MB" :
        Math.ceil(spec) + " KB";
    },
    getExtension(file: any) {
      let ext: string = file.name.split(".").pop();
      return ext.toUpperCase();
    },
    getIcon(file: any) {
      let type = file.type.split("/")[0];
      if (type == "image") {
        return "far fa-file-image";
      }
      if (type == "video") {
        return "far fa-file-video";
      }
      return "far fa-file-alt";
    },
    getColor(file: any) {
      let type = file.type.split("/")[0];
      if (type == "image") {
        return "bg-info";
      }
      if (type == "video") {
        return "bg-primary";
      }
      return "bg-success";
    },
    onChangeFiles(event: any) {
      let filesEv = event.target.files;
      for (let i = 0; i < filesEv.length; i++) {
        filesEv[i].state = 0;
        filesEv[i].error = '';
        this.files.unshift(filesEv[i]);
      }
    },
    clearFiles() {
      this.files = [];
      for (let i = 0; i < this.folders.length; i++) {
        this.folders[i].iscurr = false;
      }
      this.selected = [];
    },
    async uploadFiles() {
      console.log(this.selected);
      for (let i = 0; i < this.files.length; i++) {
        if (this.files[i].state > 0) continue;
        else {
          try {
            await this.uploadFile(this.files[i]);
          } catch (error) {
            console.log(error);
          }
        }
      }
      if (this.selected.length > 0 && this.ids.length > 0) {
        for (let i = 0; i < this.ids.length; i++) {
          for (let j = 0; j < this.selected.length; j++) {
            try {
              const response = await axios.put("/storage/files/folder", {
                folder_id: this.selected[j].id,
                file_id: this.ids[i],
              }, {
                withCredentials: true,
              })
            } catch (error) {
              console.log(error);
            }
          }
        }
        for (let i = 0; i < this.folders.length; i++) {
          this.folders[i].iscurr = false;
        }
        this.selected = [];
      }
      this.ids = [];
      return this.files;
    },
    async uploadFile(file: any) {
      let form = new FormData();
      form.append("file", file);
      const response = await axios.post("/storage/files/upload", form, {
        withCredentials: true,
        onUploadProgress: e => {
          file.status = "Loading..."
          file.state = Math.round((e.loaded / e.total) * 100);
          this.files = this.files.slice();
        }
      })
      console.log(response.status);
      if (!response.status) {
        file.status = "Error!"
      } else {
        file.status = "Finished"
        this.ids.push(response!.data);
      }
      this.files = this.files.slice();
    },

    async getFolders() {
      await axios.get<Folder[]>(`/storage/folders/`, {
        withCredentials: true
      }).then(response => {
        this.folders = response.data;
        this.getRandomColor();
      })
    },
  },
  mounted() {
    this.getFolders();
  }
});
</script>

<style lang="scss">
.file-title {
  max-width: 350px;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.file-type {
  max-width: 100px;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.icn {
  margin-right: 5px;
  padding-top: 2px;
}

.folder-item {
  padding: 5px 20px;
  margin: 0px 10px;
  display: flex;
}

.fold :hover {
  cursor: pointer;
}

.width-btn-dang {
  width: 100px;
}

.jstfc {
  display: flex;
  flex-wrap: wrap;
}
</style>