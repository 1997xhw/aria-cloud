<script setup lang="ts">
import {ref} from 'vue'
import {genFileId, UploadUserFile, ElMessageBox, ElNotification} from 'element-plus'
import type {UploadInstance, UploadProps, UploadRawFile} from 'element-plus'
import {useUserStore} from "@/stores/modules/user.ts";
import { uploadFile} from "@/api/api.ts";

const upload = ref<UploadInstance>()
const loading = ref(false);

const fileList = ref<UploadUserFile[]>([])
const handleExceed: UploadProps['onExceed'] = (files) => {
  upload.value!.clearFiles()
  const file = files[0] as UploadRawFile
  file.uid = genFileId()
  upload.value!.handleStart(file)
}
const userStore = useUserStore();
let formData = new FormData()
const submitUpload = () => {
  // upload.value!.submit()
  formData.append("username", userStore.username)
  formData.append("token", userStore.token)
  if(formData.get("file")==null) {
    ElNotification({
      title: '上传状态',
      message: '请选择至少一个文件！',
      type: 'warning',
    })
    return
  }
  try {
    loading.value=true;
    uploadFile(formData).then(res => {
      if (res.status == 200) {
        ElNotification({
          title: '上传状态',
          message: '成功',
          type: 'success',
        })
      } else {
        ElNotification({
          title: '上传状态',
          message: res.msg,
          type: 'error',
        })
      }
    })
  }finally {
    loading.value=false;
  }

}

// const progressShu = ref(0) //控制进度条
// const ulpadFile = async () => {
//   formData.append("username", userStore.username)
//   formData.append("token", userStore.token)
//
//   uploadFile(formData).then(res => {
//     if (res.code == 200) {
//       console.log(res.msg)
//       console.log(res)
//     } else {
//       console.log(res.msg)
//     }
//   })
//
//   // isShow.value = false;
//
// }

const handleRemove: UploadProps['onRemove'] = (file, uploadFiles) => {
  console.log(file, uploadFiles)
}

// const handlePreview: UploadProps['onPreview'] = (uploadFile) => {
//   console.log(uploadFile.name)
// }
const beforeRemove: UploadProps['beforeRemove'] = (uploadFile) => {
  return ElMessageBox.confirm(
      `Cancel the transfer of ${uploadFile.name} ?`
  ).then(
      () => true,
      () => false
  )
}
// const handleChange = (resp:any, file, fileList) => {
//   if (resp.code === 200) {
//     // 上传成功
//     console.log("success")
//   } else {
//     // 上传失败
//     console.log(resp.msg)
//     // 清除上传列表中的该文件
//     fileList.splice(fileList.indexOf(file), 1)
//   }
//
//
// }


const handleChange = (file, fileList) => {
  console.log(file)
  formData.append("file", file.raw)
  if (file.status !== 'uploading') {
    console.log(file, fileList)
    if (file.status == 'done') {
      console.log(file.response.msg)
      // formState.filePath = file.response.data
      return
    }
    // if (file.status == 'removed') {
    //   removeFile({data: filePath.value}).then(res => {
    //     if (res.code == 200) {
    //       message.success(res.msg)
    //     } else {
    //       message.error(res.msg)
    //     }
    //   })
    //   return
    // }
  }
};

</script>

<template>
  <div class="upload">
    <el-container class="w-full">
      <el-header class="card u-head">
        <div class="ml-5 ">
          <el-icon class="top-0.5">
            <UploadFilled/>
          </el-icon>
          <span class="ml-2 font-bold">上传文件</span>
        </div>
      </el-header>
      <el-main class="mt-5 card u-main w-full">
        <el-upload
            ref="upload"
            v-model:file-list="fileList"
            action=""

            :on-change="handleChange"
            :on-remove="handleRemove"
            :before-remove="beforeRemove"
            :max-count="1"
            :limit="1"
            :on-exceed="handleExceed"
            :auto-upload="false"
        >

          <template #trigger>
            <el-button>选择文件</el-button>
          </template>

          <el-button class="ml-3" type="success" @click="submitUpload" :loading="loading">上传文件</el-button>
          <el-button type="danger">删除文件</el-button>
          <el-button type="warning">终止上传</el-button>
        </el-upload>

      </el-main>
    </el-container>

  </div>
</template>

<style scoped lang="scss">
@import "./index.scss";
</style>