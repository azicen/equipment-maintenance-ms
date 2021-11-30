<template>
  <el-table :data="users">

    <el-table-column fixed prop="id" label="员工ID"/>

    <el-table-column prop="name" label="姓名"/>

    <el-table-column prop="tag" label="状态">
      <template #default="scope">
        <el-tag
            :type="scope.row.status === 0 ? 'danger' : ''"
            disable-transitions
        >
          <div v-if="scope.row.status === 0">已关闭</div>
          <div v-else>可使用</div>
        </el-tag>
      </template>
    </el-table-column>

    <el-table-column fixed="right" label="操作" width="120">
      <template #default="scope">

        <el-button
            type="text"
            size="small"
            @click.prevent=""
        >
          修改
        </el-button>

        <el-button
            type="text"
            size="small"
            @click.prevent="handleDelete(scope.$index, scope.row)"
        >
          删除
        </el-button>

      </template>
    </el-table-column>

  </el-table>
</template>

<script lang="ts">
import {defineComponent} from 'vue'
import {ElMessageBox, ElNotification} from 'element-plus'

export interface UserDate {
  id: number,
  name: string,
  status: number,
}

export default defineComponent({
  name: 'AdminUser',
  components: {},
  setup() {
    // 用户列表
    let users: UserDate[] = [
      {
        id: 0,
        name: "test",
        status: 0,
      },
    ]

    return {
      users,
    }
  },
  methods: {
    reviseClick() {

    },
    // 点击删除按钮事件函数
    handleDelete(index: number, data: UserDate) {
      console.log(index, data)
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
            ElNotification({
              title: 'Success',
              message: `成功删除 ${data.name} 的账号`,
              type: 'success',
            })
          })
    },
  },
})
</script>

<style scoped>

</style>