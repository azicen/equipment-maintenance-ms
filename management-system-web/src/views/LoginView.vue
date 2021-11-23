<template>
  <div class="login-view">
    <WordRotation/>
    <el-card class="login-view-box-card">
      <LoginForm/>
    </el-card>
  </div>
</template>

<script lang="ts">
import WordRotation from "@/components/data-display/WordRotation.vue";
import LoginForm from "@/components/form/LoginForm.vue";

export default {
  name: 'Login',
  components: {
    WordRotation,
    LoginForm,
  }
}
</script>

<style>
.login-view-box-card {
  margin: 0 25%;
  width: 480px;
  height: 200px;
}
</style>