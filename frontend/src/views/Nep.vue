<template>
  <div class="flex flex-col items-center justify-center h-full">
    <Table v-bind="bindProps" class="h-full w-full">
      <template #header>
        <DebouncedSearch placeholder="Search scope or name" v-model="q" />
      </template>
    </Table>
  </div>
</template>

<script setup lang="tsx">
import { GetNeps, type Nep } from "@/api/nep";
import Table from "@/components/table/index.vue";
import { useTableData } from "@/components/table/useTableData.tsx";
import { renderDate, renderActions } from "@/components/table/renders";
import Button from "primevue/button";
import DebouncedSearch from "@/components/DebouncedSearch.vue";
import { computed, ref } from "vue";

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
          onClick: (data) => console.log(data),
        },
        {
          key: "view-builds",
          label: "View Builds",
          icon: () => <div class="pi pi-eye" />,
          onClick: (data) => console.log(data),
        },
      ]),
    },
  ],
  tableProps: {
    slotHeight: 67,
  },
});
</script>
