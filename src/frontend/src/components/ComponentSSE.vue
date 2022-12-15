<script lang="ts">
import { ref, computed } from "vue";
interface ServiceEvent {
  id: string;
  data: {
    payload: string;
    time: string;
  }
}
export default {
  props: {
    inited: { type: Boolean, required: true }
  },
  setup() {
    const sseStatus = ref(false);
    const events = ref<Array<ServiceEvent>>([]);
    const sse = new EventSource("/api/events", { withCredentials: true });
    sse.addEventListener("message", function(ev){
      events.value.unshift({
        id: ev.lastEventId,
        data: JSON.parse(ev.data)
      });
    });
    sse.addEventListener("open", function() { sseStatus.value = true; });
    sse.onerror = (ev) => {
      sseStatus.value = false;
      console.error(ev);
    };
    const eventCount = computed(() => events.value.length);
    const currentEvent = computed(() => {
      if (eventCount.value > 0) {
        return events.value[0];
      }
      return null;
    });
    return {
      events,
      eventCount,
      currentEvent,
      sseStatus
    };
  }
};
</script>

<template>
  <div class="row">
    <div class="card mt-3">
      <div class="card-body">
        <div>sse status: {{ sseStatus }}</div>
        <div>auth: {{ inited }}</div>
      </div>
    </div>
    <div class="card mt-3">
      <div class="card-body">
        <h2>events</h2>
        <div>current event: {{ currentEvent }}</div>
        <div>events count: {{ eventCount }}</div>
        <ul class="mt-2">
          <li
            v-for="e in events"
            :key="e.id"
          >
            {{ e }}
          </li>
        </ul>
      </div>
    </div>
  </div>
</template>
