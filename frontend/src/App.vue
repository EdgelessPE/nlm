<template>
  <div class="layout">
    <Top class="top" />
    <SideMenu class="side-menu" />
    <router-view class="main" />
  </div>
</template>

<script setup lang="ts">
import Top from "./layout/Top.vue";
import SideMenu from "./layout/SideMenu.vue";
</script>

<style scoped lang="less">
.layout {
  @apply h-full grid gap-x-8px gap-y-4px;

  grid-template-columns: 320px 1fr;
  grid-template-rows: 64px 1fr;
  grid-template-areas:
    "top top"
    "side-menu main";

  .top {
    grid-area: top;
  }
  .side-menu {
    grid-area: side-menu;
  }
  .main {
    grid-area: main;
  }
}
</style>
