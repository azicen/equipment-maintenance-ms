<template>
  <web-nav-menu/>
  <router-view/>
</template>

<script>
// @ is an alias to /src
import WebNavMenu from "@/components/WebNavMenu.vue";

export default {
  name: "app",
  components: {
    WebNavMenu,
  },
};
</script>

<style>
</style>
