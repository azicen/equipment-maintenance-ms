<template>
  <div class="home">
    <WordRotation/>
    <img alt="Vue logo" src="../assets/logo.png">
    <HelloWorld msg="Welcome to Your Vue.js App"/>
  </div>
</template>

<script>
import HelloWorld from '@/components/HelloWorld.vue'
import WordRotation from "@/views/component/data-display/WordRotation";

export default {
  name: 'Home',
  components: {
    HelloWorld,
    WordRotation
  }
}
</script>
