<script setup lang="ts">

import Signin from "@/views/home/signin.vue";
</script>

<template>
  <div >
  <el-container class="main">
    <signin></signin>
  </el-container>
  </div>
</template>

<style scoped>
.main {
  background-image: url('@/assets/img/bg.jpg');
  background-size: cover;
  background-position: center;
  background-attachment: fixed;
  width: 100%;
  height: 100vh;
}
</style>