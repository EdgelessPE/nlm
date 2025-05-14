<template>
  <Badge v-tooltip.top="tooltip" :severity="severity">
    <div class="flex gap-4px items-center">
      {{ value }}
      <div v-if="tooltip" class="pi pi-info-circle" />
    </div>
  </Badge>
</template>

<script setup lang="ts">
import { Badge } from "primevue";

defineProps<{
  value: string;
  tooltip?: string;
  severity?: "success" | "danger" | "warning" | "info" | "help" | "none";
}>();
</script>
