<template>
  <DatePicker
    v-model="innerModel"
    class="w-320px"
    selectionMode="range"
    :manualInput="false"
    showTime
    :placeholder="placeholder"
    showButtonBar
    @clear-click="innerModel = undefined"
  />
</template>

<script setup lang="ts">
import dayjs from "dayjs";
import DatePicker from "primevue/datepicker";
import { ref, watch } from "vue";

defineProps<{
  placeholder?: string;
}>();
const outerStart = defineModel<number>("start");
const outerEnd = defineModel<number>("end");
const innerModel = ref<[Date | null, Date | null]>();

function dateToUnix(date: Date | null) {
  if (date == null) {
    return 0;
  }
  return dayjs(date).unix();
}
function isEqual(
  outer: [number | undefined, number | undefined] | undefined,
  inner: [Date | null, Date | null] | undefined,
) {
  // 两个都是 undefined
  if (outer == null && inner == null) {
    return true;
  }
  // 一个 undefined
  if (outer == null || inner == null) {
    return false;
  }
  return dateToUnix(inner[0]) === outer[0] && dateToUnix(inner[1]) === outer[1];
}

watch(innerModel, (val) => {
  if (val) {
    if (!isEqual([outerStart.value, outerEnd.value], val) && val[0] && val[1]) {
      outerStart.value = dateToUnix(val[0]);
      outerEnd.value = dateToUnix(val[1]);
    }
  } else {
    outerStart.value = undefined;
    outerEnd.value = undefined;
  }
});
watch([outerStart, outerEnd], (val) => {
  if (val[0] && val[1]) {
    if (!isEqual(val, innerModel.value)) {
      innerModel.value = [dayjs(val[0]).toDate(), dayjs(val[1]).toDate()];
    }
  } else {
    innerModel.value = undefined;
  }
});
</script>
