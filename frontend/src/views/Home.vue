<template>
  <div class="flex flex-col items-center justify-center h-full">
    <div class="mb-15vh flex flex-col gap-16px items-center w-60%">
      <h1 class="text-4xl font-bold">Get Some Real Packages</h1>
      <IconField class="w-full">
        <InputIcon class="pi pi-search" />
        <InputText
          v-model="q"
          placeholder="Search Nep Packages"
          class="w-full"
        />
      </IconField>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";

const q = ref("");
</script>
