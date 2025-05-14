import type { BaseResponse, BasicTableParams } from "@/api/type";
import type { AxiosResponse } from "axios";
import type { Ref, VNodeChild } from "vue";
import type { DataTableProps, DataTableSortEvent } from "primevue/datatable";
export type TableColumnRenderContext<IData = any> = {
  val: any;
  data: IData;
};

export type TableColumn<IData = any> = {
  label: string;
  field: Exclude<keyof IData, symbol | number> | "actions";
  render?: (ctx: TableColumnRenderContext<IData>) => VNodeChild;
  sortable?: boolean;
};

export interface TableRuntimeParams {
  offset: number;
  limit: number;
  sort?: number;
  sortBy?: string;
}

type TablePropsEnhanced = DataTableProps & {
  slotHeight?: number;
};
export interface UseTableDataProps<IData = any, IFilter = any> {
  // 返回 false 时，不进行查询
  query?: Ref<IFilter | false>;
  fetch: (
    params: Omit<BasicTableParams, "q"> & IFilter,
  ) => Promise<AxiosResponse<BaseResponse<IData[]>> | undefined>;
  getColumns: (ctx: { refresh: () => void }) => TableColumn<IData>[];
  tableProps?: TablePropsEnhanced;
}

export interface UseTableDataReturn<IData = any> {
  loading: Ref<boolean>;
  data: Ref<IData[]>;
  columns: TableColumn<IData>[];
  runtimeParams: Ref<TableRuntimeParams>;
  total: Ref<number>;
  onSort: (e: DataTableSortEvent) => void;
  onPageChange: (val: { page: number; rows: number }) => void;
  tableProps?: TablePropsEnhanced;
}
