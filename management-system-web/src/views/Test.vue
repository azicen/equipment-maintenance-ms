<template>
  <div class="test">
    <NewUserForm />
    <NewGroupForm />
    <NewEquipmentTypeForm />
    <NewEquipmentForm />
  </div>
</template>

<script>
import NewEquipmentForm from "@/components/admin/form/NewEquipmentForm.vue";
import NewEquipmentTypeForm from "@/components/admin/form/NewEquipmentTypeForm.vue";
import NewGroupForm from "@/components/admin/form/NewGroupForm.vue";
import NewUserForm from "@/components/admin/form/NewUserForm.vue";

export default {
  name: "Test",
  components: {
    NewEquipmentTypeForm,
    NewEquipmentForm,
    NewGroupForm,
    NewUserForm,
  },
};
</script>
