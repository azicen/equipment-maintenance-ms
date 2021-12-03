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
import { useRouter } from "vue-router";
import {ElNotification} from "element-plus";

export default defineComponent({
  setup() {
    const router = useRouter();

    const loginApi = new LoginApi()

    let form = ref({id: 0, passwd: ''})

    function login() {
      loginApi.login(form.value.id, form.value.passwd)
          .then((res) => {
            if (res.code === 200) {
              ElNotification({
                title: 'Success',
                message: `登录成功`,
                type: 'success',
              })

              if (form.value.id === 1) {
                console.log('admin')
                router.push({ path: "admin" })
              } else {
                console.log('user')
                router.push({ path: "/" })
              }

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
