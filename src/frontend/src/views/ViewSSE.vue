<script lang="ts">
import { defineComponent } from "vue";
import { useEventsStore } from "@/store";
import EventsChart from "@/components/EventsChart.vue";
import EventsLog from "@/components/EventsLog.vue";
import ServiceStatuses from "@/components/ServiceStatuses.vue";
export default defineComponent({
  components: {
    EventsChart,
    EventsLog,
    ServiceStatuses
  },
  props: {
    inited: { type: Boolean, required: true }
  },
  setup() {
    const eventStore = useEventsStore();
    return {
      eventStore
    };
  }
});
</script>

<template>
  <div class="row">
    <div class="col-9">
      <ServiceStatuses
        :inited="inited"
        :sse-status="eventStore.sseStatus"
      />
      <EventsChart
        :events="eventStore.events"
        :type-chart="'bar'"
      />
    </div>
    <div class="col-3">
      <EventsLog :events="eventStore.events" />
    </div>
  </div>
</template>
