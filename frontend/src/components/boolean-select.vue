<template>
  <Select
    v-model="model"
    :options="options"
    :placeholder="placeholder"
    optionLabel="label"
    optionValue="value"
    showClear
  />
</template>

<script setup lang="ts">
import Select from "primevue/select";

const model = defineModel<boolean>();
defineProps<{
  placeholder?: string;
}>();
const options = [
  { label: "Yes", value: true },
  { label: "No", value: false },
];
</script>
