<script setup lang="ts">
import {ref, onMounted} from 'vue';

// 用于getScale()函数
const fixedWidth = ref(document.documentElement.clientWidth.toString());
const fixedHeight = ref(document.documentElement.clientHeight.toString());

// 默认行内样式
const style = ref({
  'width': `${fixedWidth.value}px`,
  'height': `${fixedHeight.value}px`,
  'transform': 'scale(1) translate(-50%, -50%)',
})

onMounted(() => {
  setScale(); // 初始化时就执行一次setScale
  window.onresize = debounce(setScale, 1000); // 监听onresize事件，执行setSacle
})

// const debounce0 = (fn: any, delay: any) => {
//   let timer: any;
//   console.log('执行了debounce'); // 只执行一次，下面紧邻函数内的代码会在每次触发resize事件时执行
//   return function () {
//     // const args = arguments; // fn函数（即setScale函数）不必使用args作为fn的入参，
//     // apply()方法作用：https://blog.csdn.net/weixin_43877799/article/details/120282509
//     if (timer) {
//       clearTimeout(timer);
//     }
//     const context = this;
//     timer = setTimeout(() => {
//       timer = null;
//       console.log('正在执行自适应');
//       fn.apply(context); // 或者可以直接运行fn();
//     }, delay);
//   }
// }
const debounce = (fn: Function, delay: number) => {
  let timer: number | null = null;
  console.log('执行了debounce');
  return (...args: any[]) => {
    if (timer !== null) {
      clearTimeout(timer);
    }
    timer = window.setTimeout(() => {
      timer = null;
      console.log('正在执行自适应');
      fn(...args); // 直接传递参数给fn
    }, delay);
  }
}


// 获取放大缩小比例
const getScale = () => {
  const w = window.innerWidth / Number(fixedWidth.value);
  const h = window.innerHeight / Number(fixedHeight.value);
  return w < h ? w : h;
}

// 修改样式
const setScale = () => {
  style.value.transform = `scale${getScale()})) translate(-50% -50%)`;
  // 宽高自适应，window.innerWidth见文章：https://juejin.cn/post/6844903598489337869
  style.value.width = `${(window.innerWidth).toString()}px`; // 赋值时不要使用document.documentElement.clientWidth
  style.value.height = `${(window.innerHeight).toString()}px`;// 赋值时不要使用document.documentElement.clientHeight
  console.log('自适应后的style', style.value);
}

</script>

<template>
  <div class="screen-adapter-style" :style="style">
    <slot/>
  </div>
</template>

<style scoped>
.screen-adapter-style {
  transform-origin: 0 0;
  position: absolute;
  left: 50%;
  top: 50%;
  transition: 0.3s;
  background: transparent;
}

</style>