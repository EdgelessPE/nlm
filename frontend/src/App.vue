<template>
  <div class="layout">
    <Top class="top" />
    <SideMenu class="side-menu" />
    <Content class="content" />
  </div>
</template>

<script setup lang="ts">
import Top from "./layout/Top.vue";
import SideMenu from "./layout/SideMenu.vue";
import Content from "./layout/Content.vue";
</script>

<style scoped lang="less">
.layout {
  @apply w-full h-full grid gap-x-8px gap-y-4px;

  grid-template-columns: 280px 1fr;
  grid-template-rows: 64px 1fr;
  grid-template-areas:
    "top top"
    "side-menu content";

  .top {
    grid-area: top;
  }
  .side-menu {
    grid-area: side-menu;
  }
  .content {
    grid-area: content;
    overflow-y: hidden;
  }
}
</style>
