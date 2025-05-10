<template>
  <div
    class="flex items-center justify-between bg-[var(--p-surface-200)] dark:bg-[var(--p-surface-800)] px-32px"
  >
    <div class="flex items-center gap-16px">
      <img :src="Logo" class="w-32px h-32px" alt="logo" />
      <span>Nep Lifetime Management</span>
    </div>
    <div class="flex items-center gap-8px">
      <Button icon="pi pi-sun" size="small" @click="toggleDark()"> </Button>
      <Button icon="pi pi-user" label="Login" size="small" />
    </div>
  </div>
</template>

<script setup lang="ts">
import Logo from "@/assets/nep.ico";
import { useDark, useToggle } from "@vueuse/core";

const isDark = useDark();
const toggleDark = useToggle(isDark);
</script>
