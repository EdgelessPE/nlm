<template>
  <div ref="container" class="max-h-full overflow-hidden">
    <DataTable
      :value="data.value"
      paginator
      :loading="loading.value"
      :rows="pagination.value.limit"
      :rowsPerPageOptions="[10, 20, 50, 100]"
      lazy
      :totalRecords="total.value"
      v-bind="tableProps"
      scrollable
      :scrollHeight="tableHeight"
      paginatorTemplate="CurrentPageReport FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink JumpToPageInput RowsPerPageDropdown"
      currentPageReportTemplate="Total {totalRecords}"
      @page="onPageChange"
    >
      <Column
        v-for="column in columns"
        :key="column.field"
        :field="column.field"
        :header="column.label"
      >
        <template #body="{ data }">
          <component
            :is="column.render"
            :val="data[column.field]"
            :data="data"
          />
        </template>
      </Column>
      <!-- vue-ignore -->
      <template v-for="(_, name) in $slots" :key="name" #[name]="data">
        <slot :name="name" v-bind="data ?? {}" />
      </template>
    </DataTable>
  </div>
</template>

<script setup lang="ts">
import { computed, ref } from "vue";
import { useElementSize } from "@vueuse/core";
import type { UseTableDataReturn } from "./type";

const props = defineProps<UseTableDataReturn>();

const container = ref<HTMLDivElement>();

const { height } = useElementSize(container);
const tableHeight = computed(() => {
  const n = height.value - (props.tableProps?.slotHeight ?? 0) - 58;
  return `${n}px`;
});
</script>
