import type { VNodeChild } from "vue";

export interface ActionNode<T = any> {
  key: string;
  label: string;
  icon: () => VNodeChild;
  disabled?: (data: T) => boolean;
  onClick: (data: T, event: Event) => void | Promise<void>;
}
