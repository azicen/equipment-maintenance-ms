<template>
  <div class="login-page">
    <WordRotation/>
    <el-card class="login-form-box-card">
      <LoginForm/>
    </el-card>
  </div>
</template>

<script>
import WordRotation from "@/views/component/data-display/WordRotation.vue";
import LoginForm from "@/views/component/form/LoginForm.vue";

export default {
  name: 'Login',
  components: {
    WordRotation,
    LoginForm,
  }
}
</script>

<style>
.login-form-box-card {
  margin: 0 25%;
  width: 480px;
  height: 200px;
}
</style>