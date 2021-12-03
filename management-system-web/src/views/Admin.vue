<template>
  <el-header class="top-bar">
    <AdminTopBar/>
  </el-header>
  <el-container>
    <el-aside class="aside-bar">
      <AdminAsideBar/>
    </el-aside>
    <el-main>
      <router-view/>
    </el-main>
  </el-container>
</template>

<script lang="ts">
import {defineComponent} from 'vue'
import AdminTopBar from "@/components/menu/AdminTopBar.vue"
import AdminAsideBar from "@/components/menu/AdminAsideBar.vue"

export default defineComponent({
  name: 'Admin',
  components: {
    AdminTopBar,
    AdminAsideBar,
  },
  setup() {
  }
})
</script>

<style>
html, body, #app {
  width: 100%;
  height: 100%;
  margin: 0;
  padding: 0;
}

.top-bar {
  border: 1px #434a42 solid;
  background-color: #434a42;
  width: 100%;
}

.aside-bar {
  background-color: #626c66;
  width: 15%;
  min-width: 150px;
  max-width: 300px;
  height: 100%;
  /*background-color: rgb(255, 255, 255);*/
}
</style>