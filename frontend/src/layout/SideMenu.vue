<template>
  <div class="flex flex-col p-8px">
    <Menu :model="items" class="h-full">
      <template #item="{ item, props }">
        <router-link
          v-if="item.route"
          v-slot="{ href, navigate }"
          :to="item.route"
          custom
        >
          <a
            v-ripple
            :href="href"
            v-bind="props.action"
            @click="navigate"
            :class="curPath === item.route ? 'menu-item__active' : ''"
          >
            <span :class="item.icon" />
            <span class="ml-2">{{ item.label }}</span>
          </a>
        </router-link>
      </template>
    </Menu>
  </div>
</template>

<script setup lang="ts">
import { computed, ref } from "vue";
import type { MenuItem } from "primevue/menuitem";
import { useRoute } from "vue-router";

const items = ref<MenuItem[]>([
  {
    label: "Home",
    icon: "pi pi-home",
    route: "/",
  },
  {
    label: "Nep",
    icon: "pi pi-box",
    route: "/nep",
  },
  {
    label: "Ept",
    icon: "pi pi-hammer",
    route: "/ept",
  },
  {
    label: "Pipeline",
    icon: "pi pi-truck",
    route: "/pipeline",
  },
  {
    label: "Storage",
    icon: "pi pi-cloud",
    route: "/storage",
  },
]);
const route = useRoute();
const curPath = computed(() => route.path);
</script>

<style scoped lang="less">
.menu-item__active {
  @apply rounded-4px;
  background: var(--p-menu-item-focus-background);
}
</style>
