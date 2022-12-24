import { defineStore } from "pinia";
import { ref } from "vue";

export type ServiceEvent = {
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
  return {
    sseStatus,
    events
  };
});
