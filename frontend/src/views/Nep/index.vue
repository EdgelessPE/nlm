<template>
  <div class="flex flex-col items-center justify-center h-full">
    <Table v-bind="bindProps" class="h-full w-full">
      <template #header>
        <DebouncedSearch placeholder="Search scope or name" v-model="q" />
      </template>
    </Table>
    <Popover ref="op">
      <div class="flex items-center gap-2">
        <Button
          :label="copied ? 'Copied!' : 'Copy'"
          variant="text"
          size="small"
          class="!py-1px"
          @click="onCopy"
        />
        <code
          class="bg-gray-100 dark:bg-gray-800 rounded px-2 py-1 font-mono text-sm"
        >
          {{ installCommand }}
        </code>
        <span>and run!</span>
      </div>
    </Popover>
    <ReleasesDrawer :data="buildsData" @close="buildsData = null" />
  </div>
</template>

<script setup lang="tsx">
import { GetNeps, type Nep } from "@/api/nep";
import Table from "@/components/table/index.vue";
import { useTableData } from "@/components/table/useTableData.tsx";
import { renderDate, renderActions } from "@/components/table/renders";
import Button from "primevue/button";
import DebouncedSearch from "@/components/DebouncedSearch.vue";
import { computed, nextTick, ref } from "vue";
import Popover from "primevue/popover";
import { useClipboard } from "@vueuse/core";
import ReleasesDrawer from "./ReleasesDrawer.vue";
const q = ref<string>();
const bindProps = useTableData<Nep>({
  fetch: GetNeps,
  query: computed(() => ({
    q: q.value || undefined,
  })),
  getColumns: () => [
    {
      label: "Scope",
      field: "Scope",
      render: ({ val }) => (
        <Button variant="link" label={val} onClick={() => (q.value = val)} />
      ),
    },
    {
      label: "Name",
      field: "Name",
    },
    {
      label: "Latest Version",
      field: "LatestReleaseVersion",
    },
    {
      label: "Updated At",
      field: "UpdatedAt",
      render: renderDate(),
    },
    {
      label: "Actions",
      field: "actions",
      render: renderActions([
        {
          key: "install",
          label: "Install",
          icon: () => <div class="pi pi-cloud-download" />,
          onClick: (data, event) => {
            installData.value = data;
            togglePopover(event);
          },
        },
        {
          key: "view-releases",
          label: "View Releases",
          icon: () => <div class="pi pi-eye" />,
          onClick: (data) => {
            buildsData.value = data;
          },
        },
      ]),
    },
  ],
  tableProps: {
    slotHeight: 67,
  },
});

const op = ref<InstanceType<typeof Popover>>();
const installData = ref<Nep>();
const buildsData = ref<Nep | null>(null);
const togglePopover = async (event: Event) => {
  op.value?.hide();
  await nextTick();
  op.value?.show(event);
};

const installCommand = computed(
  () => `ept i "${installData.value?.Scope}/${installData.value?.Name}"`,
);
const { copy, copied } = useClipboard();
const onCopy = () => {
  copy(installCommand.value);
};
</script>
