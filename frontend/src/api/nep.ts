import instance from "./index";
import type { BaseResponse } from "./type";

export interface Nep {
  ID: string;
  CreatedAt: string;
  UpdatedAt: string;
  DeletedAt: string | null;
  Scope: string;
  Name: string;
  LatestReleaseVersion: string;
}
export async function GetNeps(params: { offset: number; limit: number }) {
  return instance.get<BaseResponse<Nep[]>>("/api/nep/neps", {
    params,
  });
}
