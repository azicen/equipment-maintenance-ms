<template>

  <!--  添加用户弹窗-->
  <div>
    <el-dialog v-model="addVisible" title="添加设备" width="30%">
      <el-form :model="addForm">
        <el-form-item label="设备名称">
          <el-input v-model="addForm.name" autocomplete="off"></el-input>
        </el-form-item>

        <el-form-item label="地点">
          <el-input v-model="addForm.location" autocomplete="off"></el-input>
        </el-form-item>

        <el-form-item label="状况">
          <el-input v-model="addForm.status" autocomplete="off"></el-input>
        </el-form-item>

        <el-form-item label="结束服务日期">
          <el-date-picker
              v-model="addForm.deadline"
              type="date"
              placeholder="一年后"
              :default-value="new Date(new Date().getFullYear()+1)"
          >
          </el-date-picker>
        </el-form-item>

        <el-form-item label="类型">
          <el-input v-model="addForm.type_id" autocomplete="off"></el-input>
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

  <!--  添加设备按钮-->
  <div>
    <el-button type="primary" @click="addVisible = true">添加设备</el-button>
  </div>

  <!--  设备信息列表-->
  <div>
    <el-table :data="equipments">
      <el-table-column fixed prop="id" label="设备ID"/>
      <el-table-column prop="name" label="名称"/>
      <el-table-column prop="location" label="地点"/>
      <el-table-column prop="status" label="状况"/>
      <el-table-column prop="start_date" label="开始服役日期"/>
      <el-table-column prop="deadline" label="结束服务日期"/>
      <el-table-column prop="type_id" label="类型"/>
      <el-table-column prop="user_id" label="创建人员编号"/>

      <el-table-column fixed="right" label="操作" width="120">
        <template #default="scope">
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
import {defineComponent, onBeforeMount, onMounted, ref} from 'vue'
import {EquipmentApi, GetEquipmentResponse} from "@/api/equipment/equipment";
import {ElMessageBox, ElNotification} from "element-plus";
import {GetUserResponse} from "@/api/user/user";
import {ResponseData} from "@/api/response-data";

export default defineComponent({
  name: 'AdminEquipment',
  components: {},
  setup() {
    const addVisible = ref(false)

    const equipmentApi = new EquipmentApi()

    // 单页数量
    const pageSize = ref(10)
    // 总数
    let total = ref(0)
    // 页码
    let page = ref(1)

    // 设备列表
    let equipments = ref<GetEquipmentResponse[]>([])

    let addForm = ref({
      name: "",
      location: "",
      status: "",
      deadline: new Date(new Date().getFullYear() + 1),
      type_id: 0,
    })

    function resetAddForm() {
      addForm.value = {
        name: "",
        location: "",
        status: "",
        deadline: new Date(new Date().getFullYear() + 1),
        type_id: 0,
      }
    }

    // 组件挂载之前
    onBeforeMount(() => {
      equipmentApi.getEquipmentListSize()
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
      equipmentApi.getEquipments(pageSize.value, page.value - 1)
          .then((res) => {
            equipments.value = res.data
          })
    }

    // 添加设备
    function handleAdd() {
      equipmentApi.addEquipment(
          addForm.value.name,
          addForm.value.location,
          addForm.value.status,
          addForm.value.deadline.getTime(),
          addForm.value.type_id
      )
          .then((res) => {
            if (res.code === 200) {
              ElNotification({
                title: 'Success',
                message: `成功添加 ${addForm.value.name} 设备，ID: ${res.data.id}`,
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
            resetAddForm()
          })
          .catch(() => {
            // 请求失败
            ElNotification({
              title: 'Error',
              message: '添加失败！',
              type: 'error',
            })
            resetAddForm()
          })

      addVisible.value = false
    }

    // 点击删除按钮事件函数
    function handleDelete(i: number, data: GetUserResponse) {
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
            equipmentApi.deleteUser(data.id)
                .then(() => {
                  // 从列表中删除数据
                  delete equipments.value[i]
                  // 删除成功提示
                  ElNotification({
                    title: 'Success',
                    message: `成功删除 ${data.name} 设备`,
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

    return {
      pageSize,
      total,
      page,
      equipments,
      addForm,
      addVisible,

      updatePage,
      handleAdd,
      handleDelete,
    }
  }
})
</script>

<style scoped>

</style>