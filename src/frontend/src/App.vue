<template>
  <div class="container-fluid">
    <ViewSSE
      v-if="inited"
      :inited="inited"
    />
  </div>
</template>

<script lang="ts">
import { ref, defineComponent } from "vue";
import ViewSSE from "@/views/ViewSSE.vue";
export default defineComponent({
  components: {
    ViewSSE
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
});
</script>
