<template>
  <el-menu
    :default-active="activeIndex"
    class="el-menu-demo"
    mode="horizontal"
    @select="handleSelect"
    router
  >
    <el-menu-item index="/">首页</el-menu-item>
    <el-menu-item index="/maintain">设备维护中心</el-menu-item>
    <el-menu-item index="/root_home">管理系统</el-menu-item>
  </el-menu>
</template>

<script>
export default {
  name: "WebNavMenu",
  props: {
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
