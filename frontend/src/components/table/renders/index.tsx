import dayjs from "dayjs";
import type { TableColumnRenderContext } from "../type";
import Actions from "../components/actions/index.vue";
import type { ActionNode } from "../components/actions/type";

export function renderDate() {
  return ({ val }: TableColumnRenderContext) => {
    const d = dayjs(val);
    if (d.isValid()) {
      return <span>{d.format("YYYY-MM-DD HH:mm:ss")}</span>;
    }
    return <span>{val}</span>;
  };
}

export function renderActions<T>(actions: ActionNode<T>[]) {
  return (ctx: TableColumnRenderContext<T>) => {
    return <Actions actions={actions} ctx={ctx} />;
  };
}
