<template>
  <div class="container">
    <Filter v-bind="bindFilterProps" />
    <Table v-bind="bindTableProps" class="w-full mt-8px" />
  </div>
</template>

<script setup lang="tsx">
import Table from "@/components/table/index.vue";
import { defineTableProps } from "@/components/table/utils";
import { GetEpts, type Ept, type GetEptsParams } from "@/api/ept";
import { ref } from "vue";
import { defineFilterProps } from "@/components/filter/utils";
import DebouncedSearch from "@/components/DebouncedSearch.vue";
import {
  renderActions,
  renderDate,
  renderFileSize,
} from "@/components/table/renders";
import { GetStorageUrl } from "@/api/storage";

type IFilter = GetEptsParams;
const query = ref<IFilter>({});

const bindFilterProps = defineFilterProps<IFilter>({
  model: query,
  getConfig: (form) => [
    {
      field: "q",
      component: () => (
        <DebouncedSearch
          placeholder="Search file name"
          v-model={form.value.q}
        />
      ),
    },
  ],
});
const bindTableProps = defineTableProps<Ept>({
  fetch: GetEpts,
  query,
  getColumns: () => [
    {
      label: "Name",
      field: "Name",
    },
    {
      label: "Version",
      field: "Version",
      sortable: true,
    },
    {
      label: "File Size",
      field: "FileSize",
      render: renderFileSize(),
      sortable: true,
    },
    {
      label: "Integrity (blake3)",
      field: "Integrity",
    },
    {
      label: "Created At",
      field: "CreatedAt",
      render: renderDate(),
      sortable: true,
    },
    {
      label: "Actions",
      field: "actions",
      render: renderActions([
        {
          key: "download",
          label: "Download",
          icon: () => <div class="pi pi-download" />,
          onClick: async (data) => {
            const {
              data: { data: url },
            } = await GetStorageUrl(data.StorageKey);
            window.open(url, "_blank");
          },
        },
      ]),
    },
  ],
});
</script>
