import instance from ".";
import type { BaseResponse, BasicTableParams } from "./type";

export function GetStorageUrl(uuid: string) {
  return instance.get<BaseResponse<string>>(`/api/storage/url/${uuid}`);
}

export interface GetStoragesParams extends BasicTableParams {
  is_compressed?: boolean;
}
export interface Storage {
  ID: string;
  CreatedAt: string;
  UpdatedAt: string;
  DeletedAt: string | null;
  FileName: string;
  FileSize: number;
  SyncFinishedAt: string;
  Compressed: boolean;
}
export function GetStorages(params: GetStoragesParams) {
  return instance.get<BaseResponse<Storage[]>>(`/api/storage/list`, { params });
}
