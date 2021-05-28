<template>
  <el-table :data="tableData" style="width: 100%">
    <el-table-column type="expand">
      <template #default="props">
        <el-form label-position="left" inline class="demo-table-expand">
          <el-form-item label="类型 ID">
            <span>{{ props.row.type_id }}</span>
          </el-form-item>
          <el-form-item label="位置">
            <span>{{ props.row.location }}</span>
          </el-form-item>
          <el-form-item label="服役时间">
            <span>{{ props.row.start_date }}</span>
          </el-form-item>
          <el-form-item label="预计退役时间">
            <span>{{ props.row.deadline }}</span>
          </el-form-item>
          <el-form-item>
            <el-footer height="25px">
              <el-button
                size="mini"
                @click="handleRetired(props.$index, props.row)"
                >提前退役</el-button
              >
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
                  <el-button size="mini" type="danger" @click="visible = true"
                    >删除</el-button
                  >
                </template>
              </el-popover>
            </el-footer>
          </el-form-item>
        </el-form>
      </template>
    </el-table-column>
    <el-table-column label="设备 ID" prop="id"> </el-table-column>
    <el-table-column label="设备名称" prop="name"> </el-table-column>
    <el-table-column label="状态" prop="status"> </el-table-column>
    <el-table-column label="近期维护时间" prop="date"> </el-table-column>
    <el-table-column label="近期维护人 ID" prop="user_id"> </el-table-column>
  </el-table>
</template>

<script>
import * as EquipmentAPI from "@/api/equipment";

export default {
  data() {
    return {
      tableData: [],
      search: "",
      i: 0,
      defaultHeight: {
        height: "",
      },
      visible: true,
    };
  },
  methods: {
    getList() {
      this.i = this.i + 1;
      EquipmentAPI.getEquipments(this.i).then((res) => {
        this.tableData = this.tableData.concat(res.data.equipments);
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
};
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
