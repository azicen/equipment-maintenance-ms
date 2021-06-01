<template>
  <el-container>
    <el-main :style="defaultHeight">
      <el-table :data="tableData" style="width: 100%" stripe>
        <el-table-column type="expand">
          <template #default="props">
            <el-form label-position="left" inline class="demo-table-expand">
              <el-form-item label="异常信息">
                <span>{{ props.row.remark }}</span>
              </el-form-item>
              <el-form-item>
                <el-footer height="25px">
                  <el-popover placement="top" :width="160" v-model="visible">
                    <p>确定删除吗？</p>
                    <div style="text-align: right; margin: 0">
                      <el-button
                        plain
                        type="danger"
                        size="mini"
                        @click="handleDelete(props.$index, props.row)"
                        >我确定要删除它</el-button
                      >
                    </div>
                    <template #reference>
                      <el-button
                        size="mini"
                        type="danger"
                        @click="visible = true"
                        >删除</el-button
                      >
                    </template>
                  </el-popover>
                </el-footer>
              </el-form-item>
            </el-form>
          </template>
        </el-table-column>
        <el-table-column label="维护人 ID" prop="user_id"> </el-table-column>
        <el-table-column label="设备 ID" prop="equipment_id"> </el-table-column>
        <el-table-column label="维护时间" prop="date"> </el-table-column>
        <el-table-column label="状态" prop="status"> </el-table-column>
      </el-table>
    </el-main>

    <el-footer height="10px">
      <el-button size="small" @click="getList()">获取更多</el-button>
    </el-footer>
  </el-container>
</template>

<script>
import * as MaintainAPI from "@/api/maintain";
import { defineComponent } from "vue";

export default defineComponent({
  props: ["userId"],
  setup() {},
  data() {
    return {
      tableData: [],
      search: "",
      currentUserId: 0,
      i: 0,
      defaultHeight: {
        height: "",
      },
      visible: true,
    };
  },
  methods: {
    getList() {
      console.log("user_id: " + this.userId);
      if (this.userId == undefined) {
        if (this.currentUserId != 0) {
          this.i = 0;
          this.tableData = [];
          this.currentUserId = 0;
        }
      } else {
        if (this.currentUserId != this.userId) {
          this.i = 0;
          this.tableData = [];
          this.currentUserId = this.userId;
        }
      }
      this.i = this.i + 1;
      MaintainAPI.getMaintains(this.i, this.currentUserId).then((res) => {
        this.tableData = this.tableData.concat(res.data.maintains);
        console.log(res.data.maintains);
      });
      this.i = this.i + 9;
    },
    handleRetired(index, row) {
      console.log(index, row);
    },
    handleDelete(index, row) {
      this.visible = false;
      console.log(index, row);
      this.visible = true;
      MaintainAPI.delMaintain(row.id).then((res) => {
        this.$notify({
          title: "删除成功",
          message: `您所删除的维护信息ID为${res.data.id}`,
          type: "success",
        });
      });
    },
    //定义方法，获取高度减去头尾
    getHeight() {
      this.defaultHeight.height = window.innerHeight - 65 - 55 + "px";
    },
    dateFormat(fmt, date) {
      let ret;
      const opt = {
        "Y+": date.getFullYear().toString(), // 年
        "m+": (date.getMonth() + 1).toString(), // 月
        "d+": date.getDate().toString(), // 日
        "H+": date.getHours().toString(), // 时
        "M+": date.getMinutes().toString(), // 分
        "S+": date.getSeconds().toString(), // 秒
        // 有其他格式化字符需求可以继续添加，必须转化成字符串
      };
      for (let k in opt) {
        ret = new RegExp("(" + k + ")").exec(fmt);
        if (ret) {
          fmt = fmt.replace(
            ret[1],
            ret[1].length == 1 ? opt[k] : opt[k].padStart(ret[1].length, "0")
          );
        }
      }
      return fmt;
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
});
</script>

<style>
.demo-table-expand {
  font-size: 0;
}
.demo-table-expand label {
  width: 100px;
  color: #99a9bf;
}
.demo-table-expand .el-form-item {
  margin-right: 0;
  margin-bottom: 0;
  width: 50%;
}
</style>
