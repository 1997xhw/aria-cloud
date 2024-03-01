<script setup lang="ts">
import {ref} from 'vue'
import {genFileId, UploadUserFile, ElMessageBox} from 'element-plus'
import type {UploadInstance, UploadProps, UploadRawFile} from 'element-plus'

const upload = ref<UploadInstance>()

const fileList = ref<UploadUserFile[]>([])
const handleExceed: UploadProps['onExceed'] = (files) => {
  upload.value!.clearFiles()
  const file = files[0] as UploadRawFile
  file.uid = genFileId()
  upload.value!.handleStart(file)
}

const submitUpload = () => {
  upload.value!.submit()
}
const handleRemove: UploadProps['onRemove'] = (file, uploadFiles) => {
  console.log(file, uploadFiles)
}

const handlePreview: UploadProps['onPreview'] = (uploadFile) => {
  console.log(uploadFile)
}
const beforeRemove: UploadProps['beforeRemove'] = (uploadFile, uploadFiles) => {
  return ElMessageBox.confirm(
      `Cancel the transfer of ${uploadFile.name} ?`
  ).then(
      () => true,
      () => false
  )
}

</script>

<template>
  <div class="upload">
    <el-container class="w-full" >
      <el-header class="card u-head">
        <div class="ml-5 ">
          <el-icon class="top-0.5">
            <UploadFilled/>
          </el-icon>
          <span class="ml-2 font-bold">上传文件</span>
        </div>
      </el-header>
      <el-main class="mt-5 card u-main w-full">
        <div class=" relative  w-full ">
          <el-upload
              v-model:file-list="fileList"
              class="upload-demo w-full"
              action="https://run.mocky.io/v3/9d059bf9-4660-45f2-925d-ce80ad6c4d15"
              multiple
              :on-preview="handlePreview"
              :on-remove="handleRemove"
              :before-remove="beforeRemove"
              :limit="3"
              :on-exceed="handleExceed"
              :auto-upload="false"
          >

              <template #trigger>
                <el-button>选择文件</el-button>
              </template>

              <el-button class="ml-3" type="success" @click="submitUpload">上传文件</el-button>
              <el-button type="danger">删除文件</el-button>
              <el-button type="warning">终止上传</el-button>


          </el-upload>

        </div>

      </el-main>
    </el-container>

  </div>
</template>

<style scoped lang="scss">
@import "./index.scss";
</style>