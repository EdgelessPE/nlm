<template>
  <div class="container">
    <Filter v-bind="bindFilterProps" />
    <Table v-bind="bindTableProps" class="w-full mt-8px" />
  </div>
</template>

<script setup lang="tsx">
import Table from "@/components/table/index.vue";
import { defineTableProps } from "@/components/table/utils";
import { ref } from "vue";
import { defineFilterProps } from "@/components/filter/utils";
import DebouncedSearch from "@/components/DebouncedSearch.vue";
import {
  renderActions,
  renderDate,
  renderFileSize,
  renderBoolean,
} from "@/components/table/renders";
import {
  GetStorages,
  GetStorageUrl,
  type GetStoragesParams,
  type Storage,
} from "@/api/storage";
import BooleanSelect from "@/components/boolean-select.vue";

type IFilter = GetStoragesParams;
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
    {
      field: "is_compressed",
      component: () => (
        <BooleanSelect
          v-model={form.value.is_compressed}
          placeholder="Is Compressed"
        />
      ),
    },
  ],
});
const bindTableProps = defineTableProps<Storage>({
  fetch: GetStorages,
  query,
  getColumns: () => [
    {
      label: "File Name",
      field: "FileName",
    },
    {
      label: "File Size",
      field: "FileSize",
      render: renderFileSize(),
      sortable: true,
    },
    {
      label: "Sync Finished At",
      field: "SyncFinishedAt",
      render: renderDate(),
      sortable: true,
    },
    {
      label: "Compressed",
      field: "Compressed",
      render: renderBoolean(),
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
            } = await GetStorageUrl(data.ID);
            window.open(url, "_blank");
          },
        },
      ]),
    },
  ],
});
</script>
