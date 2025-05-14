<template>
  <div class="flex flex-col items-center h-full overflow-hidden">
    <Filter v-bind="bindFilterProps" />
    <Table v-bind="bindTableProps" class="w-full mt-8px" />
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
import { GetNeps, type GetNepsParams, type Nep } from "@/api/nep";
import Table from "@/components/table/index.vue";
import { defineTableProps } from "@/components/table/utils";
import { renderDate, renderActions } from "@/components/table/renders";
import Button from "primevue/button";
import DebouncedSearch from "@/components/DebouncedSearch.vue";
import { computed, nextTick, ref } from "vue";
import Popover from "primevue/popover";
import { useClipboard } from "@vueuse/core";
import ReleasesDrawer from "./ReleasesDrawer.vue";
import { defineFilterProps } from "@/components/filter/utils";
import Filter from "@/components/filter/index.vue";
import AsyncSelect from "@/components/async-select/index.vue";
import { GetScopes } from "@/api/nep";

type IFilter = GetNepsParams;
const query = ref<IFilter>({});

const bindFilterProps = defineFilterProps<IFilter>({
  model: query,
  getConfig: (form) => [
    {
      field: "q",
      component: () => (
        <DebouncedSearch
          placeholder="Search scope or name"
          v-model={form.value.q}
        />
      ),
    },
    {
      field: "scope",
      component: () => (
        <AsyncSelect
          v-model={form.value.scope}
          fetch={GetScopes}
          placeholder="Scope"
        />
      ),
    },
  ],
});
const bindTableProps = defineTableProps<Nep>({
  fetch: GetNeps,
  query,
  getColumns: () => [
    {
      label: "Scope",
      field: "Scope",
      render: ({ val }) => (
        <Button
          variant="link"
          label={val}
          onClick={() => (query.value.scope = val)}
        />
      ),
      sortable: true,
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
      sortable: true,
    },
    {
      label: "Actions",
      field: "actions",
      render: renderActions([
        {
          key: "install",
          label: "Install",
          icon: () => <div class="pi pi-plus" />,
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
