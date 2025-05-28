<template>
  <Filter v-bind="bindFilterProps" class="mt-8px" />
  <Table v-bind="bindProps" class="mt-8px" />
</template>

<script setup lang="tsx">
import {
  GetPipelineList,
  type Pipeline,
  type PipelineParams,
} from "@/api/pipeline";
import { defineTableProps } from "@/components/table/utils";
import Table from "@/components/table/index.vue";
import {
  renderActions,
  renderDate,
  renderTag,
} from "@/components/table/renders";
import dayjs from "dayjs";
import { defineFilterProps } from "@/components/filter/utils";
import { ref } from "vue";
import Select from "primevue/select";

type IFilter = PipelineParams;
const query = ref<IFilter>({});

const bindFilterProps = defineFilterProps<IFilter>({
  model: query,
  getConfig: (form) => [
    {
      field: "status",
      component: () => (
        <Select
          v-model={form.value.status}
          placeholder="Status"
          options={["running", "success", "failed"]}
          showClear
        />
      ),
    },
  ],
});

const bindProps = defineTableProps<Pipeline, PipelineParams>({
  getColumns: () => [
    {
      label: "Name",
      field: "CreatedAt",
      render: ({ val }) => {
        const str = dayjs(val).format("YYYY-MM-DD--HH-mm-ss");
        return <span>{str}</span>;
      },
    },
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
    // {
    //   label: "Created At",
    //   field: "CreatedAt",
    //   render: renderDate(),
    // },
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
    {
      label: "Actions",
      field: "actions",
      render: renderActions<Pipeline>([
        {
          key: "view",
          label: "View Log",
          icon: () => <i class="pi pi-eye" />,
          onClick: (data) => {
            console.log(data);
          },
        },
      ]),
    },
  ],
  query,
  fetch: (params) => GetPipelineList({ ...params, model_name: "bot" }),
});
</script>
