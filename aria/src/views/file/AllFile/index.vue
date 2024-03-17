<script setup lang="ts">
import {onMounted, reactive} from "vue";
import {getFileAllList} from "@/api/api.ts";
import {ElNotification} from "element-plus";
import {useUserStore} from "@/stores/modules/user.ts";


interface file {
  filename: string,
  filehash: string,
  filesize: string,
  updatetime: string,
}

const tableData = reactive<file[]>([])
// tableData.push({filehash: "asasas", filename: "asasas", filesize: 2130, updatetime: "asdsad"})
const userStore = useUserStore();
const getFileList = ()=> {
  getFileAllList(userStore.token,userStore.username).then(res=> {
    if (res.status==200) {
      console.log(res)
      let data = res.data.data
      data.forEach(item=> {
        tableData.push({
          filename: item.file_name,
          filehash: item.file_sha,
          filesize: (item.file_size/1024/1024).toFixed(2).toString() + " M",
          updatetime: item.last_update,
        })
      })
    }else {
      ElNotification({
        title: '获取数据错误',
        message: res.msg,
        type: 'error',
      })
    }
  })
}

onMounted(()=> {
  getFileList()
})

</script>

<template>
  <div class="all-file">
    <el-container class="w-full">
      <el-header class="card u-head">
        <div class="ml-5 ">
          <el-icon class="top-0.5">
            <UploadFilled/>
          </el-icon>
          <span class="ml-2 font-bold">全部文件</span>
        </div>
      </el-header>
      <el-main class="mt-5 card u-main w-full">
        <el-table height="100%" :data="tableData">
          <el-table-column prop="filename" label="文件名"></el-table-column>
          <el-table-column prop="filehash" label="哈希"></el-table-column>
          <el-table-column prop="filesize" label="文件大小" width="100px"></el-table-column>
          <el-table-column prop="updatetime" label="上传时间"></el-table-column>
          <el-table-column  label="操作" width="200px">
            <el-button type="success" size="default">下载</el-button>
            <el-button type="danger" size="default">删除</el-button>
          </el-table-column>
        </el-table>
      </el-main>
    </el-container>

  </div>
</template>

<style scoped lang="scss">
@import "./index.scss";
</style>