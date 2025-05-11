<template>
  <Drawer
    v-model:visible="visible"
    class="!w-60vw"
    position="right"
    :header="`Releases of ${data?.Scope}/${data?.Name}`"
    @hide="$emit('close')"
  >
    {{ data }}
  </Drawer>
</template>

<script setup lang="tsx">
import { computed } from "vue";
import type { Nep } from "@/api/nep";
import Drawer from "primevue/drawer";

const props = defineProps<{
  data: Nep | null;
}>();
defineEmits(["close"]);

const visible = computed(() => props.data !== null);
</script>
