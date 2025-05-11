import type { VNodeChild } from "vue";

export interface ActionNode<T = any> {
  key: string;
  label: string;
  icon: () => VNodeChild;
  onClick: (data: T) => void;
}
