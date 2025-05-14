<template>
  <MultiSelect
    v-model="innerModel"
    :options="FLAGS"
    :showToggleAll="false"
    placeholder="Flags"
    class="w-96px"
    showClear
  />
</template>

<script setup lang="ts">
import { ref, watch } from "vue";
import { FLAGS } from "./const";

const outerModel = defineModel<string>();
const innerModel = ref<string[]>(
  outerModel.value ? outerModel.value.split("") : [],
);

watch(innerModel, (val) => {
  const newVal = val?.join("");
  if (newVal !== outerModel.value) {
    outerModel.value = newVal;
  }
});
watch(outerModel, (val) => {
  innerModel.value = val ? val.split("") : [];
});
</script>
