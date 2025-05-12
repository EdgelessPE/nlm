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
    <code>{{ release?.Meta ? base64Decode(release.Meta) : "--" }}</code>
  </Dialog>
</template>

<script setup lang="ts">
import type { Nep, Release } from "@/api/nep";
import { computed } from "vue";
import { base64Decode } from "@/utils";
const props = defineProps<{
  release: Release | null;
  nep: Nep | null;
}>();
defineEmits(["close"]);

const visible = computed(() => props.release !== null);
</script>
