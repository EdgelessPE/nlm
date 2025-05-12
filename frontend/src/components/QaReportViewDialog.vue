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
    <div v-html="md" />
  </Dialog>
</template>

<script setup lang="ts">
import type { Release } from "@/api/nep";
import { GetStorageUrl } from "@/api/storage";
import { computedAsync } from "@vueuse/core";
import Dialog from "primevue/dialog";
import { computed, ref } from "vue";
import markdownIt from "markdown-it";
import { decodeZstd } from "@/utils";
const props = defineProps<{
  release: Release | null;
}>();
defineEmits(["close"]);

const visible = computed(() => props.release !== null);
const loading = ref(false);
const md = computedAsync<string>(
  async () => {
    // 获取 url
    const {
      data: { data: url },
    } = await GetStorageUrl(props.release!.QaResultStorageKey);
    // 下载 zstd 压缩的 md 内容
    const response = await fetch(url, {
      method: "GET",
      mode: "no-cors",
    });
    const data = await response.arrayBuffer();

    // 解压
    const decompressed = await decodeZstd(
      new Uint8Array(data),
      props.release?.FileSize ?? 0,
    );
    // 转换为字符串
    const str = new TextDecoder().decode(decompressed);
    // 转换为 markdown
    return markdownIt().render(str);
  },
  "",
  { evaluating: loading },
);
</script>
