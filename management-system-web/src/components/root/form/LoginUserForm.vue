<template>
  <el-form
    ref="form"
    :model="form"
    :label-position="labelPosition"
    label-width="120px"
  >
    <el-form-item label="账号">
      <el-input
        placeholder="请输入账号"
        v-model="form.id"
        onkeyup="this.value = this.value.replace(/[^\d.]/g,'');"
      ></el-input>
    </el-form-item>
    <el-form-item label="密码">
      <el-input
        placeholder="请输入密码"
        show-password
        v-model="form.password"
      ></el-input>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="onSubmit">登陆</el-button>
    </el-form-item>
  </el-form>
</template>

<script>
export default {
  data() {
    return {
      form: {
        id: "",
        password: "",
      },
    };
  },
  methods: {
    onSubmit() {},
  },
  beforeMount() {},
};
</script>

<style>
</style>
