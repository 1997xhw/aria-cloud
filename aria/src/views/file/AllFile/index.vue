<script setup lang="ts">
import {onMounted, reactive, ref} from "vue";
import {DeleteFileOne, getFileAllList} from "@/api/api.ts";
import {ElMessageBox, ElNotification, ElMessage} from "element-plus";
import {useUserStore} from "@/stores/modules/user.ts";
import router from "@/routes";


interface file {
  filename: string,
  filehash: string,
  filesize: string,
  updatetime: string,
}

const tableData = reactive<file[]>([])
// tableData.push({filehash: "asasas", filename: "asasas", filesize: 2130, updatetime: "asdsad"})
const userStore = useUserStore();
const getFileList = () => {
  tableData.length = 0
  getFileAllList(userStore.token, userStore.username).then(res => {
    if (res.status == 200) {
      console.log(res)
      let data = res.data.data
      data.forEach(item => {
        tableData.push({
          filename: item.file_name,
          filehash: item.file_sha,
          filesize: (item.file_size / 1024 / 1024).toFixed(2).toString() + " M",
          updatetime: item.last_update,
        })
      })
    } else {
      ElNotification({
        title: '获取数据错误',
        message: res.msg,
        type: 'error',
      })
    }
  })
}

onMounted(() => {
  getFileList()
})

const delVis = ref(false)

const DeleteFileHandler = (file: file) => {
  ElMessageBox.confirm(
      '确认删除文件【 '+file.filename+'】？',
      '注意',
      {
        confirmButtonText: 'OK',
        cancelButtonText: 'Cancel',
        type: 'warning',
      }
  ).then(()=>{
    console.log('Deleting file', file)
    delVis.value = false
    let formData = new FormData()
    formData.append("username", userStore.username)
    formData.append("filehash", file.filehash)
    console.log("formData: ", formData.get("username"))
    DeleteFileOne(formData)
        .then(res => {
          if (res.status==200) {
            ElNotification({
              title: '删除状态',
              message: '成功',
              type: 'success',
            })
            tableData.length = 0
            getFileList()
          } else {
            ElNotification({
              title: '上传状态',
              message: res.msg,
              type: 'error',
            })
          }
        })
  }).catch(()=>{
    ElMessage({
      type: 'info',
      message: 'Delete canceled',
    })
  })

}

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
          <el-table-column label="操作" width="200px">
            <template #default="scope">
              <el-button type="success" size="default">下载</el-button>
              <el-button type="danger" size="default" @click="DeleteFileHandler(scope.row)">删除</el-button>

<!--              <el-popover :visible="delVis" placement="top" :width="160">-->
<!--                <p>Are you sure to delete this?</p>-->
<!--                <div style="text-align: right; margin: 0">-->
<!--                  <el-button size="small" text @click="delVis = false">取消</el-button>-->
<!--                  <el-button size="small" type="primary" @click="DeleteFileHandler(scope.row)"-->
<!--                  >确认</el-button-->
<!--                  >-->
<!--                </div>-->
<!--                <template #reference>-->
<!--&lt;!&ndash;                  <el-button @click="visible = true">Delete</el-button>&ndash;&gt;-->
<!--                  <el-button type="danger" size="default" @click="delVis = true">删除</el-button>-->
<!--                </template>-->
<!--              </el-popover>-->
            </template>
          </el-table-column>
        </el-table>
      </el-main>
    </el-container>

  </div>
</template>

<style scoped lang="scss">
@import "./index.scss";
</style>