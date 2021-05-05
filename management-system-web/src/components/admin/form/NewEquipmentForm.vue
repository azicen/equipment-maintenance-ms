<template>
  <el-form
    ref="form"
    :model="form"
    :label-position="labelPosition"
    label-width="120px"
  >
    <el-form-item label="设备类型">
      <el-select
        v-model="form.typeId"
        clearable
        filterable
        placeholder="请选择设备类型"
      >
        <el-option
          v-for="item in types"
          :key="item.id"
          :label="item.name"
          :value="item.id"
        >
        </el-option>
      </el-select>
    </el-form-item>
    <el-form-item label="设备名称">
      <el-input v-model="form.name"></el-input>
    </el-form-item>
    <el-form-item label="地点">
      <el-input v-model="form.location"></el-input>
    </el-form-item>
    <el-form-item label="状况">
      <el-input v-model="form.status"></el-input>
    </el-form-item>
    <el-form-item label="开始服役时间">
      <el-col :span="11">
        <el-date-picker
          type="date"
          placeholder="选择日期"
          v-model="form.date"
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
export default {
  data() {
    return {
      form: {
        id: 0,
        typeId: "",
        name: "",
        location: "",
        status: "",
        date: "",
        deadline: "",
      },
      types: [
        {
          id: 1,
          name: "设备类型1",
        },
        {
          id: 2,
          name: "设备类型2",
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
