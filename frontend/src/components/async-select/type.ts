import type { BaseResponse } from "@/api/type";
import type { AxiosResponse } from "axios";
import type { Ref } from "vue";

export interface AsyncSelectProps<IFilter = any> {
  fetch: (
    query?: IFilter,
  ) => Promise<
    AxiosResponse<BaseResponse<string[] | { label: string; value: string }[]>>
  >;
  query?: Ref<IFilter | false>;
  lazy?: boolean;
}
