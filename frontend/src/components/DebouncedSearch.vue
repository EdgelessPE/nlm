<template>
  <IconField ref="inputRef" class="w-320px">
    <InputIcon :class="icon" />
    <InputText v-model="innerModel" :placeholder="placeholder" class="w-full" />
    <InputIcon
      v-if="innerModel && isHoverInput"
      class="pi pi-times-circle"
      @click="outerModel = undefined"
    />
  </IconField>
</template>

<script setup lang="ts">
import { ref, useTemplateRef, watch } from "vue";
import { useDebounceFn, useElementHover } from "@vueuse/core";

const outerModel = defineModel<string>();
const innerModel = ref<string>();

const props = withDefaults(
  defineProps<{
    placeholder?: string;
    debounce?: number;
    icon?: string;
  }>(),
  {
    placeholder: "Search",
    debounce: 500,
    icon: "pi pi-search",
  },
);

const debouncedInput = useDebounceFn((value?: string) => {
  outerModel.value = value;
}, props.debounce);

watch(innerModel, debouncedInput);
watch(outerModel, (val) => {
  innerModel.value = val;
});

const inputRef = useTemplateRef<HTMLInputElement>("inputRef");
const isHoverInput = useElementHover(inputRef);
</script>
