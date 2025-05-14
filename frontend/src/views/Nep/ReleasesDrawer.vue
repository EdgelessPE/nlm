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
    <Table v-bind="bindProps" class="h-full w-full">
      <template #header>
        <DebouncedSearch placeholder="Search file name" v-model="q" />
      </template>
    </Table>
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
import { GetReleases, type Nep, type Release } from "@/api/nep";
import Drawer from "primevue/drawer";
import { defineTableProps } from "@/components/table/utils";
import {
  renderActions,
  renderDate,
  renderFileSize,
  renderSuccess,
} from "@/components/table/renders";
import Badge from "primevue/badge";
import { Button } from "primevue";
import MetaViewDialog from "@/components/MetaViewDialog.vue";
import QaReportViewDialog from "@/components/QaReportViewDialog.vue";
import { GetStorageUrl } from "@/api/storage";
const props = defineProps<{
  data: Nep | null;
}>();
defineEmits(["close"]);

const visible = computed(() => props.data !== null);
const q = ref("");
const metaDialogRelease = ref<Release | null>(null);
const qaReportDialogRelease = ref<Release | null>(null);
const bindProps = defineTableProps<Release>({
  query: computed(() => ({
    nep_id: props.data?.ID,
    q: q.value,
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
            {data.IsLastMajor && <Badge value="LastMajor" severity="success" />}
          </div>
        );
      },
    },
    {
      label: "Flags",
      field: "Flags",
      render: ({ val }) => {
        const flags = (val as string).split("").map((flag) => {
          return <Button variant="outlined" size="small" label={flag.trim()} />;
        });
        return <div class="flex gap-2">{flags}</div>;
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
