<template>
  <div class="card mt-3">
    <div class="card-body">
      <h2>events</h2>
      <h6>current event:</h6>
      <pre>{{ currentEvent }}</pre>
      <h6>events count: {{ eventCount }}</h6>
      <div>
        <button
          type="button"
          class="btn btn-primary"
          @click="toggleShow"
        >
          Show log
        </button>
      </div>
      <ul
        v-if="showLog"
        class="mt-2"
      >
        <li
          v-for="e in events"
          :key="e.id"
        >
          <pre>{{ e }}</pre>
        </li>
      </ul>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, PropType, computed } from "vue";
import { ServiceEvent } from "@/store";
export default defineComponent({
  props: {
    events: { type: Array as PropType<Array<ServiceEvent>>, required: true }
  },
  setup (props) {
    const showLog = ref(false);
    const toggleShow = () => showLog.value = !showLog.value;
    const eventCount = computed(() => props.events.length);
    const currentEvent = computed(() => {
      if (eventCount.value > 0) {
        return props.events[props.events.length-1];
      }
      return null;
    });
    return {
      showLog,
      toggleShow,
      eventCount,
      currentEvent
    };
  }
});
</script>
