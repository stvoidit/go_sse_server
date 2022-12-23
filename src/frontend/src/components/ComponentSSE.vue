<script lang="ts">
import Chart from "chart.js/auto";
import moment from "moment";
import { ChartOptions } from "chart.js";
import { ref, onMounted, watch, shallowRef, defineComponent } from "vue";
import { useEventsStore } from "@/store";
export default defineComponent({
  props: {
    inited: { type: Boolean, required: true }
  },
  setup() {
    const eventStore = useEventsStore();
    const chart = shallowRef<Chart>();
    const refChart = ref<HTMLCanvasElement>();
    const showLog = ref(false);

    watch(eventStore.events, () => {
      if (!chart.value || eventStore.events.length === 0) return;
      const lastElem = eventStore.events[eventStore.events.length-1];
      chart.value.data.labels?.push(moment(lastElem.data.time).format("YYYY-MM-DD HH:mm:ss"));
      chart.value.data.datasets.forEach(dataset => dataset.data.push(lastElem.data.payload));
      chart.value.update();
    });

    async function renderChart() {
      if (!refChart.value) return;
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
      eventStore,
      refChart,
      showLog
    };
  }
});
</script>

<template>
  <div class="row">
    <div class="col-9">
      <div class="card mt-3">
        <div class="card-body">
          <div>sse status: {{ eventStore.sseStatus }}</div>
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
