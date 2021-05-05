<template>
  <div class="test">
    <NewEmployeeForm />
    <NewGroupForm />
    <NewEquipmentTypeForm />
    <NewEquipmentForm />
  </div>
</template>

<script>
import NewEquipmentForm from "@/components/admin/form/NewEquipmentForm.vue";
import NewEquipmentTypeForm from "@/components/admin/form/NewEquipmentTypeForm.vue";
import NewGroupForm from "@/components/admin/form/NewGroupForm.vue";
import NewEmployeeForm from "@/components/admin/form/NewEmployeeForm.vue";

export default {
  name: "Test",
  components: {
    NewEquipmentTypeForm,
    NewEquipmentForm,
    NewGroupForm,
    NewEmployeeForm,
  },
};
</script>
