import { computedAsync } from "@vueuse/core";
import type {
  TablePaginationParams,
  UseTableDataProps,
  UseTableDataReturn,
} from "./type";
import { ref } from "vue";

export function useTableData<T>(
  props: UseTableDataProps<T>,
): UseTableDataReturn {
  const loading = ref(false);
  const total = ref(1);
  const pagination = ref<TablePaginationParams>({
    offset: 0,
    limit: 20,
  });
  const onPageChange = (val: { page: number; rows: number }) => {
    pagination.value.offset = val.page * val.rows;
    pagination.value.limit = val.rows;
  };

  const refreshKey = ref(0);
  const data = computedAsync(
    async () => {
      // eslint-disable-next-line @typescript-eslint/no-unused-expressions
      refreshKey.value;
      const {
        data: { data, total: t },
      } = await props.fetch({
        ...pagination.value,
        ...props.query?.value,
      });
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

  return {
    total,
    loading,
    data,
    columns,
    pagination,

    onPageChange,
  };
}
