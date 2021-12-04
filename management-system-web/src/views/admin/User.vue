<template>

  <!--  添加用户弹窗-->
  <div>
    <el-dialog v-model="addVisible" title="添加用户" width="30%">
      <el-form :model="addForm">
        <el-form-item label="用户名">
          <el-input v-model="addForm.name" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="账号状态">
          <el-switch
              v-model="addForm.status"
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
        <el-button @click="addVisible = false">取消</el-button>
        <el-button type="primary" @click.prevent="handleAdd()">确定</el-button>
      </span>
      </template>
    </el-dialog>
  </div>

  <!--  修改用户信息弹窗-->
  <div>
    <el-dialog v-model="updateVisible" :title="'修改用户 '+ updateForm.id +' 的信息'" width="30%">
      <el-form :model="updateForm">
        <el-form-item label="用户名">
          <el-input v-model="updateForm.name" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="账号状态">
          <el-switch
              v-model="updateForm.status"
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
        <el-button @click="updateVisible = false">取消</el-button>
        <el-button type="primary" @click.prevent="handleUpdate()">确定</el-button>
      </span>
      </template>
    </el-dialog>
  </div>

  <!--  添加用户按钮-->
  <div>
    <el-button type="primary" @click="addVisible = true">添加用户</el-button>
  </div>

  <!--  用户信息列表-->
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

  <div>
    <el-pagination
        background
        layout="prev, pager, next"
        v-model:current-page="page"
        v-model:page-size="pageSize"
        v-model:total="total"
        @current-change="updatePage()"
        @prev-click="page-1"
        @next-click="page+1"
    />
  </div>

</template>

<script lang="ts">
import {defineComponent, onBeforeMount, onMounted, Ref, ref} from 'vue'
import {ElMessageBox, ElNotification} from 'element-plus'

import {GetUserResponse, UserApi} from "@/api/user/user"
import {ResponseData} from "@/api/response-data"

export default defineComponent({
  name: 'AdminUser',
  components: {},
  setup() {
    const addVisible = ref(false)
    const updateVisible = ref(false)

    const userApi = new UserApi()

    // 单页数量
    const pageSize = ref(10)
    // 总数
    let total = ref(0)
    // 页码
    let page = ref(1)

    // 用户列表
    let users: Ref<GetUserResponse[]> = ref<GetUserResponse[]>([])

    let addForm = ref({name: "", status: true})
    let updateForm = ref<GetUserResponse>({id: 0, name: "xxx", status: true})

    // 组件挂载之前
    onBeforeMount(() => {
      userApi.getUserListSize()
          .then((res) => {
            total.value = res.data.size
          })
    })

    // 组件挂载完成
    onMounted(() => {
      updatePage()
    })

    // 更新列表信息
    function updatePage() {
      userApi.getUsers(pageSize.value, page.value - 1)
          .then((res) => {
            users.value = res.data
          })
    }

    function reviseClick(i: number) {
      updateForm.value = users.value[i]
      updateVisible.value = true
    }

    // 添加用户按钮
    function handleAdd() {
      userApi.addUser(addForm.value.name, addForm.value.status)
          .then((res) => {
            if (res.code === 200) {
              ElNotification({
                title: 'Success',
                message: `成功添加 ${addForm.value.name} 的账号，ID: ${res.data.id}`,
                type: 'success',
              })
            } else {
              // 删除失败提示
              ElNotification({
                title: 'Error',
                message: '添加失败！' + res.msg,
                type: 'error',
              })
            }
            addForm.value = {name: "", status: true}
          })
          .catch(() => {
            // 请求失败
            ElNotification({
              title: 'Error',
              message: '添加失败！',
              type: 'error',
            })
            addForm.value = {name: "", status: true}
          })

      addVisible.value = false
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
            userApi.deleteUser(data.id)
                .then(() => {
                  // 从列表中删除数据
                  delete users.value[i]
                  // 删除成功提示
                  ElNotification({
                    title: 'Success',
                    message: `成功删除 ${data.name} 的账号`,
                    type: 'success',
                  })
                })
                .catch((res: ResponseData<any>) => {
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
      userApi.updateUser(updateForm.value.id, updateForm.value.name, updateForm.value.status)
          .then(() => {
            // 修改成功提示
            ElNotification({
              title: 'Success',
              message: `成功修改 ${updateForm.value.name} 的账号`,
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
      updateVisible.value = false
    }

    return {
      users,
      pageSize,
      total,
      page,
      addForm,
      updateForm,
      addVisible,
      updateVisible,

      updatePage,
      handleAdd,
      reviseClick,
      handleDelete,
      handleUpdate,
    }
  },
})
</script>

<style scoped>

</style>