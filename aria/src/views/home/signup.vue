<script setup lang="ts">
import {reactive, ref} from 'vue'
import {register} from "@/api/api.ts";
import {ElNotification, FormRules} from "element-plus";
import type {ElForm} from 'element-plus'

import { useRouter } from 'vue-router';
const router = useRouter();
type FormInstance = InstanceType<typeof ElForm>

interface RegisterForm {
  username: string
  password: string
}

const registerFormRef = ref<FormInstance>()
const registerForm = reactive<RegisterForm>({
  username: "",
  password: ""

})
const rules = reactive<FormRules<RegisterForm>>({
  username: [
    {required: true, message: 'Please input Activity name', trigger: 'blur'},
    {min: 3, max: 5, message: 'Length should be 3 to 5', trigger: 'blur'},
  ],
  password: [
    {required: true, message: 'Please input Password', trigger: 'blur'},
    {min: 5, max: 12, message: 'Length should be 3 to 12', trigger: 'blur'},
  ]
})


const handleRegister = async (formEl: FormInstance) => {
  if (!registerFormRef) return
  await formEl.validate((vaild) => {
    if (vaild) {
      const form: any = {
        username: registerForm.username,
        password: registerForm.password
      }
      // 在这里添加登录逻辑
      // 例如，发送请求到后端验证用户名和密码
      register(form).then(res => {
        if (res.code == 200) {
          ElNotification({
            title: 'Success',
            message: '注册成功！',
            type: 'success',
          })

          console.log(res)
          router.push('/login')
        } else {
          ElNotification({
            title: 'Error',
            message: res.msg,
            type: 'error',
          })
          console.log(res.msg)
        }
      })
    }
  })


}
</script>

<template>
  <el-container class="main">
    <el-card class="sign-card flex justify-center items-center h-4/6 w-1/4" shadow="hover">
      <el-row>
        <el-col :span="12">
          <img class="sigin-logo" src="@/assets/TEAMlogo.png">
        </el-col>
      </el-row>
      <div class="font-bold mb-5 flex left-0">
        <el-text style="font-size: 30px">注册</el-text>
      </div>
      <el-form label-position="top"
               class="font-bold"
               ref="registerFormRef"
               :rules="rules"
               :hide-required-asterisk="true"
               :model="registerForm"
      >
        <el-form-item label="用户名" prop="username">
          <el-input v-model="registerForm.username"></el-input>
        </el-form-item>
        <el-form-item label="密码" class="mt-5" prop="password">
          <el-input v-model="registerForm.password"></el-input>
        </el-form-item>

        <div class="text-right">
          <el-button type="text" style="font-size: 12px;" @click="$router.push('/login')">已有账号?</el-button>
        </div>
          <div  class=" text-left mb-5">
            <el-button type="primary" @click="handleRegister(registerFormRef)">注册</el-button>
          </div>



      </el-form>
    </el-card>
  </el-container>
</template>

<style scoped>
.sign-card {
  position: relative;
  left: 8%;
  top: 12%;
  padding: 10px 10px 20px 10px;

  border-radius: 30px; /* 设置圆角大小 */
  background-color: rgba(255, 255, 255, 0.7); /* 设置背景透明 */
  border-color: rgba(255, 255, 255, 0);
}

.elsignup {
  text-align: right;
  position: relative;
  right: -20px;
  top: 60%;
}

.sigin-logo {
  position: relative;
  left: 0;
  top: 10%;

}

.main {
  background-image: url('@/assets/img/bg.jpg');
  background-size: cover;
  background-position: center;
  background-attachment: fixed;
  width: 100%;
  height: 100vh;
}
</style>