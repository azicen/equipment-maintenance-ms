<template>

  <div>
    <el-dialog v-model="dialogVisible" :title="'修改用户 '+ dialogForm.id +' 的信息'" width="30%">
      <el-form :model="dialogForm">
        <el-form-item label="用户名">
          <el-input v-model="dialogForm.name" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="账号状态">
          <el-switch
              v-model="dialogForm.status"
              inline-prompt
              active-color="#13ce66"
              inactive-color="#ff4949"
              active-text="Y"
              inactive-text="N"
          />
        </el-form-item>
      </el-form>

      <template #footer>
      <span class="dialog-footer">
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click.prevent="handleUpdate()">确定</el-button>
      </span>
      </template>
    </el-dialog>
  </div>

  <div>
    <el-button type="primary" @click.prevent="">添加用户</el-button>
  </div>

  <div>
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
          <el-button type="text" size="small" @click.prevent="reviseClick(scope.$index)">修改</el-button>
          <el-button type="text" size="small" @click.prevent="handleDelete(scope.$index, scope.row)">
            删除
          </el-button>
        </template>
      </el-table-column>
    </el-table>
  </div>

</template>

<script lang="ts">
import {defineComponent, onMounted, Ref, ref} from 'vue'
import {ElMessageBox, ElNotification} from 'element-plus'

import {GetUserResponse, UserApi} from "@/api/user/user"
import {ResponseData} from "@/api/response-data"

export default defineComponent({
  name: 'AdminUser',
  components: {},
  setup() {
    const dialogVisible = ref(false)

    const userApi = new UserApi()

    let page: number = 0

    // 用户列表
    let users: Ref<GetUserResponse[]> = ref<GetUserResponse[]>([])

    let dialogForm = ref<GetUserResponse>({id: 0, name: "xxx", status: true})

    onMounted(() => {
      userApi.getUsers(10, page).then((res) => {
        users.value = res.data
      })
    })

    function reviseClick(i: number) {
      dialogForm.value = users.value[i]
      console.log('reviseClick', dialogForm.value)
      dialogVisible.value = true
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

    // 更新用户数据
    function handleUpdate() {
      userApi.updateUser(dialogForm.value.id, dialogForm.value.name, dialogForm.value.status)
          .then(() => {
            // 修改成功提示
            ElNotification({
              title: 'Success',
              message: `成功修改 ${dialogForm.value.name} 的账号`,
              type: 'success',
            })
          })
          .catch((res) => {
            // 修改失败提示
            ElNotification({
              title: 'Error',
              message: '修改失败！' + res.msg,
              type: 'error',
            })
          })
      dialogVisible.value = false
    }

    return {
      users,
      page,
      dialogForm,
      dialogVisible,

      reviseClick,
      handleDelete,
      handleUpdate,
    }
  },
})
</script>

<style scoped>

</style>