<template>
  <Table v-bind="bindProps" />
</template>

<script setup lang="tsx">
import {
  GetPipelineList,
  type Pipeline,
  type PipelineParams,
} from "@/api/pipeline";
import { defineTableProps } from "@/components/table/utils";
import Table from "@/components/table/index.vue";
import { renderDate, renderTag } from "@/components/table/renders";

const bindProps = defineTableProps<Pipeline, PipelineParams>({
  getColumns: () => [
    // {
    //   label: "ID",
    //   field: "ID",
    // },
    {
      label: "Status",
      field: "Status",
      render: (params) => {
        return renderTag({
          map: {
            running: {
              label: "Running",
              severity: "info",
            },
            success: {
              label: "Success",
              severity: "success",
            },
            failed: {
              label: "Failed",
              severity: "danger",
            },
          },
          tooltip:
            params.data.Status === "failed" ? params.data.ErrMsg : undefined,
        })(params);
      },
    },
    {
      label: "Stage",
      field: "Stage",
    },
    {
      label: "Created At",
      field: "CreatedAt",
      render: renderDate(),
    },
    {
      label: "Finished At",
      field: "FinishedAt",
      render: (params) => {
        if (params.data.Status === "running") {
          return <span>--</span>;
        }
        return renderDate()(params);
      },
    },
  ],
  fetch: (params) => GetPipelineList({ ...params, model_name: "bot" }),
});
</script>
