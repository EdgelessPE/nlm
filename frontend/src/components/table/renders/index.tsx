import dayjs from "dayjs";
import type { TableColumnRenderContext } from "../type";

export function renderDate({ val }: TableColumnRenderContext) {
  const d = dayjs(val);
  if (d.isValid()) {
    return <span>{d.format("YYYY-MM-DD HH:mm:ss")}</span>;
  }
  return <span>{val}</span>;
}
