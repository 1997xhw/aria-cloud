<script setup lang="ts">
import {reactive, ref} from 'vue'
import {register} from "@/api/api.ts";
import {ElNotification, FormRules} from "element-plus";
import type {ElForm} from 'element-plus'

import {useRouter} from 'vue-router';

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
const rules = reactive({
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
  if (!formEl) return
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
  <el-container class="login-container flx-center">
    <div class="login-box">
      <el-card class="login-card " shadow="hover">

          <div class="team-logo">
            <img class="team-icon" src="../../assets/TEAMlogo.png">
          </div>

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
          <div class=" text-left mb-5">
            <el-button type="primary" @click="handleRegister(registerFormRef)">注册</el-button>
          </div>
        </el-form>
      </el-card>
    </div>
  </el-container>
</template>

<style scoped>
@import "./index.css";
</style>