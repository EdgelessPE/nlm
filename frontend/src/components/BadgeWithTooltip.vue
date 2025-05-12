<template>
  <Badge v-tooltip.top="tooltip" :value="value" :severity="severity" />
</template>

<script setup lang="ts">
import { Badge } from "primevue";

defineProps<{
  value: string;
  tooltip?: string;
  severity?: "success" | "danger" | "warning" | "info" | "help" | "none";
}>();
</script>
