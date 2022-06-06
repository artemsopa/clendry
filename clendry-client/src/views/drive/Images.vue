<template>
  <div>
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
            <div v-if="files?.length > 0" class="card-header-toolbar d-flex align-items-center">
              <div class="card-header-toolbar">
                <div class="dropdown">
                  <span class="dropdown-toggle dropdown-bg btn bg-white" id="dropdownMenuButton1"
                    data-toggle="dropdown">
                    Name<i class="ri-arrow-down-s-line ml-1"></i>
                  </span>
                  <div class="dropdown-menu dropdown-menu-right shadow-none" aria-labelledby="dropdownMenuButton1">
                    <div @click="sortDefault()" class="dropdown-item">Default</div>
                    <div @click="sortAZ()" class="dropdown-item">Title A-Z</div>
                    <div @click="sortZA()" class="dropdown-item">Title Z-A</div>
                    <div @click="sortSizeLow()" class="dropdown-item">Size smaller</div>
                    <div @click="sortSizeHigh()" class="dropdown-item">Size bigger</div>
                  </div>
                </div>
              </div>
                  <h1>⠀</h1>
              <div class="card-header-toolbar">
                <div>
                  <div class="list-grid-toggle mr-4" @click="change()">
                    <span class="icon icon-grid i-grid" v-if="data"><i
                        class="ri-layout-grid-line font-size-20"></i></span>
                    <span class="icon i-list" v-else><i class="ri-list-check font-size-20"></i></span>
                    <span class="label label-list">List</span>
                  </div>
                </div>
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

                          <!-- Modal -->
                          <div class="modal fade" id="exampleModal" tabindex="-1" role="dialog"
                            aria-labelledby="exampleModalLabel" aria-hidden="true">
                            <div class="modal-dialog" role="document">
                              <div class="modal-content">
                                <div class="modal-header">
                                  <div class="ojhjs d-flex align-items-center mb-3">
                                    <div class="profile-icon iq-icon-box rounded-small bg-light svg-danger text-center">
                                      <i class="ri-window-line text-primary"></i>
                                    </div>
                                    <div class="pl-3">
                                      <h5 :title="list.title" class="txt-hide modal-title" id="exampleModalLabel">
                                        {{ list.title }}</h5>
                                      <p class="txt-hide mb-0">{{ list.id }}</p>
                                    </div>
                                  </div>
                                </div>
                                <div class="modal-body">
                                  <h5>Click to remove from folder</h5>
                                </div>

                                <div v-if="folders[0].length || folders[0].length > 0">

                                  <div class="jstfc">
                                    <div class="fold" v-for="(folder, index) in folders[0]" :key="index">
                                      <div @click="deleteFolder(index, folder, list.id)"
                                        :class="'folder-item ' + getStyle(folder)">
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
                                <div class="text-center" v-else>
                                  <h6 style="margin-bottom:10px">No any selected folder...</h6>
                                </div>

                                <div class="modal-body">
                                  <h5>Click to add to folder</h5>
                                </div>
                                <div v-if="folders[1].length || folders[1].length > 0">

                                  <div class="jstfc">
                                    <div class="fold" v-for="(folder, index) in folders[1]" :key="index">
                                      <div @click="chooseFolder(index, folder, list.id)"
                                        :class="'folder-item ' + getStyle(folder)">
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
                                <div class="text-center" v-else>
                                  <h6 style="margin-bottom:10px">No any created folder...</h6>
                                </div>

                                <div class="modal-footer">
                                  <button type="button" class="btn btn-secondary" data-dismiss="modal">Close</button>
                                </div>
                              </div>
                            </div>
                          </div>

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
                                <div v-if="!list.isEdit" @click="downloadFile(list)">
                                  <div class="dropdown-item"><i class="ri-download-fill mr-2"></i>download</div>
                                </div>
                                <div v-if="!list.isEdit" class="dropdown-item" @click="edit(list)">
                                  <i class="ri-pencil-fill mr-2"></i>edit
                                </div>
                                <div v-if="!list.isEdit">
                                  <div class="dropdown-item">
                                    <!-- Button trigger modal -->
                                    <div @click="getSelectedFolders(list.id)" type="button" data-toggle="modal"
                                      data-target="#exampleModal">
                                      <i class="las la-share-alt mr-2"></i>struct
                                    </div>
                                  </div>
                                </div>
                                <hr style="margin-top:0px;margin-bottom:0px" v-if="!list.isEdit">
                                <div @click="addToFav(list.id)" v-if="!list.isEdit && !list.is_fav"
                                  class="dropdown-item">
                                  <i class="lar la-star mr-2"></i>like
                                </div>
                                <div @click="removeFav(list.id)" v-if="!list.isEdit && list.is_fav"
                                  class="dropdown-item">
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
  </div>
