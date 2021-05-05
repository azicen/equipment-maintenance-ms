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
        v-model="form.types"
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
export default {
  data() {
    return {
      form: {
        groupId: 0,
        name: "",
        types: [],
      },
      allType: [
        {
          id: 1,
          name: "设备类别1",
        },
        {
          id: 2,
          name: "设备类别2",
        },
      ],
    };
  },
  methods: {
    onSubmit() {
      console.log("submit!");
    },
  },
};
</script>

<style>
</style>
