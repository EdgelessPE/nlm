<template>
  <div
    class="flex items-center justify-between bg-[var(--p-surface-300)] dark:bg-[var(--p-surface-800)] px-32px"
  >
    <div class="flex items-center gap-16px">
      <img :src="Logo" class="w-32px h-32px" alt="logo" />
      <span>Nep Lifetime Management</span>
    </div>
    <div class="flex items-center gap-8px">
      <Avatar
        icon="pi pi-user"
        class="cursor-pointer"
        shape="circle"
        @click="toggleMenu"
      />
      <TieredMenu ref="menu" id="overlay_menu" :model="items" :popup="true" />
    </div>
  </div>
</template>

<script setup lang="ts">
import Logo from "@/assets/nep.ico";
import { useDark, useLocalStorage } from "@vueuse/core";
import { computed, ref } from "vue";
import type { MenuItem } from "primevue/menuitem";

const darkStorage = useLocalStorage("vueuse-color-scheme", "auto");
useDark();

const menu = ref();
const items = computed<MenuItem[]>(() => [
  {
    label: "Login",
    icon: "pi pi-sign-in",
  },
  {
    label: "Language",
    icon: "pi pi-globe",
    items: [
      {
        label: "简体中文",
        icon: "icon-zh",
      },
      {
        label: "English",
        icon: "pi pi-us",
      },
    ],
  },
  {
    label: "Color Scheme",
    icon: "pi pi-palette",
    items: [
      {
        label: "System",
        icon: darkStorage.value === "auto" ? "pi pi-check" : "pi pi-desktop",
        disabled: darkStorage.value === "auto",
        command: () => {
          darkStorage.value = "auto";
        },
      },
      {
        label: "Light",
        icon: darkStorage.value === "light" ? "pi pi-check" : "pi pi-sun",
        disabled: darkStorage.value === "light",
        command: () => {
          darkStorage.value = "light";
        },
      },
      {
        label: "Dark",
        icon: darkStorage.value === "dark" ? "pi pi-check" : "pi pi-moon",
        disabled: darkStorage.value === "dark",
        command: () => {
          darkStorage.value = "dark";
        },
      },
    ],
  },
]);

const toggleMenu = (event: Event) => {
  menu.value?.toggle(event);
};
</script>

<style scoped lang="less"></style>
