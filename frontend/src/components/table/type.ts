import type { BaseResponse } from "@/api/type";
import type { AxiosResponse } from "axios";
import type { Ref, VNodeChild } from "vue";
import type { DataTableProps } from "primevue/datatable";
export type TableColumnRenderContext<IData = any> = {
  val: any;
  data: IData;
};

export type TableColumn<IData = any> = {
  label: string;
  field: Exclude<keyof IData, symbol | number> | "actions";
  render?: (ctx: TableColumnRenderContext<IData>) => VNodeChild;
};

export interface TablePaginationParams {
  offset: number;
  limit: number;
}

export interface UseTableDataProps<IData = any, IFilter = any> {
  query?: Ref<IFilter>;
  fetch: (
    params: TablePaginationParams & IFilter,
  ) => Promise<AxiosResponse<BaseResponse<IData[]>>>;
  getColumns: (ctx: { refresh: () => void }) => TableColumn<IData>[];
  tableProps?: DataTableProps;
}

export interface UseTableDataReturn<IData = any> {
  loading: Ref<boolean>;
  data: Ref<IData[]>;
  columns: TableColumn<IData>[];
  pagination: Ref<TablePaginationParams>;
  total: Ref<number>;
  onPageChange: (val: { page: number; rows: number }) => void;
  tableProps?: DataTableProps;
}
