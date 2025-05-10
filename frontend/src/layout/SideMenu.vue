<template>
  <div class="flex flex-col p-8px">
    <Menu :model="items">
      <template #item="{ item, props }">
        <router-link
          v-if="item.route"
          v-slot="{ href, navigate }"
          :to="item.route"
          custom
        >
          <a v-ripple :href="href" v-bind="props.action" @click="navigate">
            <span :class="item.icon" />
            <span class="ml-2">{{ item.label }}</span>
          </a>
        </router-link>
      </template>
    </Menu>
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";

const items = ref<
  {
    label: string;
    icon: string;
    route: string;
  }[]
>([
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
</script>
