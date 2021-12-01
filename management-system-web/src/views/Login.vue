<template>
  <div class="login-view">
    <WordRotation class="login-word-rotation"/>
    <el-card class="login-view-box-card">
      <LoginForm/>
    </el-card>
  </div>
</template>

<script lang="ts">
import WordRotation from "@/components/data-display/WordRotation.vue";
import LoginForm from "@/components/form/LoginForm.vue";
import {defineComponent} from "vue";

export default defineComponent({
  name: 'Login',
  components: {
    WordRotation,
    LoginForm,
  }
})
</script>

<style>
.login-word-rotation{
  /*margin: 0 20%;*/
  margin: 5% 0;
  width: 100%;
  height: 300px;
  display: flex;
  display: -webkit-flex;
  justify-content: center;
  align-items: center;
}

.login-view-box-card {
  margin: auto;
  width: 400px;
  height: 200px;
}
</style>