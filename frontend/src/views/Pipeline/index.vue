<template>
  <div class="container" style="background: var(--p-tabs-tablist-background)">
    <Tabs class="w-full h-full" :value="tab">
      <TabList>
        <Tab value="nep">Nep</Tab>
        <Tab value="ept">Ept</Tab>
      </TabList>
      <TabPanel class="panel" value="nep">
        <Nep />
      </TabPanel>
      <TabPanel class="panel" value="ept">
        <Ept />
      </TabPanel>
    </Tabs>
  </div>
</template>

<script setup lang="tsx">
import Nep from "./nep.vue";
import Ept from "./ept.vue";
import { ref } from "vue";

import Tabs from "primevue/tabs";
import TabList from "primevue/tablist";
import Tab from "primevue/tab";
import TabPanel from "primevue/tabpanel";

const tab = ref("nep");
</script>

<style lang="less">
.panel {
  @apply flex flex-col overflow-hidden;
  height: calc(100% - 32px);
}
</style>
