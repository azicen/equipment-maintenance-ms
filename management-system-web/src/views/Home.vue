<template>
  <div class="home">

    <el-card class="home-card">
      <el-form
          ref="form"
          :model="form"
          label-width="120px"
      >
        <el-form-item label="维护设备 ID">
          <el-input
              onkeyup="this.value = this.value.replace(/[^\d.]/g,'');"
              v-model="form.equipment_id"
          ></el-input>
        </el-form-item>
        <el-form-item label="维护时间">
          <el-col :span="11">
            <el-date-picker
                type="datetime"
                placeholder="选择时间"
                v-model="form.date"
                style="width: 100%"
                :default-time="new Date()"
            ></el-date-picker>
          </el-col>
        </el-form-item>
        <el-form-item label="设备状态">
          <el-input v-model="form.status"></el-input>
        </el-form-item>
        <el-form-item label="设备异常">
          <el-input
              type="textarea"
              :autosize="{ minRows: 5, maxRows: 20 }"
              v-model="form.remark"
          ></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="onSubmit()">完成维护</el-button>
          <el-button>取消</el-button>
        </el-form-item>
      </el-form>
    </el-card>

  </div>
</template>

<script lang="ts">
import {defineComponent, ref} from "vue";
import {MaintainApi} from "@/api/maintain/maintain";
import {ElNotification} from "element-plus";

export default defineComponent({
  name: 'Home',
  components: {},
  setup() {
    const maintainApi = new MaintainApi()

    // 设备列表
    let form = ref({
      equipment_id: 0,
      date: 0,
      status: "",
      remark: "",
    })

    function onSubmit() {
      maintainApi.addMaintain(form.value.equipment_id, form.value.date, form.value.status, form.value.remark)
          .then(() => {
            ElNotification({
              title: 'Success',
              message: `添加成功`,
              type: 'success',
            })
          })
          .catch(() => {
            ElNotification({
              title: 'Error',
              message: '添加失败',
              type: 'error',
            })
          })
    }

    return {
      form,

      onSubmit
    }
  }
})
</script>

<style>
.home-card {
  margin: auto;
  width: 80%;
}
</style>