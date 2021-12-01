<template>
  <div class="login-form">
    <el-form ref="form" :model="form">
      <el-form-item label="账户">
        <el-input v-model="form.id" placeholder="account"/>
      </el-form-item>
      <el-form-item label="密码">
        <el-input v-model="form.password" placeholder="password" show-password/>
      </el-form-item>
      <div class="login-form-button">
        <el-form-item>
          <el-button type="primary" @click="login">登录</el-button>
        </el-form-item>
      </div>
    </el-form>
  </div>
</template>

<script>
import {UserLogin} from "@/api/login/userLogin";

export default {
  data() {
    return {
      form: {
        id: '',
        password: '',
      },
    }
  },
  setup(){
    const userApi = new UserLogin()
    function login(){
      console.log("aaa")
      userApi.Login(this.data().form.id,this.data().form.password).then(
          (res)=>{
            if(res.code===200){
              alert("aaa")
            }
          console.log(res.data)
        }
      )
    }
    return {
      login
    }
  },
}
</script>

<style>
.login-form {
}

.login-form-button {
  text-align: right;
}
</style>