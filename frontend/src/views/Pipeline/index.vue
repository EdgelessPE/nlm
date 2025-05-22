<template>
  <Tabs :value="tab">
    <TabList>
      <Tab value="nep">Nep</Tab>
      <Tab value="ept">Ept</Tab>
    </TabList>
    <TabPanel value="nep">
      <Nep />
    </TabPanel>
    <TabPanel value="ept">
      <Ept />
    </TabPanel>
  </Tabs>
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
