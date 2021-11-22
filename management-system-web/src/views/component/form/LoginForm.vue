<template>
  <div>
    <el-card class="login-form-box-card">
      <el-form ref="form" :model="form">
        <el-form-item label="账户">
          <el-input v-model="form.account" placeholder="account"/>
        </el-form-item>
        <el-form-item label="密码">
          <el-input v-model="form.password" placeholder="password" show-password/>
        </el-form-item>
        <div class="login-form-button">
          <el-form-item>
            <el-button type="primary" @click="onSubmit">登录</el-button>
          </el-form-item>
        </div>
      </el-form>
    </el-card>
  </div>
</template>

<script>
export default {
  data() {
    return {
      form: {
        account: '',
        password: '',
      },
    }
  },
  methods: {
    onSubmit() {
      console.log('submit!')
    },
  },
}
</script>

<style>

.login-form-box-card {
  margin: 0 25%;
  width: 480px;
}

.login-form-button {
  margin-left: 70%;
}
</style>