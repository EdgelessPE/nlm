import type { Ref, VNode } from "vue";

export interface FilterNode<T> {
  field: Exclude<keyof T, number | symbol>;
  component: () => VNode;
}

export interface FilterProps<T> {
  model: Ref<T>;
  getConfig: (form: Ref<T>) => FilterNode<T>[];
}