</template>

<script lang="ts">
import { defineComponent } from "@vue/runtime-core";
import File from '@/components/lists/File.vue';
import axios from "axios";
import FileServ from "@/models/file";
import { reactive } from "vue";
import Folder from "../../models/folder";

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
      input,
    }
  },
  data() {
    return {
      data: true,
      files: [] as FileServ[],
      ext: "",
      folders: [] as Folder[][],
      ids: [] as any[],
    };
  },
  methods: {
    async sortDefault() {
      this.getAllImages();
    },
    sortAZ() {
      this.files = this.files.sort((a, b) => a.title.localeCompare(b.title));
      },
    sortZA() {
      this.files = this.files.sort((a, b) => b.title.localeCompare(a.title));
      },
    sortSizeLow() {
      this.files = this.files.sort((a, b) => {return a.size - b.size;});
    },
    sortSizeHigh() {
      this.files = this.files.sort((a, b) => {return b.size - a.size;});
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
      if (this.folders[1]) {
        for (let i = 0; i < this.folders[1].length; i++) {
          this.folders[1][i].color = arr[count];
          count++;
          if (count == 3) {
            count = 0;
          }
        }
      }

      if (this.folders[0]) {
        for (let i = 0; i < this.folders[0].length; i++) {
          this.folders[0][i].color = "bg-success rounded mb-4";
        }
      }
    },
    async getSelectedFolders(id: string) {
      console.log(id);
      await axios.get<Folder[][]>(`/storage/files/folders/${id}`, {
        withCredentials: true
      }).then(response => {
        if (response.data[0]) {
          this.folders[0] = response.data[0]
        }
        else this.folders[0] = [];
        if (response.data[1]) {
          this.folders[1] = response.data[1]
        }
        else this.folders[1] = [];
        // console.log(response.data);
        // this.folders = response.data;
        this.getRandomColor();
      })
    },
    async chooseFolder(index: number, folder: Folder, id: string) {
      folder.iscurr = true;
      this.folders[0].push(folder);
      this.folders[1].splice(index, 1);
      try {
        await axios.put("/storage/files/folder", {
          folder_id: folder.id,
          file_id: id,
        }, {
          withCredentials: true,
        })
      } catch (error) {
        console.log(error);
      }
    },
    async deleteFolder(index: number, folder: Folder, id: string) {
      try {
        folder.iscurr = false;
        this.folders[1].push(folder);
        this.folders[0].splice(index, 1);
        this.getRandomColor();
        const response = await axios.get(`/storage?folder=${folder.id}&file=${id}`, {
          withCredentials: true
        });
        await axios.delete(`/storage/folders/member/${response.data.id}`, {
          withCredentials: true
        })

      } catch (error) {
        console.log(error);
      }
    },




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
    async downloadFile(file: FileServ) {
      const response = await axios.get(file.url, { responseType: "blob" });
      const blob = new Blob([response.data], { type: "application/pdf" });
      const link = document.createElement("a");
      link.href = URL.createObjectURL(blob);
      link.download = file.title;
      link.click();
      URL.revokeObjectURL(link.href);
      await axios.put(`/profile/downloads`, {}, {
        withCredentials: true
      });
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
        if(response.data) {
        this.files = response!.data;
        } else this.files = [];
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
      }, {
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
    this.folders.push([]);
    this.folders.push([]);
    await this.getAllImages();
  },
});
</script>

<style lang="scss">
.dropdown-item:hover {
  cursor: pointer;
}

.ojhjs {
  margin-bottom: -2px !important;
}

.modal-footer {
  display: flex;
  justify-content: space-around;
}

.modal-body {
  padding-left: 30px;
}

.txt-hide {
  width: 350px;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.jstfc {
  display: flex;
  flex-wrap: wrap;
  border-radius: 10px;
  margin-left: 30px;
  margin-right: 30px;
  padding-top: 10px;
  margin-bottom: 10px;
  background-color: #FAFAFA !important;
}

.folder-item {
  padding: 5px 20px;
  margin: 0px 10px;
  display: flex;
}

.fold :hover {
  cursor: pointer;
}


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