<template>
  <Select
    v-bind="props"
    :options="options"
    v-model="model"
    :loading="loading"
    optionLabel="label"
    optionValue="value"
    filter
    filterPlaceholder="Search"
    resetFilterOnClear
    showClear
    @show="working = true"
  />
</template>

<script setup lang="ts">
import type { BaseResponse } from "@/api/type";
import { computedAsync } from "@vueuse/core";
import Select from "primevue/select";
import type { SelectProps } from "primevue/select";
import type { AxiosResponse } from "axios";
import { ref } from "vue";
const model = defineModel<SelectProps["modelValue"]>();
const props = withDefaults(
  defineProps<
    Omit<SelectProps, "options" | "modelValue"> & {
      fetch: () => Promise<
        AxiosResponse<
          BaseResponse<string[] | { label: string; value: string }[]>
        >
      >;
      lazy?: boolean;
    }
  >(),
  {
    lazy: () => true,
  },
);

const loading = ref(false);
const working = ref(!props.lazy);
const options = computedAsync<SelectProps["options"]>(
  async () => {
    if (!working.value) {
      return [];
    }
    const res = await props.fetch();
    return res.data.data
      .map((item) => {
        if (typeof item === "string") {
          return {
            label: item,
            value: item,
          };
        }
        return item;
      })
      .sort((a, b) => a.label.localeCompare(b.label));
  },
  [],
  {
    evaluating: loading,
  },
);
</script>
