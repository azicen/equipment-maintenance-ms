<template>
  <h5>test</h5>
  <img src="@/assets/logo.png" alt="图片"/>
</template>

<script lang="ts">
import {defineComponent} from 'vue'
import {onMounted} from 'vue'
import {UserApi} from "@/api/user/user"

interface User {
  name: string,
  status: number,
  group: number,
}

export default defineComponent({
  // 已启用类型推断
  name: 'Test',
  setup() {
    const userApi = new UserApi()

    let user: User

    onMounted(() => {
      userApi.getUser(1).then((res) => {
        console.log(res)
        user = res.data.data
        console.log(user)
      })
    })
  }
})
</script>

<style>

</style>