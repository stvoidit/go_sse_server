<script lang="ts">
import { ref, computed, onBeforeUnmount } from 'vue'
export default {
  setup() {
    const events = ref<Array<any>>([])
    const sse = new EventSource("/api/events", { withCredentials: true })
    sse.addEventListener("message", function(ev){
      events.value.unshift({
        id: ev.lastEventId,
        ...JSON.parse(ev.data)
      })
    })
    sse.addEventListener("open", function(ev) {console.log(ev)})
    sse.onerror =(ev) => console.error(ev)
    const eventCount = computed(() => events.value.length)
    const currentEvent = computed(() => {
      if (eventCount.value > 0) {
        return events.value[0]
      }
      return null
    })

    return {
      events,
      eventCount,
      currentEvent
    }
  }
}
</script>

<template>
  <h2>events</h2>
  <div class="container">
    <div>current event: {{currentEvent}}</div>
    <div>events count: {{eventCount}}</div>
    <div>
      <div v-for="e in events" :key="e.id">{{e}}</div>
    </div>
  </div>
</template>

<style scoped>

</style>
