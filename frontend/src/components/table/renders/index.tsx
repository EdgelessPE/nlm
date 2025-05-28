import dayjs from "dayjs";
import type { TableColumnRenderContext } from "../type";
import Actions from "../components/actions/index.vue";
import type { ActionNode } from "../components/actions/type";
import { formatFileSize } from "@/utils";
import BadgeWithTooltip from "@/components/BadgeWithTooltip.vue";

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

export function renderFileSize() {
  return ({ val }: TableColumnRenderContext) => {
    return <span>{formatFileSize(val)}</span>;
  };
}

export function renderBoolean({ tooltip }: { tooltip?: string } = {}) {
  return ({ val }: TableColumnRenderContext) => {
    return (
      <BadgeWithTooltip
        value={val ? "Yes" : "No"}
        tooltip={tooltip}
        severity={val ? "success" : "danger"}
      />
    );
  };
}

export function renderTag(config: {
  map: Record<
    string,
    {
      label: string;
      severity: "success" | "danger" | "warning" | "info" | "help" | "none";
    }
  >;
  tooltip?: string;
}) {
  return ({ val }: TableColumnRenderContext) => {
    return (
      <BadgeWithTooltip
        value={config.map[val]?.label || "--"}
        severity={config.map[val]?.severity || "primary"}
        tooltip={config.tooltip}
      />
    );
  };
}
