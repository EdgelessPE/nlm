import instance from ".";
import type { BaseResponse } from "./type";

export function GetStorageUrl(uuid: string) {
  return instance.get<BaseResponse<string>>(`/api/storage/url/${uuid}`);
}
