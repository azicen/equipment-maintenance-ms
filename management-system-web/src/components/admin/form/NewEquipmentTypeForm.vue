<template>
  <el-form
    ref="form"
    :model="form"
    :label-position="labelPosition"
    label-width="120px"
  >
    <el-form-item label="类别名称">
      <el-input v-model="form.name"></el-input>
    </el-form-item>
    <el-form-item label="维护周期">
      <el-input-number v-model="form.cycle" :min="1">
        <template #append>天</template>
      </el-input-number>
    </el-form-item>
    <el-form-item label="权限组">
      <el-select
        v-model="groups"
        multiple
        filterable
        allow-create
        default-first-option
        placeholder="请选择权限组"
      >
        <el-option
          v-for="item in allGroups"
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
        cycle: "",
      },
      groups: [],
      allGroups: [],
    };
  },
  methods: {
    getGroupList() {
      GroupAPI.getGroups().then((res) => {
        this.allGroups = res.data.groups;
      });
    },
    addEquipmentType() {
      EquipmentTypeAPI.postEquipmentType(this.form).then((res) => {
        this.$notify({
          title: "创建成功",
          message: `您索创建的设备类型ID为${res.data.id}`,
          type: "success",
        });
      });
    },
    onSubmit() {
      this.addEquipmentType();
      //console.log("submit!");
    },
  },
  beforeMount() {
    this.getGroupList();
  },
};
</script>

<style>
</style>
