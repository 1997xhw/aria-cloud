<script setup lang="ts">
import {reactive, ref} from 'vue'
import {login} from "@/api/api.ts";
import {ElNotification} from "element-plus";
import type {ElForm} from 'element-plus'
import {useUserStore} from "@/stores/modules/user.ts";
import router from "@/routes";
import {HOME_URL} from "@/config";


const userStore = useUserStore();


interface LoginForm {
  username: string
  password: string
}

type FormInstance = InstanceType<typeof ElForm>
const loginFormRef = ref<FormInstance>()
const loginForm = reactive<LoginForm>({
  username: "",
  password: ""

})
const rules = reactive({
  username: [
    {required: true, message: 'Please input Activity name', trigger: 'blur'},
    // { min: 3, max: 5, message: 'Length should be 3 to 5', trigger: 'blur' },
  ],
  password: [
    {required: true, message: 'Please input Password', trigger: 'blur'},
  ]
})

const loading = ref(false);


const handleLogin = async (formEl: FormInstance | undefined) => {
  if (!formEl) return
  await formEl.validate((vaild, fields) => {
    if (vaild) {
      const form: any = {
        username: loginForm.username,
        password: loginForm.password
      }
      // 在这里添加登录逻辑
      // 例如，发送请求到后端验证用户名和密码
      try {
        loading.value = true
        login(form).then(res => {
          console.log("res: ", res);

          if (res.status == 200) {
            let data = res.data.data;
            userStore.setToken(data.token);
            userStore.setUserInfo(data.username);
            ElNotification({
              title: 'Success',
              message: '登陆成功！',
              type: 'success',
            })
            // console.log(res)
            router.push(HOME_URL)
          } else {
            ElNotification({
              title: 'Error',
              message: res.msg,
              type: 'error',
            })
            // console.log(res.msg)
          }
        })
      } finally {
        loading.value = false
      }
    } else {
      console.log(fields)
      return
    }
  })
}
</script>

<template>
  <el-container class="login-container flx-center">
    <div class="login-box">
      <el-card class="login-card" shadow="hover">
        <div class="team-logo">
          <img class="team-icon" src="../../assets/TEAMlogo.png">
          <el-button class="signup-button" type="text" @click="$router.push('/register')">No
            account?<br>Sign up
          </el-button>
        </div>

        <div class="font-bold mb-5 flex left-0">
          <el-text style="font-size: 30px">登陆</el-text>
        </div>
        <el-form label-position="top"
                 class="font-bold"
                 ref="loginFormRef"
                 :rules="rules"
                 :hide-required-asterisk="true"
                 :model="loginForm"
        >
          <el-form-item label="用户名" prop="username">
            <el-input v-model="loginForm.username"></el-input>
          </el-form-item>
          <el-form-item label="密码" class="mt-5" prop="password">
            <el-input v-model="loginForm.password"></el-input>

          </el-form-item>
          <div class="">
            <div class="text-right">
              <el-button type="text" style="font-size: 12px;">Forget Password?</el-button>
            </div>
            <div class="text-left mb-5">
              <el-button type="primary" @click="handleLogin(loginFormRef)" :loading="loading">登陆</el-button>
            </div>
          </div>
        </el-form>
      </el-card>
    </div>
  </el-container>
</template>
<style scoped>


@import "index.scss";

</style>