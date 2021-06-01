<template>
  <el-aside width="250px" style="background-color: rgb(255, 255, 255)">
    <el-menu
      :default-active="$route.path"
      router
      @select="handleSelect"
      :default-openeds="[]"
    >
      <el-submenu index="1">
        <template #title><i class="el-icon-user-solid"></i>用户管理</template>
        <el-menu-item-group>
          <el-menu-item index="/root_home/add_user">添加用户</el-menu-item>
          <el-menu-item index="/root_home/put_user">修改用户</el-menu-item>
          <el-menu-item index="/root_home/get_user">用户列表</el-menu-item>
        </el-menu-item-group>
      </el-submenu>
      <el-submenu index="2">
        <template #title><i class="el-icon-menu"></i>设备管理</template>
        <el-menu-item-group>
          <el-menu-item index="/root_home/add_equipment">添加设备</el-menu-item>
          <el-menu-item index="/root_home/put_equipment">修改设备</el-menu-item>
          <el-menu-item index="/root_home/get_equipment">设备列表</el-menu-item>
        </el-menu-item-group>
      </el-submenu>
      <el-submenu index="3">
        <template #title><i class="el-icon-menu"></i>权限组管理</template>
        <el-menu-item-group>
          <el-menu-item index="/root_home/add_group">添加权限组</el-menu-item>
          <el-menu-item index="/root_home/put_group">修改权限组</el-menu-item>
          <el-menu-item index="/root_home/get_group">权限组列表</el-menu-item>
        </el-menu-item-group>
      </el-submenu>
      <el-submenu index="4">
        <template #title><i class="el-icon-menu"></i>设备类型管理</template>
        <el-menu-item-group>
          <el-menu-item index="/root_home/add_equipment_type"
            >添加设备类型</el-menu-item
          >
          <el-menu-item index="/root_home/put_equipment_type"
            >修改设备类型</el-menu-item
          >
          <el-menu-item index="/root_home/get_equipment_type"
            >设备类型列表</el-menu-item
          >
        </el-menu-item-group>
      </el-submenu>
      <el-submenu index="5">
        <template #title><i class="el-icon-menu"></i>维护信息管理</template>
        <el-menu-item-group>
          <el-menu-item index="/root_home/get_maintain"
            >维护信息列表</el-menu-item
          >
        </el-menu-item-group>
      </el-submenu>
    </el-menu>
  </el-aside>
</template>

<script>
export default {
  methods: {},
};
</script>

<style>
</style>
