<template>
<div class="container-fluid">
            <div class="row">
                <div class="col-lg-12">
                    <div class="card card-block card-stretch card-transparent ">
                        <div class="card-header d-flex justify-content-between pb-0">
                            <div class="header-title">
                                <h4 class="card-title">Documents</h4>
                            </div>
                            <div class="card-header-toolbar d-flex align-items-center">
                                <div class="card-header-toolbar">
                                    <div class="dropdown">
                                        <span class="dropdown-toggle dropdown-bg btn bg-white" id="dropdownMenuButton1"
                                            data-toggle="dropdown">
                                            Name<i class="ri-arrow-down-s-line ml-1"></i>
                                        </span>
                                        <div class="dropdown-menu dropdown-menu-right shadow-none"
                                            aria-labelledby="dropdownMenuButton1">
                                            <a class="dropdown-item" href="#">Last modified</a>
                                            <a class="dropdown-item" href="#">Last modifiedby me</a>
                                            <a class="dropdown-item" href="#">Last opened by me</a>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="col-lg-3 col-md-6 col-sm-6">
                    <div class="card card-block card-stretch card-height">
                        <div class="card-body image-thumb">
                            <a href="#" data-title="Terms.pdf" data-load-file="file" data-load-target="#resolte-contaniner" data-url="../../assets/vendor/doc-viewer/files/demo.pdf" data-toggle="modal" data-target="#exampleModal">
                                <div class="mb-4 text-center p-3 rounded iq-thumb">
                                    <div class="iq-image-overlay"></div>
                                    <img src="../../assets/images/layouts/page-1/pdf.png" class="img-fluid" alt="image1">       
                                </div>
                                <h6>Terms.pdf</h6> 
                            </a>             
                        </div>
                    </div>
                </div>
                <div class="col-lg-3 col-md-6 col-sm-6">
                    <div class="card card-block card-stretch card-height">
                        <div class="card-body image-thumb">
                            <a href="#" data-title="New-one.docx" data-load-file="file" data-load-target="#resolte-contaniner" data-url="../../assets/vendor/doc-viewer/files/demo.docx" data-toggle="modal" data-target="#exampleModal">
                            <div class="mb-4 text-center p-3 rounded iq-thumb">
                                <div class="iq-image-overlay"></div>
                                <img src="../../assets/images/layouts/page-1/doc.png" class="img-fluid" alt="image1">
                            </div>
                            <h6>New-one.docx</h6>  
                            </a>   
                        </div>
                    </div>
                </div>
                <div class="col-lg-3 col-md-6 col-sm-6">
                    <div class="card card-block card-stretch card-height">
                        <div class="card-body image-thumb">
                            <a href="#" data-title="Woo-box.xlsx" data-load-file="file" data-load-target="#resolte-contaniner"  data-url="../../assets/vendor/doc-viewer/files/demo.xlsx" data-toggle="modal" data-target="#exampleModal">
                                <div class="mb-4 text-center p-3 rounded iq-thumb">
                                    <div class="iq-image-overlay"></div>
                                    <img src="../../assets/images/layouts/page-1/xlsx.png" class="img-fluid" alt="image1">
                                </div>
                                <h6>Woo-box.xlsx</h6>
                            </a>
                        </div>
                    </div>
                </div>
                <div class="col-lg-3 col-md-6 col-sm-6">
                    <div class="card card-block card-stretch card-height">
                        <div class="card-body image-thumb doc-text">
                            <a href="#" data-title="IOS-content.pptx" data-load-file="file" data-load-target="#resolte-contaniner" data-url="../../assets/vendor/doc-viewer/files/demo.pptx" data-toggle="modal" data-target="#exampleModal">
                            <div class="mb-4 text-center p-3 rounded iq-thumb">
                                <div class="iq-image-overlay"></div>
                                <img src="../../assets/images/layouts/page-1/ppt.png" class="img-fluid" alt="image1">      
                            </div>
                            <h6>IOS-content.pptx</h6>   
                            </a>       
                        </div>
                    </div>
                </div>
                <div class="col-lg-12">
                    <div class="card card-block card-stretch card-transparent">
                        <div class="card-header d-flex justify-content-between pb-0">
                            <div class="header-title">
                                <h4 class="card-title">Folders</h4>
                            </div>
                            <div class="card-header-toolbar">
                                <div class="dropdown">
                                    <span class="dropdown-toggle dropdown-bg btn bg-white" id="dropdownMenuButton1"
                                        data-toggle="dropdown">
                                        Name<i class="ri-arrow-down-s-line ml-1"></i>
                                    </span>
                                    <div class="dropdown-menu dropdown-menu-right shadow-none"
                                        aria-labelledby="dropdownMenuButton1">
                                        <a class="dropdown-item" href="#">Last modified</a>
                                        <a class="dropdown-item" href="#">Last modifiedby me</a>
                                        <a class="dropdown-item" href="#">Last opened by me</a>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>                
                
                <div class="col-md-6 col-sm-6 col-lg-3"  v-for="(folder, index) in Folders" :key="index" >
                    <div class="card card-block card-stretch card-height">
                        <div class="card-body">
                            <div class="d-flex justify-content-between">
                                <router-link :to="folder.link" class="folder">
                                        <div :class="folder.color">
                                            <i class="ri-file-copy-line"></i>
                                        </div>
                                    </router-link>
                                <div class="card-header-toolbar">
                                    <div class="dropdown">
                                        <span class="dropdown-toggle" id="dropdownMenuButton5" data-toggle="dropdown">
                                            <i class="ri-more-2-fill"></i>
                                        </span>
                                        <div class="dropdown-menu dropdown-menu-right" aria-labelledby="dropdownMenuButton5">
                                            <a class="dropdown-item" href="#"><i class="ri-eye-fill mr-2"></i>View</a>
                                            <a class="dropdown-item" href="#"><i class="ri-delete-bin-6-fill mr-2"></i>Delete</a>
                                            <a class="dropdown-item" href="#"><i class="ri-pencil-fill mr-2"></i>Edit</a>
                                            <a class="dropdown-item" href="#"><i class="ri-printer-fill mr-2"></i>Print</a>
                                            <a class="dropdown-item" href="#"><i class="ri-file-download-fill mr-2"></i>Download</a>
                                        </div>
                                    </div>
                                </div>
                            </div>
                                <router-link :to="folder.link" class="folder">
                                    <h5 class="mb-2">{{folder.name}}</h5>
                                    <p class="mb-2"><i :class="folder.clock "></i>{{folder.date}}</p>
                                    <p class="mb-0"><i :class="folder.filecolor"></i> {{folder.file}}</p>
                                </router-link>
                        </div>
                    </div>
                </div>
                <div class="col-lg-12">
                    <div class="card card-block card-stretch card-transparent">
                        <div class="card-header d-flex justify-content-between pb-0">
                            <div class="header-title">
                                <h4 class="card-title">Files</h4>
                            </div>
                        </div>
                    </div>
                </div>
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
                                    <tr v-for="(file, index) in  Files" :key="index">
                                        <td>
                                            <div class="d-flex align-items-center">
                                                <div :class="file.color">
                                                    <i :class="file.icon"></i>
                                                </div>
                                                <a href="#"  :data-title="file.name" :data-url="file.url" @click="$root?.$emit('bv::show::modal', 'viewer-modal', $event.target)" v-b-modal.viewer-modal >
                                                <span style="color:#535F6B;">{{file.name}}</span>
                                                </a>
                                            </div>
                                        </td>
                                        <td>{{file.owner}}</td>
                                        <td>{{file.lastEdit}}</td>
                                        <td>{{file.size}}</td>
                                        <td>
                                            <div class="dropdown">
                                                <span class="dropdown-toggle" id="dropdownMenuButton12" data-toggle="dropdown">
                                                    <i class="ri-more-fill"></i>
                                                </span>
                                                <div class="dropdown-menu dropdown-menu-right" aria-labelledby="dropdownMenuButton12">
                                                    <a class="dropdown-item" href="#"><i class="ri-eye-fill mr-2"></i>View</a>
                                                    <a class="dropdown-item" href="#"><i class="ri-delete-bin-6-fill mr-2"></i>Delete</a>
                                                    <a class="dropdown-item" href="#"><i class="ri-pencil-fill mr-2"></i>Edit</a>
                                                    <a class="dropdown-item" href="#"><i class="ri-printer-fill mr-2"></i>Print</a>
                                                    <a class="dropdown-item" href="#"><i class="ri-file-download-fill mr-2"></i>Download</a>
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
</template>

