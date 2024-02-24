<script setup lang="ts">
import {ref} from 'vue'
import {login} from "@/api/api.ts";

const username = ref("")
const password = ref("")

const handleLogin = () => {
  const form: any = {
    username :username.value,
    password:password.value
  }
  // 在这里添加登录逻辑
  // 例如，发送请求到后端验证用户名和密码
  login(form).then(res=> {
    if (res.code== 200) {
      console.log(res)
    } else {
      console.log(res.msg)
    }
  })


}
</script>

<template>
  <el-card class="sign-card flex justify-center items-center h-4/6 w-1/4" shadow="hover">
    <el-row>
      <el-col :span="12" >
          <img  class="sigin-logo" src="@/assets/TEAMlogo.png">
      </el-col>
      <el-col :span="12">
        <el-button class="elsignup" type="text" style="font-size: 15px">No account?<br>Sign up</el-button>
      </el-col>
    </el-row>
    <div class="font-bold mb-5 flex left-0">
      <el-text style="font-size: 30px">Sign in</el-text>
    </div>
    <el-form label-position="top" class="font-bold">
      <el-form-item label="用户名"  >
        <el-input v-model="username"></el-input>
      </el-form-item>
      <el-form-item label="密码" class="mt-5">
        <el-input v-model="password"></el-input>
        <el-button type="text" style="font-size: 12px;">Forget Password?</el-button>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="handleLogin">登陆</el-button>
      </el-form-item>
    </el-form>
  </el-card>
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

</style>