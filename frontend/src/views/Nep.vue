<template>
  <div class="flex flex-col items-center justify-center h-full">
    <Table v-bind="bindProps" />
  </div>
</template>

<script setup lang="ts">
import { GetNeps } from "@/api/nep";
import Table from "@/components/table/index.vue";
import { useTableData } from "@/components/table/useTableData";
import { renderDate } from "@/components/table/renders";
const bindProps = useTableData({
  fetch: GetNeps,
  getColumns: (ctx) => [
    {
      label: "Scope",
      field: "Scope",
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
      render: renderDate,
    },
  ],
});
</script>
