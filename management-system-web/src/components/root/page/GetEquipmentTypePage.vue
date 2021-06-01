<template>
  <el-table
    :data="
      list.filter(
        (data) =>
          !search || data.name.toLowerCase().includes(search.toLowerCase())
      )
    "
    style="width: 100%"
    stripe
  >
    <el-table-column label="设备类型 ID" prop="id"> </el-table-column>
    <el-table-column label="设备类型称名" prop="name"> </el-table-column>
    <el-table-column align="right">
      <template #header>
        <el-row>
          <el-col :span="19">
            <el-input
              v-model="search"
              size="mini"
              placeholder="输入ID或姓名搜索"
            />
          </el-col>
          <el-col :span="5"
            ><el-button
              size="mini"
              @click="handleSearch(scope.$index, scope.row)"
              >搜索</el-button
            ></el-col
          >
        </el-row>
      </template>
      <template #default="scope">
        <el-button size="mini" @click="handleEdit(scope.$index, scope.row)"
          >编辑</el-button
        >
        <el-popover placement="top" :width="160" v-model="visible">
          <p>确定删除吗？</p>
          <div style="text-align: right; margin: 0">
            <el-button
              plain
              type="danger"
              size="mini"
              @click="handleDelete(scope.$index, scope.row)"
              >我确定要删除它</el-button
            >
          </div>
          <template #reference>
            <el-button size="mini" type="danger" @click="visible = true"
              >删除</el-button
            >
          </template>
        </el-popover>
      </template>
    </el-table-column>
  </el-table>
</template>

<script>
import * as EquipmentTypeAPI from "@/api/equipment_type";

export default {
  data() {
    return {
      list: [],
      search: "",
      defaultHeight: {
        height: "",
      },
      visible: true,
    };
  },
  methods: {
    getList() {
      EquipmentTypeAPI.getEquipmentTypes().then((res) => {
        this.list = this.list.concat(res.data.equipment_types);
      });
    },
    handleSearch(index, row) {
      console.log(index, row);
    },
    handleEdit(index, row) {
      console.log(index, row);
    },
    handleDelete(index, row) {
      this.visible = false;
      console.log(index, row);
      this.visible = true;
      EquipmentTypeAPI.delEquipmentType(row.id).then((res) => {
        this.$notify({
          title: "删除成功",
          message: `您所删除的设备类型ID为${res.data.id}`,
          type: "success",
        });
      });
    },
    //定义方法，获取高度减去头尾
    getHeight() {
      this.defaultHeight.height = window.innerHeight - 65 - 55 + "px";
    },
  },
  beforeMount() {
    this.getList();
  },
  created() {
    //页面创建时执行一次getHeight进行赋值，顺道绑定resize事件
    window.addEventListener("resize", this.getHeight);
    this.getHeight();
  },
};
</script>

<style>
</style>