<script lang="ts">
import { defineComponent } from "@vue/runtime-core";

export default defineComponent({
  name: "Recent",
  data() {
    return {
      Folders: [
        {
          link: "./alexworkshop",
          color: "icon-small bg-danger rounded mb-4 ",
          name: "Alexa Workshop",
          date: "10 Dec, 2020",
          file: "08 Files",
          clock: "lar la-clock text-danger mr-2 font-size-20",
          filecolor: "las la-file-alt text-danger mr-2 font-size-20",
        },
        {
          link: "./android",
          color: "icon-small bg-primary rounded mb-4 ",
          name: "Android",
          date: "09 Dec, 2020",
          file: "08 Files",
          clock: "lar la-clock text-primary mr-2 font-size-20",
          filecolor: "las la-file-alt text-primary mr-2 font-size-20",
        },
        {
          link: "./brightspot",
          color: "icon-small bg-info rounded mb-4 ",
          name: "Brightspot",
          date: " 07 Dec, 2020",
          file: "08 Files",
          clock: "lar la-clock text-info mr-2 font-size-20",
          filecolor: "las la-file-alt text-info mr-2 font-size-20",
        },
        {
          link: "./ionic-chat-app",
          color: "icon-small bg-success rounded mb-4 ",
          name: "Ionic Chat App",
          date: " 06 Dec, 2020",
          file: "08 Files",
          clock: "lar la-clock text-success mr-2 font-size-20",
          filecolor: "las la-file-alt text-success mr-2 font-size-20",
        },
      ],
      Files: [
        {
          icon: "ri-file-excel-line",
          color: "icon-small bg-danger rounded mr-3",
          name: "Weekly Report.pdf",
          url: this.baseUrl + "/doc-viewer/files/demo.pdf",
          owner: "Me",
          lastEdit: "Mar 30, 2020 Gail Forcewind",
          size: "10 MB",
        },
        {
          icon: "ri-file-download-line",
          color: "icon-small bg-primary rounded mr-3",
          name: "Milestone.docx",
          url: this.baseUrl + "/doc-viewer/files/demo.docx",
          owner: "Penny",
          lastEdit: "Mar 31, 2020 Penny",
          size: "65 MB",
        },
        {
          icon: "ri-file-excel-line",
          color: "icon-small bg-info rounded mr-3",
          name: "Training center.xlsx",
          url: this.baseUrl + "/doc-viewer/files/demo.xlsx",
          owner: "Banny",
          lastEdit: "Mar 30, 2020 Banny",
          size: "90 MB",
        },
        {
          icon: "ri-file-excel-line",
          color: "icon-small bg-success rounded mr-3",
          name: "Training center.pptx",
          url: this.baseUrl + "/doc-viewer/files/demo.pptx",
          owner: "Me",
          lastEdit: "Apr 04, 2020 me",
          size: "10 MB",
        },
      ],
    };
  },
});
</script>