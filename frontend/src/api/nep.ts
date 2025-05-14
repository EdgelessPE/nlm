import instance from "./index";
import type { BaseResponse, BasicTableParams } from "./type";

export interface GetNepsParams extends BasicTableParams {
  scope?: string;
}
export interface Nep {
  ID: string;
  CreatedAt: string;
  UpdatedAt: string;
  DeletedAt: string | null;
  Scope: string;
  Name: string;
  LatestReleaseVersion: string;
}
export async function GetNeps(params: GetNepsParams) {
  return instance.get<BaseResponse<Nep[]>>("/api/nep/neps", {
    params,
  });
}

export interface GetReleasesParams extends BasicTableParams {
  nep_id: string;
  is_bot_success?: boolean;
  is_qa_success?: boolean;
  version?: string;
}
export interface Release {
  ID: string;
  CreatedAt: string;
  UpdatedAt: string;
  DeletedAt: string | null;
  Version: string;
  Flags: string;
  FileName: string;
  FileSize: number;
  StorageKey: string;
  Meta: string;
  IsLastMajor: boolean;
  NepId: string;
  Nep: null;
  PipelineId: string;
  IsBotSuccess: boolean;
  BotErrMsg: string;
  IsQaSuccess: boolean;
  QaResultStorageKey: string;
}
export async function GetReleases(params: GetReleasesParams) {
  return instance.get<BaseResponse<Release[]>>("/api/nep/releases", {
    params,
  });
}
