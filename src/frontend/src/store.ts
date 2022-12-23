import { defineStore } from "pinia";
import { ref, computed } from "vue";

type ServiceEvent = {
  id: string;
  data: {
    payload: number;
    time: string;
  }
}

export const useEventsStore = defineStore("events", () => {
  const sseStatus = ref(false);
  const events = ref<Array<ServiceEvent>>([]);
  const sse = new EventSource("/api/events", { withCredentials: true });
  sse.addEventListener("message", function(ev){
    events.value.push({
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
      return events.value[events.value.length-1];
    }
    return null;
  });
  return {
    sseStatus,
    events,
    eventCount,
    currentEvent
  };
});
