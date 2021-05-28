<template>
  <el-container>
    <el-main :style="defaultHeight">
      <el-table
        :data="
          userList.filter(
            (data) =>
              !search || data.name.toLowerCase().includes(search.toLowerCase())
          )
        "
        style="width: 100%"
        stripe
      >
        <el-table-column type="selection" width="55"> </el-table-column>
        <el-table-column label="用户ID" prop="id"> </el-table-column>
        <el-table-column label="用户名" prop="name"> </el-table-column>
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
    </el-main>
    <el-footer height="10px">
      <el-button size="small" @click="getUserList()">获取更多</el-button>
      <el-button size="small" type="danger" @click="handleDeleteAll()"
        >全部删除</el-button
      >
    </el-footer>
  </el-container>
</template>

<script>
import * as UserAPI from "@/api/user";

export default {
  data() {
    return {
      userList: [],
      search: "",
      i: 0,
      defaultHeight: {
        height: "",
      },
      visible: true,
    };
  },
  methods: {
    getUserList() {
      this.i = this.i + 1;
      UserAPI.getUserList(this.i).then((res) => {
        this.userList = this.userList.concat(res.data.users);
      });
      this.i = this.i + 9;
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
    },
    //定义方法，获取高度减去头尾
    getHeight() {
      this.defaultHeight.height = window.innerHeight - 65 - 55 + "px";
    },
  },
  beforeMount() {
    this.getUserList();
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
