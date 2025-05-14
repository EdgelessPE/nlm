<template>
  <div class="w-full flex items-center gap-4px">
    <div v-for="node in nodes" :key="node.field">
      <component :is="node.component" />
    </div>
  </div>
</template>

<script setup lang="ts">
import type { FilterProps } from "./type";
import { computed } from "vue";

const props = defineProps<FilterProps<any>>();

const nodes = computed(() => props.getConfig(props.model));
</script>
