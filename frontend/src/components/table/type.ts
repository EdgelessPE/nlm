import type { BaseResponse } from "@/api/type";
import type { AxiosResponse } from "axios";
import type { Ref, VNodeChild } from "vue";

export type TableColumnRenderContext<IData = any> = {
  val: any;
  data: IData;
};

export type TableColumn<IData = any> = {
  label: string;
  field: Exclude<keyof IData, symbol | number>;
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
}

export interface UseTableDataReturn<IData = any> {
  loading: Ref<boolean>;
  data: Ref<IData[]>;
  columns: TableColumn<IData>[];
  pagination: Ref<TablePaginationParams>;
  total: Ref<number>;
  onPageChange: (val: { page: number; rows: number }) => void;
}
