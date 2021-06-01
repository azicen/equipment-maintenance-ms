<template>
  <el-dialog title="用户维护信息" width="80%" :before-close="handleClose">
    <GetMaintainListPage :user-id="userId" ref="list" />
  </el-dialog>
</template>

<script>
import GetMaintainListPage from "@/components/root/page/GetMaintainListPage.vue";
import { defineComponent, ref } from "vue";
export default defineComponent({
  components: {
    GetMaintainListPage,
  },
  props: ["userId"],
  setup() {
    const list = ref();
    const refresh = () => {
      list.value.getList();
    };
    return { list, refresh };
  },
  data() {
    return {
      dialogVisible: false,
    };
  },
  methods: {
    handleClose(done) {
      this.$confirm("确认关闭？")
        .then(() => {
          done();
        })
        .catch(() => {});
    },
  },
});
</script>
