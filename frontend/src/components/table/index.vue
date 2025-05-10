<template>
  <DataTable
    :value="data.value"
    paginator
    :loading="loading.value"
    :rows="pagination.value.limit"
    :rowsPerPageOptions="[10, 20, 50, 100]"
    lazy
    :totalRecords="total.value"
    scrollable
    @page="onPageChange"
  >
    <Column
      v-for="column in columns"
      :key="column.field"
      :field="column.field"
      :header="column.label"
    >
      <template #body="{ data }">
        <component :is="column.render" :val="data[column.field]" :data="data" />
      </template>
    </Column>
  </DataTable>
</template>

<script setup lang="ts">
import type { UseTableDataReturn } from "./type";

defineProps<UseTableDataReturn>();
</script>
