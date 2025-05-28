import instance from "./index";
import type { BaseResponse, BasicTableParams } from "./type";

export type PipelineParams = BasicTableParams & {
  model_name?: string;
  status?: "running" | "success" | "failed";
};
export interface Pipeline {
  ID: string;
  CreatedAt: string;
  UpdatedAt: string;
  DeletedAt: string | null;
  ModelName: string;
  FinishedAt: string | null;
  Status: "running" | "success" | "failed";
  ErrMsg: string;
  Stage: string;
}
export async function GetPipelineList(params: PipelineParams) {
  return instance.get<BaseResponse<Pipeline[]>>("/api/pipeline/list", {
    params,
  });
}
