<script lang="ts">
import { ref, computed, onMounted, watch, shallowRef } from "vue";

import Chart from "chart.js/auto";
import { ChartOptions } from "chart.js";
import moment from "moment";
interface ServiceEvent {
  id: string;
  data: {
    payload: number;
    time: string;
  }
}
export default {
  props: {
    inited: { type: Boolean, required: true }
  },
  setup() {
    const chart = shallowRef<Chart>();
    const refChart = ref();
    const sseStatus = ref(false);
    const showLog = ref(false);
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

    watch(events.value, () => {
      if (!chart.value || events.value.length === 0) return;
      const lastElem = events.value[events.value.length-1];
      chart.value.data.labels?.push(moment(lastElem.data.time).format("YYYY-MM-DD HH:mm:ss"));
      chart.value.data.datasets.forEach(dataset => dataset.data.push(lastElem.data.payload));
      chart.value.update();
    });


    function renderChart() {
      const chartOptions: ChartOptions = {
        responsive: true,
        interaction: {
          intersect: false,
          mode: "index"
        },
        plugins: {
          title: {
            display: true,
            text: "random number"
          },
          legend: {
            display: false
          }
        }
      };
      chart.value = new Chart(
        refChart.value,
        {
          type: "bar",
          data: {
            labels: [],
            datasets: [
              {
                backgroundColor: "#850391",
                data: []
              }
            ]
          },
          options: chartOptions
        }
      );
    }
    onMounted(renderChart);
    return {
      events,
      eventCount,
      currentEvent,
      sseStatus,
      refChart,
      showLog
    };
  }
};
</script>

<template>
  <div class="row">
    <div class="col-9">
      <div class="card mt-3">
        <div class="card-body">
          <div>sse status: {{ sseStatus }}</div>
          <div>auth: {{ inited }}</div>
        </div>
      </div>
      <div class="card mt-3">
        <div class="card-body">
          <canvas ref="refChart" />
        </div>
      </div>
    </div>
    <div class="col-3">
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
              v-for="e in events"
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
