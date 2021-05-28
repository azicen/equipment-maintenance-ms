<template>
  <el-container :style="defaultHeight">
    <RootNavMenu />
    <el-main>
      <router-view />
    </el-main>
  </el-container>
</template>

<script>
import RootNavMenu from "@/components/root/RootNavMenu.vue";

export default {
  data() {
    return {
      defaultHeight: {
        height: "",
      },
    };
  },
  methods: {
    //定义方法，获取高度减去头尾
    getHeight() {
      this.defaultHeight.height = window.innerHeight - 65 + "px";
    },
  },
  created() {
    //页面创建时执行一次getHeight进行赋值，顺道绑定resize事件
    window.addEventListener("resize", this.getHeight);
    this.getHeight();
  },
  components: {
    RootNavMenu,
  },
};
</script>

<style>
</style>
