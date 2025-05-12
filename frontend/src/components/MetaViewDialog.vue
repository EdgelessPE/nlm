<template>
  <Dialog
    v-model:visible="visible"
    :header="`Meta of ${nep?.Scope}/${nep?.Name}`"
    @hide="$emit('close')"
  >
    <pre>{{ data?.Meta }}</pre>
  </Dialog>
</template>

<script setup lang="ts">
import type { Nep, Release } from "@/api/nep";
import { computed } from "vue";

const props = defineProps<{
  data: Release | null;
  nep: Nep | null;
}>();
defineEmits(["close"]);

const visible = computed(() => props.data !== null);
</script>
