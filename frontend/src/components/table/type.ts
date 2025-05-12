import type { BaseResponse, BasicTableParams } from "@/api/type";
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

export interface TablePagination {
  offset: number;
  limit: number;
}

type TablePropsEnhanced = DataTableProps & {
  slotHeight?: number;
};
export interface UseTableDataProps<IData = any, IFilter = any> {
  query?: Ref<IFilter>;
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
  pagination: Ref<TablePagination>;
  total: Ref<number>;
  onPageChange: (val: { page: number; rows: number }) => void;
  tableProps?: TablePropsEnhanced;
}
