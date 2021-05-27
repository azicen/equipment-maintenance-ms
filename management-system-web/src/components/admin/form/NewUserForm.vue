<template>
  <el-form
    ref="form"
    :model="form"
    :label-position="labelPosition"
    label-width="120px"
  >
    <el-form-item label="人员姓名">
      <el-input v-model="form.name"></el-input>
    </el-form-item>
    <el-form-item label="分配权限组">
      <el-select
        v-model="form.types"
        multiple
        filterable
        allow-create
        default-first-option
        placeholder="请选权限组"
      >
        <el-option
          v-for="item in allGroups"
          :key="item.id"
          :label="item.name"
          :value="item.id"
        >
        </el-option>
      </el-select>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="onSubmit">立即创建</el-button>
      <el-button>取消</el-button>
    </el-form-item>
  </el-form>
</template>

<script>
import * as GroupAPI from "@/api/group";
import * as UserAPI from "@/api/user";

export default {
  data() {
    return {
      form: {
        name: "",
        password: "123456",
      },
      allGroups: [],
      groups: [],
    };
  },
  methods: {
    getGroupList() {
      GroupAPI.getGroups().then((res) => {
        this.allGroups = res.data.groups;
      });
    },
    addUser() {
      UserAPI.postUser(this.form).then((res) => {
        this.$notify({
          title: "创建成功",
          message: `您索创建的用户ID为${res.data.id}`,
          type: "success",
        });
      });
    },
    onSubmit() {
      this.addUser();
    },
  },
  beforeMount() {
    this.getGroupList();
  },
};
</script>

<style>
</style>
