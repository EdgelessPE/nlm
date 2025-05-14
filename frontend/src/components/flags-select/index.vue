<template>
  <MultiSelect
    v-model="innerModel"
    :options="FLAGS"
    :showToggleAll="false"
    placeholder="Flags"
    class="w-120px"
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
  const arr = val?.join("");
  const newVal = arr?.length ? arr : undefined;
  if (newVal !== outerModel.value) {
    outerModel.value = newVal;
  }
});
watch(outerModel, (val) => {
  innerModel.value = val ? val.split("") : [];
});
</script>
