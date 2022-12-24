<script lang="ts">
import { ref, defineComponent } from "vue";
import { useEventsStore } from "@/store";
import EventsChart from "@/components/EventsChart.vue";
import ServiceStatuses from "@/components/ServiceStatuses.vue";
export default defineComponent({
  components: {
    EventsChart,
    ServiceStatuses
  },
  props: {
    inited: { type: Boolean, required: true }
  },
  setup() {
    const eventStore = useEventsStore();
    const showLog = ref(false);
    return {
      eventStore,
      showLog
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
      <EventsChart :events="eventStore.events" />
    </div>
    <div class="col-3">
      <div class="card mt-3">
        <div class="card-body">
          <h2>events</h2>
          <h6>current event:</h6>
          <pre>{{ eventStore.currentEvent }}</pre>
          <h6>events count: {{ eventStore.eventCount }}</h6>
          <div>
            <button
              type="button"
              class="btn btn-primary"
              @click="() => showLog = !showLog"
            >
              Show log
            </button>
          </div>
          <ul
            v-if="showLog"
            class="mt-2"
          >
            <li
              v-for="e in eventStore.events"
              :key="e.id"
            >
              <pre>{{ e }}</pre>
            </li>
          </ul>
        </div>
      </div>
    </div>
  </div>
</template>
