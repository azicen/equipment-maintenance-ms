<template>
  <div class="login">
    <WordRotation/>
    <LoginForm/>
  </div>
</template>

<script>
import WordRotation from "@/views/component/data-display/WordRotation";
import LoginForm from "@/views/component/form/LoginForm";

export default {
  name: 'Login',
  components: {
    WordRotation,
    LoginForm,
  }
}
</script>
