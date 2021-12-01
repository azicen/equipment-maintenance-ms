<template>

  <el-table :data="users">

    <el-table-column fixed prop="id" label="员工ID"/>

    <el-table-column prop="name" label="姓名"/>

    <el-table-column prop="tag" label="状态">
      <template #default="scope">
        <el-tag
            :type="scope.row.status ? '' : 'danger'"
            disable-transitions
        >
          <div v-if="scope.row.status">可使用</div>
          <div v-else>已关闭</div>
        </el-tag>
      </template>
    </el-table-column>

    <el-table-column fixed="right" label="操作" width="120">
      <template #default="scope">

        <UpdateUserDialog v-model="dialogVisible" :data="users[index]"/>

        <el-button type="text" size="small" @click="dialogVisible = true">修改</el-button>

        <el-button type="text" size="small" @click.prevent="handleDelete(scope.$index, scope.row)">
          删除
        </el-button>

      </template>
    </el-table-column>

  </el-table>

</template>

<script lang="ts">
import {defineComponent, onMounted, Ref, ref} from 'vue'
import {ElMessageBox, ElNotification} from 'element-plus'

import {GetUserResponse, UserApi} from "@/api/user/user"
import {ResponseData} from "@/api/response-data"

import UpdateUserDialog from "@/components/dialog/UpdateUserDialog.vue"

export default defineComponent({
  name: 'AdminUser',
  components: {
    UpdateUserDialog,
  },
  setup() {
    const dialogVisible = ref(false)
    const index = ref(0)

    const userApi = new UserApi()

    let page: number = 0

    // 用户列表
    let users: Ref<GetUserResponse[]> = ref<GetUserResponse[]>([])

    onMounted(() => {
      userApi.getUsers(10, page).then((res) => {
        users.value = res.data
      })
    })

    function reviseClick() {

    }

    // 点击删除按钮事件函数
    function handleDelete(i: number, data: GetUserResponse) {
      console.log(data)
      ElMessageBox.confirm(
          `你确定要删除 ${data.name} 的账号吗？`,
          '警告',
          {
            confirmButtonText: '确认',
            cancelButtonText: '取消',
            type: 'warning',
            center: true,
          }
      )
          .then(() => {
            userApi.deleteUser(data.id).then(() => {
              // 从列表中删除数据
              delete users.value[i]
              // 删除成功提示
              ElNotification({
                title: 'Success',
                message: `成功删除 ${data.name} 的账号`,
                type: 'success',
              })
            }).catch((res: ResponseData<any>) => {
              // 删除失败提示
              ElNotification({
                title: 'Error',
                message: '删除失败！' + res.msg,
                type: 'error',
              })
            })
          })
          .catch(() => {
          })
    }

    return {
      users,
      page,
      dialogVisible,
      index,
      reviseClick,
      handleDelete,
    }
  },
})
</script>

<style scoped>

</style>