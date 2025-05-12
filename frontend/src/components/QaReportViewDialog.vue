<template>
  <Dialog
    :visible="visible"
    :header="`QA report of ${release?.FileName}`"
    class="w-70vw"
    @update:visible="
      (v) => {
        if (!v) {
          $emit('close');
        }
      }
    "
  >
    114514
  </Dialog>
</template>

<script setup lang="ts">
import type { Release } from "@/api/nep";
import Dialog from "primevue/dialog";
import { computed } from "vue";
const props = defineProps<{
  release: Release | null;
}>();
defineEmits(["close"]);

const visible = computed(() => props.release !== null);
</script>
