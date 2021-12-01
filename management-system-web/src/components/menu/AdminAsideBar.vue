<template>
  <div>
    <el-menu
        :default-active="$route.path"
        :default-openeds="[]"
        class="admin-aside-bar"
        text-color="#fff"
        active-text-color="#aca885"
        background-color="#626c66"
        router
    >

      <el-menu-item index="/admin/user">
        <i class="ri-user-3-fill"></i>
        <span>用户管理</span>
      </el-menu-item>

      <el-menu-item index="/admin/equipment">
        <i class="ri-user-3-fill"></i>
        <span>设备管理</span>
      </el-menu-item>

      <el-menu-item index="/admin/group">
        <i class="ri-user-3-fill"></i>
        <span>权限组管理</span>
      </el-menu-item>

      <el-menu-item index="/admin/maintain">
        <i class="ri-user-3-fill"></i>
        <span>维护信息管理</span>
      </el-menu-item>

    </el-menu>
  </div>
</template>

<script lang="ts">
import {defineComponent} from 'vue'

export default defineComponent({
  name: 'AdminAsideBar',
  components: {},
  setup() {
  }
})
</script>

<style>
i {
  margin: 5%;
  font-size: 1.25em;
}
.admin-aside-bar{
  background-color: #626c66;
}
</style>