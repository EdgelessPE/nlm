<template>
  <Dialog
    :visible="visible"
    :header="`Meta of ${release?.FileName}`"
    class="w-40vw"
    @update:visible="
      (v) => {
        if (!v) {
          $emit('close');
        }
      }
    "
  >
    <code class="whitespace-pre">{{
      release?.Meta
        ? JSON.stringify(JSON.parse(base64Decode(release.Meta)), null, 2)
        : "--"
    }}</code>
  </Dialog>
</template>

<script setup lang="ts">
import type { Release } from "@/api/nep";
import { computed } from "vue";
import { base64Decode } from "@/utils";
const props = defineProps<{
  release: Release | null;
}>();
defineEmits(["close"]);

const visible = computed(() => props.release !== null);
</script>
