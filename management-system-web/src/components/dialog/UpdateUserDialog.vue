<template>

  <el-dialog v-model="dialogFormVisible" :title="'修改用户 '+ form.id +' 的信息'" width="30%">

    <el-form :model="form">
      <el-form-item label="用户名">
        <el-input v-model="form.name" autocomplete="off"></el-input>
      </el-form-item>
      <el-form-item label="账号状态">
        <el-switch
            v-model="form.status"
            inline-prompt
            active-color="#13ce66"
            inactive-color="#ff4949"
            active-text="Y"
            inactive-text="N"
        />
      </el-form-item>
    </el-form>

    <template #footer>
      <span class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取消</el-button>
        <el-button type="primary" @click.prevent="handleUpdate()">
          确定
        </el-button>
      </span>
    </template>

  </el-dialog>

</template>

<script lang="ts">
import {computed, defineComponent, onMounted, PropType, ref} from "vue"
import {GetUserResponse, UserApi} from "@/api/user/user"

export default defineComponent({
  name: 'UpdateUserDialog',

  components: {},

  props: {
    modelValue: Boolean,
    data: {
      type: Object as PropType<GetUserResponse>,
      required: true,
    },
  },

  // 双向绑定的回调事件
  emits: ['update:modelValue'],

  setup(props, context) {
    const userApi = new UserApi()

    // 嵌套双向绑定
    const dialogFormVisible = computed({
      get: () => props.modelValue,
      set: (v) => {
        context.emit('update:modelValue', v)
      }
    })

    // 表单数据
    const form = ref<GetUserResponse>({id: 0, name: '', status: true})

    function handleUpdate() {
      userApi.updateUser(form.value.id, form.value.name, form.value.status)
      dialogFormVisible.value = false
    }

    onMounted(() => {
      form.value = props.data
    })

    return {
      form,
      dialogFormVisible,
      handleUpdate,
    }
  },
})
</script>

<style scoped>

</style>