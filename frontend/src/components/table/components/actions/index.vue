<template>
  <div class="flex items-center gap-4px">
    <template v-for="action in actions" :key="action.key">
      <Button
        v-tooltip.top="action.label"
        variant="text"
        size="small"
        @click="action.onClick(ctx.data, $event)"
      >
        <template #icon>
          <component :is="action.icon" />
        </template>
      </Button>
    </template>
  </div>
</template>

<script setup lang="ts">
import type { ActionNode } from "./type";
import type { TableColumnRenderContext } from "../../type";

defineProps<{
  actions: ActionNode[];
  ctx: TableColumnRenderContext;
}>();
</script>
