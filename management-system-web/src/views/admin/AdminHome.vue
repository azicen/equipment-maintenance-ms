<template>
  <div>
    <Summary :data="data"></Summary>
  </div>
</template>

<script lang="ts">
import {defineComponent} from 'vue'
import Summary, {SummaryData} from "@/components/data-display/Summary.vue"

export default defineComponent({
  name: 'AdminHome',
  props: {},
  components: {
    Summary,
  },
  setup() {
    const data: SummaryData = {
      today: 1,
      wait: 2,
      alarm: 3,
      scrap: 4,
    }
    return {
      data
    }
  }
})
</script>

<style>

</style>