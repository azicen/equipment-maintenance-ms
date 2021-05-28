<template>
  <el-form
    ref="form"
    :model="form"
    :label-position="labelPosition"
    label-width="120px"
  >
    <el-form-item label="权限组名称">
      <el-input v-model="form.name"></el-input>
    </el-form-item>
    <el-form-item label="拥有设备类别">
      <el-select
        v-model="types"
        multiple
        filterable
        allow-create
        default-first-option
        placeholder="请选设备类别"
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
    <el-form-item>
      <el-button type="primary" @click="onSubmit">立即创建</el-button>
      <el-button>取消</el-button>
    </el-form-item>
  </el-form>
</template>

<script>
import * as GroupAPI from "@/api/group";
import * as EquipmentTypeAPI from "@/api/equipment_type";

export default {
  data() {
    return {
      form: {
        name: "",
      },
      types: [],
      allType: [],
    };
  },
  methods: {
    getEquipmentTypeList() {
      EquipmentTypeAPI.getEquipmentTypes().then((res) => {
        this.allType = res.data.equipment_types;
      });
    },
    addGroup() {
      GroupAPI.postGroup(this.form).then((res) => {
        this.$notify({
          title: "创建成功",
          message: `您所创建的权限组ID为${res.data.id}`,
          type: "success",
        });
      });
    },
    onSubmit() {
      this.addGroup()
    },
  },
  beforeMount() {
    this.getEquipmentTypeList();
  },
};
</script>

<style>
</style>
