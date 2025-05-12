<template>
  <div class="flex items-center gap-4px">
    <template v-for="action in actions" :key="action.key">
      <Button
        v-tooltip.top="action.label"
        variant="text"
        size="small"
        :disabled="action.disabled?.(ctx.data)"
        :loading="loadingMap[action.key]"
        @click="handleClick(action, ctx.data, $event)"
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
import { ref } from "vue";
defineProps<{
  actions: ActionNode[];
  ctx: TableColumnRenderContext;
}>();
const loadingMap = ref<Record<string, boolean>>({});
const handleClick = async (action: ActionNode, data: any, event: Event) => {
  const r = action.onClick(data, event);
  if (r instanceof Promise) {
    loadingMap.value[action.key] = true;
    await r;
    loadingMap.value[action.key] = false;
  }
};
</script>
