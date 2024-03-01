<script setup lang="ts">
import {useRoute} from "vue-router";
import {useAuthStore} from "@/stores/modules/auth.ts";
import {computed} from "vue";
import SubMenu from "@/layouts/components/Menu/SubMenu.vue";
import {useGlobalStore} from "@/stores/modules/global.ts";
import ToolBarLeft from "@/layouts/components/Header/ToolBarLeft.vue";
import ToolBarRight from "@/layouts/components/Header/ToolBarRight.vue";
import Main from "@/layouts/components/Main/index.vue";
const route = useRoute();
const authStore = useAuthStore();
const globalStore = useGlobalStore();
const accordion = computed(() => globalStore.accordion);
const isCollapse = computed(() => globalStore.isCollapse);
console.log(isCollapse)
const menuList = computed(() => authStore.showMenuListGet);
console.log(menuList)
const activeMenu = computed(() => (route.meta.activeMenu ? route.meta.activeMenu : route.path) as string);
</script>

<template>
  <el-container class="layout">
    <el-header >
      <div class="header-lf mask-image">
        <div class="logo flx-center">
          <img class="logo-img ml-5" src="@/assets/TEAMlogo.png"/>
          <span class="logo-text">Aria</span>
        </div>
        <ToolBarLeft/>
      </div>
      <div>
        <ToolBarRight />
      </div>
    </el-header>
    <el-container class="classic-content">
      <el-aside>
        <div class="aside-box" :style="{width: isCollapse ? '65px' : '150px'}">
          <el-scrollbar>
            <el-menu
                :router="false"
                :default-active="activeMenu"
                :collapse=isCollapse
                :unique-opened=accordion
                :collapse-transition="false"
            >
              <SubMenu :menu-list="menuList"/>
            </el-menu>
          </el-scrollbar>
        </div>
      </el-aside>
      <el-container class="classic-main">
        <Main />
      </el-container>
    </el-container>
  </el-container>
</template>

<style scoped>
@import "index.scss";


</style>