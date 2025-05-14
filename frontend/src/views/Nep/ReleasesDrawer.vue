<template>
  <Drawer
    :visible="visible"
    class="!w-80vw"
    position="right"
    :header="`Releases of ${data?.Scope}/${data?.Name}`"
    @update:visible="
      (v) => {
        if (!v) {
          $emit('close');
        }
      }
    "
  >
    <Filter v-bind="bindFilterProps" />
    <Table v-bind="bindTableProps" class="h-full w-full mt-8px" />
  </Drawer>
  <MetaViewDialog
    :release="metaDialogRelease"
    @close="metaDialogRelease = null"
  />
  <QaReportViewDialog
    :release="qaReportDialogRelease"
    @close="qaReportDialogRelease = null"
  />
</template>

<script setup lang="tsx">
import { computed, ref } from "vue";
import {
  GetReleases,
  GetReleaseVersions,
  type GetReleasesParams,
  type Nep,
  type Release,
} from "@/api/nep";
import Drawer from "primevue/drawer";
import { defineTableProps } from "@/components/table/utils";
import {
  renderActions,
  renderDate,
  renderFileSize,
  renderSuccess,
} from "@/components/table/renders";
import { Button } from "primevue";
import MetaViewDialog from "@/components/MetaViewDialog.vue";
import QaReportViewDialog from "@/components/QaReportViewDialog.vue";
import { GetStorageUrl } from "@/api/storage";
import { defineFilterProps } from "@/components/filter/utils";
import DebouncedSearch from "@/components/DebouncedSearch.vue";
import LastMajorTag from "@/components/LastMajorTag.vue";
import FlagsSelect from "@/components/flags-select/index.vue";
import AsyncSelect from "@/components/async-select/index.vue";
import BooleanSelect from "@/components/boolean-select.vue";

type IFilter = Omit<GetReleasesParams, "nep_id">;

const props = defineProps<{
  data: Nep | null;
}>();
defineEmits(["close"]);

const visible = computed(() => props.data !== null);
const query = ref<IFilter>({});
const metaDialogRelease = ref<Release | null>(null);
const qaReportDialogRelease = ref<Release | null>(null);

const versionQuery = computed(() => ({
  id: props.data?.ID,
}));

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
      field: "version",
      component: () => (
        <AsyncSelect
          v-model={form.value.version}
          fetch={GetReleaseVersions}
          placeholder="Version"
          query={versionQuery}
        />
      ),
    },
    {
      field: "flags",
      component: () => <FlagsSelect v-model={form.value.flags} />,
    },
    {
      field: "is_bot_success",
      component: () => (
        <BooleanSelect
          v-model={form.value.is_bot_success}
          placeholder="Bot Success"
        />
      ),
    },
    {
      field: "is_qa_success",
      component: () => (
        <BooleanSelect
          v-model={form.value.is_qa_success}
          placeholder="QA Success"
        />
      ),
    },
  ],
});
const bindTableProps = defineTableProps<Release>({
  query: computed(() => ({
    nep_id: props.data?.ID,
    ...query.value,
  })),
  getColumns: () => [
    {
      label: "File Name",
      field: "FileName",
    },
    {
      label: "Version",
      field: "Version",
      render: ({ data }) => {
        return (
          <div class="flex items-center gap-2">
            <span>{data.Version}</span>
            {data.IsLastMajor && <LastMajorTag />}
          </div>
        );
      },
    },
    {
      label: "Flags",
      field: "Flags",
      render: ({ val }: { val: string }) => {
        return (
          <Button
            variant="outlined"
            size="small"
            label={val.trim()}
            onClick={() => {
              query.value.flags = val;
            }}
          />
        );
      },
    },
    {
      label: "File Size",
      field: "FileSize",
      render: renderFileSize(),
    },
    {
      label: "Meta",
      field: "Meta",
      render: ({ val, data }) => {
        return val ? (
          <Button
            label="View"
            variant="outlined"
            size="small"
            onClick={() => (metaDialogRelease.value = data)}
          />
        ) : (
          "--"
        );
      },
    },
    {
      label: "Bot Success",
      field: "IsBotSuccess",
      render: (ctx) => {
        return renderSuccess({
          tooltip: ctx.val ? undefined : ctx.data.BotErrMsg || undefined,
        })(ctx);
      },
    },
    {
      label: "QA Success",
      field: "IsQaSuccess",
      render: renderSuccess(),
    },
    {
      label: "QA Report",
      field: "QaResultStorageKey",
      render: ({ val, data }) => {
        return val ? (
          <Button
            label="View"
            variant="outlined"
            size="small"
            onClick={() => (qaReportDialogRelease.value = data)}
          />
        ) : (
          "--"
        );
      },
    },
    {
      label: "Created At",
      field: "CreatedAt",
      render: renderDate(),
    },
    {
      label: "Actions",
      field: "actions",
      render: renderActions([
        {
          key: "download",
          label: "Download",
          icon: () => <div class="pi pi-cloud-download" />,
          disabled: (data) => !data.StorageKey,
          onClick: async (data) => {
            const {
              data: { data: url },
            } = await GetStorageUrl(data.StorageKey);
            window.open(url, "_blank");
          },
        },
        {
          key: "view-pipeline",
          label: "View Pipeline",
          icon: () => <div class="pi pi-truck" />,
          onClick: async (data) => {
            window.open(`/pipeline?id=${data.PipelineId}`, "_blank");
          },
        },
      ]),
    },
  ],
  fetch: async (params) => {
    if (!params.nep_id) {
      return undefined;
    }
    return GetReleases(params);
  },
});
</script>
