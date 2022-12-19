<template>
  <div class="container-fluid">
    <ComponentSSE
      v-if="inited"
      :inited="inited"
    />
  </div>
</template>

<script lang="ts">
import ComponentSSE from "@/components/ComponentSSE.vue";
import { ref } from "vue";
export default {
  components: {
    ComponentSSE
  },
  setup() {
    const inited = ref(false);
    fetch("/api/init").then(response => {
      if (response.status === 200) {
        inited.value = true;
      } else {
        alert("not auth");
      }
    });
    return {
      inited
    };
  }
};
</script>
