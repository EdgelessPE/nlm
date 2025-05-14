import { computedAsync } from "@vueuse/core";
import type {
  TableRuntimeParams,
  UseTableDataProps,
  UseTableDataReturn,
} from "./type";
import { ref } from "vue";
import type { DataTableSortEvent } from "primevue/datatable";

export function defineTableProps<T>(
  props: UseTableDataProps<T>,
): UseTableDataReturn {
  const loading = ref(false);
  const total = ref(1);
  const runtimeParams = ref<TableRuntimeParams>({
    offset: 0,
    limit: 20,
  });
  const onPageChange = (val: { page: number; rows: number }) => {
    runtimeParams.value.offset = val.page * val.rows;
    runtimeParams.value.limit = val.rows;
  };

  const refreshKey = ref(0);
  const data = computedAsync(
    async () => {
      // eslint-disable-next-line @typescript-eslint/no-unused-expressions
      refreshKey.value;
      const {
        data: { data, total: t },
      } = (await props.fetch({
        ...runtimeParams.value,
        ...props.query?.value,
      })) ?? {
        data: {
          data: [],
          total: 0,
        },
      };
      total.value = t ?? 1;
      return data;
    },
    [],
    {
      evaluating: loading,
    },
  );

  const columns = props
    .getColumns({
      refresh: () => {
        refreshKey.value++;
      },
    })
    .map((column) => ({
      ...column,
      render: column.render ?? (({ val }) => <span>{val || "--"}</span>),
    }));

  const onSort = (e: DataTableSortEvent) => {
    runtimeParams.value.offset = 0;
    if (e.sortField) {
      runtimeParams.value.sort = e.sortOrder || undefined;
      runtimeParams.value.sortBy = e.sortField as string;
    } else {
      runtimeParams.value.sort = undefined;
      runtimeParams.value.sortBy = undefined;
    }
  };

  return {
    total,
    loading,
    data,
    columns,
    runtimeParams,
    tableProps: props.tableProps,
    onSort,
    onPageChange,
  };
}
