import instance from "./index";
import type { BaseResponse, BasicTableParams } from "./type";

export interface GetNepsParams extends BasicTableParams {
  scope?: string;
  updated_at_start?: number;
  updated_at_end?: number;
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
  return instance.get<BaseResponse<Nep[]>>("/api/nep/list", {
    params,
  });
}

export interface GetReleasesParams extends BasicTableParams {
  nep_id: string;
  is_bot_success?: boolean;
  is_qa_success?: boolean;
  version?: string;
  flags?: string;
  created_at_start?: number;
  created_at_end?: number;
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

export async function GetScopes() {
  return instance.get<BaseResponse<string[]>>("/api/nep/scopes");
}

export interface GetReleaseVersionsParams extends BasicTableParams {
  id: string;
}
export async function GetReleaseVersions(params: GetReleaseVersionsParams) {
  return instance.get<BaseResponse<string[]>>("/api/nep/release_versions", {
    params,
  });
}
