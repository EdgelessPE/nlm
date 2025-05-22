import instance from "./index";
import type { BaseResponse } from "./type";
import type { BasicTableParams } from "./type";

export type GetEptsParams = BasicTableParams;

export interface Ept {
  ID: string;
  CreatedAt: string;
  UpdatedAt: string;
  DeletedAt: string | null;
  Name: string;
  Version: string;
  StorageKey: string;
  FileSize: number;
  Integrity: string;
}

export async function GetEpts(params: GetEptsParams) {
  return instance.get<BaseResponse<Ept[]>>("/api/ept/epts", {
    params,
  });
}
