<script setup lang="ts">
// import VueCompareImage from 'vue3-compare-image';

import {onMounted, onUnmounted, reactive, ref} from 'vue';

// const images = ["./1.jpg", "./2.jpg"]


const state = reactive({
  offsetX: 0
});

const slider = ref<HTMLElement | null>(null);
const coverImage = ref<HTMLElement | null>(null);

const startChange = (event: MouseEvent) => {
  if (slider.value) {
    state.offsetX = event.pageX - slider.value.offsetLeft;
    window.addEventListener("mousemove", moveHandler);
    window.addEventListener("mouseup", endChange);
  }
};

const moveHandler = (e: MouseEvent) => {
  if (slider.value && coverImage.value) {
    const newLeft = e.pageX - state.offsetX - 4;
    slider.value.style.left = `${newLeft || 0}px`;
    coverImage.value.style.width = `${newLeft || 0}px`;
  }
};

const endChange = () => {
  window.removeEventListener("mousemove", moveHandler);
  window.removeEventListener("mouseup", endChange);
};

onMounted(() => {
  window.addEventListener("mousedown", startChange);
});

onUnmounted(() => {
  window.removeEventListener("mousedown", startChange);
  // It's good practice to also remove these listeners, just in case.
  window.removeEventListener("mousemove", moveHandler);
  window.removeEventListener("mouseup", endChange);
});


;
</script>

<template>
  <div class="dehaze">
    <div class="image-comparison" ref="box">
      <div class="cover-image" ref="coverImage">
        <img src="./1.jpg" alt=""/>
      </div>
      <div class="slider" ref="slider" @mousedown="startChange">
        <span class="tag-icon"></span>
      </div>
      <img class="base-image" src="./2.jpg" alt=""/>
    </div>
    <el-container class="w-full">
      <el-header class="card u-head">
        <div class="ml-5 ">
          <el-icon class="top-0.5">
            <UploadFilled/>
          </el-icon>
          <span class="ml-2 font-bold">图像去雾</span>
        </div>
      </el-header>
      <el-main class="mt-5 card u-main w-full">
        <div>
          <!--          <compare-image  value="45">-->
          <!--            <img slot="first"  src="./1.jpg" alt="">-->
          <!--            <img slot="second" src="./2.jpg" alt="">-->
          <!--          </compare-image>-->
          <!--          <VueCompareImage   left-image="./1.jpg" right-image="./2.jpg"/>-->
        </div>
      </el-main>
    </el-container>


  </div>
</template>

<style scoped lang="scss">
.dehaze {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 100%;
  height: 100%;

  :deep(.u-main) {
    position: relative;
    display: flex;
    height: 520px;
    //align-items: start;
  }

  :deep(.u-head) {
    display: flex;
    align-items: center;
    height: 100px;
    color: #303133;
    font-size: 18px;
  }
}

.image-comparison {
  width: 100%;
  height: 100%;
  overflow: hidden;
  position: relative;

  * {
    user-select: none;
  }

  .slider {
    position: absolute;
    left: calc(50% - 4px);
    width: 8px;
    height: 100%;
    cursor: col-resize;
    z-index: 20;
    background: rgba(0, 0, 0, 0.5);

    .tag-icon {
      position: absolute;
      top: 50%;
      left: 50%;
      border: 2px solid rgba(255, 255, 255, 0.5);
      width: 40px;
      height: 40px;
      border-radius: 50%;
      background: rgba(0, 0, 0, 0.5);
      transform: translate(-50%, -50%);
    }
  }

  .cover-image,
  .base-image {
    position: absolute;
    left: 0;
    height: 100%;
    overflow: hidden;
  }

  .cover-image {
    left: 0;
    width: 50%;
    z-index: 20;
  }

  img {
    max-height: 100%;
  }
}

</style>