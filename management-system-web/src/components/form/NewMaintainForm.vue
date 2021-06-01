<template>
  <el-form
    ref="form"
    :model="form"
    :label-position="labelPosition"
    label-width="120px"
  >
  <el-form-item label="维护人 ID">
      <el-input
        onkeyup="this.value = this.value.replace(/[^\d.]/g,'');"
        v-model="form.user_id"
      ></el-input>
    </el-form-item>
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
          :default-time="currentTime"
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
      <el-button type="primary" @click="onSubmit">完成维护</el-button>
      <el-button>取消</el-button>
    </el-form-item>
  </el-form>
</template>

<script>
import * as MaintainAPI from "@/api/maintain";

export default {
  data() {
    return {
      currentTime: new Date(),
      form: {
        user_id: 1,
        equipment_id: "",
        date: "",
        status: "",
        remark: "",
      },
    };
  },
  methods: {
    addMaintain() {
      this.form.equipment_id = parseInt(this.form.equipment_id);
      this.form.date = this.form.date.getTime() / 1000;
      MaintainAPI.postMaintain(this.form).then((res) => {
        this.$notify({
          title: "创建成功",
          message: `您所创建的维护表单ID为${res.data.id}`,
          type: "success",
        });
      });
    },
    onSubmit() {
      this.addMaintain();
    },
  },
  beforeMount() {},
};
</script>

<style>
</style>
