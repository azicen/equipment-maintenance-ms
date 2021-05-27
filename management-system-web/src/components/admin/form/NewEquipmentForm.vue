<template>
  <el-form
    ref="form"
    :model="form"
    :label-position="labelPosition"
    label-width="120px"
  >
    <el-form-item label="设备名称">
      <el-input v-model="form.name"></el-input>
    </el-form-item>
    <el-form-item label="地点">
      <el-input v-model="form.location"></el-input>
    </el-form-item>
    <el-form-item label="状况">
      <el-input v-model="form.status"></el-input>
    </el-form-item>
    <el-form-item label="设备类型">
      <el-select
        v-model="form.type_id"
        clearable
        filterable
        placeholder="请选择设备类型"
      >
        <el-option
          v-for="item in allType"
          :key="item.id"
          :label="item.name"
          :value="item.id"
        >
        </el-option>
      </el-select>
    </el-form-item>
    <el-form-item label="开始服役时间">
      <el-col :span="11">
        <el-date-picker
          type="date"
          placeholder="选择日期"
          v-model="form.start_date"
          style="width: 100%"
        ></el-date-picker>
      </el-col>
    </el-form-item>
    <el-form-item label="结束服役时间">
      <el-col :span="11">
        <el-date-picker
          type="date"
          placeholder="选择日期"
          v-model="form.deadline"
          style="width: 100%"
        ></el-date-picker>
      </el-col>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="onSubmit">立即创建</el-button>
      <el-button>取消</el-button>
    </el-form-item>
  </el-form>
</template>

<script>
import * as EquipmentAPI from "@/api/equipment";
import * as EquipmentTypeAPI from "@/api/equipment_type";

export default {
  data() {
    return {
      form: {
        name: "",
        location: "",
        status: "",
        start_date: "",
        deadline: "",
        type_id: "",
        user_id: 1,
      },
      allType: [],
    };
  },
  methods: {
    getEquipmentTypeList() {
      EquipmentTypeAPI.getEquipmentTypes().then((res) => {
        this.allType = res.data.equipment_types;
      });
    },
    addEquipment() {
      this.form.start_date = this.form.start_date.getTime() / 1000
      this.form.deadline = this.form.deadline.getTime() / 1000

      EquipmentAPI.postEquipment(this.form).then((res) => {
        this.$notify({
          title: "创建成功",
          message: `您索创建的设备ID为${res.data.id}`,
          type: "success",
        });
      })
      //new Date().getTime();
    },
    onSubmit() {
      this.addEquipment();
    },
  },
  beforeMount() {
    this.getEquipmentTypeList();
  },
};
</script>

<style>
</style>
