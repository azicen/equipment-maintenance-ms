<template>

  <el-dialog v-model="visibleBind" :title="'修改用户 '+ form.id +' 的信息'" width="30%">

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
        <el-button @click="visibleBind = false">取消</el-button>
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
    'visible': Boolean,
    'data': Object as PropType<GetUserResponse>,
  },

  // 双向绑定的回调事件
  emits: ['update:visible', 'update:data'],

  setup(props, context) {
    const userApi = new UserApi()

    // 嵌套双向绑定
    const visibleBind = computed({
      get: () => props.visible,
      set: (v) => {
        context.emit('update:visible', v)
        console.log('update:visible')
      }
    })
    const dataBind = computed({
      get: () => props.data,
      set: (v) => {
        context.emit('update:data', v)
        if (props.data === undefined) {
          return
        }
        form.value = props.data
        console.log('update:data')
      }
    })

    // 表单数据
    const form = ref<GetUserResponse>({id: 0, name: '', status: true})

    function handleUpdate() {
      userApi.updateUser(form.value.id, form.value.name, form.value.status)
      visibleBind.value = false
    }

    onMounted(() => {
      if (props.data === undefined) {
        return
      }
      form.value = props.data
      console.log('UpdateUserDialog',form.value)
    })

    return {
      form,
      visibleBind,
      dataBind,
      handleUpdate,
    }
  },
})
</script>

<style scoped>

</style>