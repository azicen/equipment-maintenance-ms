<template>
  <div class="login-form">
    <el-form :model="form">

      <el-form-item label="账号">
        <el-input v-model.number="form.id" placeholder="请输入账号"/>
      </el-form-item>

      <el-form-item label="密码">
        <el-input v-model="form.passwd" placeholder="请输入密码" show-password/>
      </el-form-item>

      <div class="login-form-button">
        <el-form-item>
          <el-button type="primary" @click="login()">登录</el-button>
        </el-form-item>
      </div>

    </el-form>
  </div>

</template>

<script lang="ts">
import {LoginApi} from "@/api/login/login";
import {defineComponent, ref} from "vue";

export default defineComponent({
  setup() {
    const loginApi = new LoginApi()

    let form = ref({id: 0, passwd: ''})

    function login() {
      console.log("aaa")
      loginApi.login(form.value.id, form.value.passwd)
          .then((res) => {
            if (res.code === 200) {

            }
          })
    }

    return {
      form,
      login
    }
  },
})
</script>

<style>
.login-form {
}

.login-form-button {
  text-align: right;
}
</style>
