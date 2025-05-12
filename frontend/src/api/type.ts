export interface BaseResponse<T> {
  code: number;
  msg: string;
  data: T;
  total?: number;
}

export interface BasicTableParams {
  offset?: number;
  limit?: number;
  q?: string;
  sort?: string;
  order?: string;
}
