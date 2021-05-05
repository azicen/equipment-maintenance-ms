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
      <el-input v-model="form.cycle">
        <template #append>天</template>
      </el-input>
    </el-form-item>
    <el-form-item label="权限组">
      <el-select
        v-model="form.groups"
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
export default {
  data() {
    return {
      form: {
        typeId: 0,
        name: "",
        cycle: "",
        groups: [],
      },
      allGroups: [
        {
          id: 1,
          name: "权限组1",
        },
        {
          id: 2,
          name: "权限组2",
        },
        {
          id: 3,
          name: "权限组3",
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
