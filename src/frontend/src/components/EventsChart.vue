<template>
  <div class="card mt-3">
    <div class="card-body">
      <canvas ref="refChart" />
    </div>
  </div>
</template>

<script lang="ts">
import Chart from "chart.js/auto";
import moment from "moment";
import { ChartOptions } from "chart.js";
import { defineComponent, onMounted, watch, shallowRef, PropType, ref } from "vue";
interface ServiceEvent {
  id: string;
  data: {
    payload: number;
    time: string;
  }
}
export default defineComponent({
  props: {
    events: { type: Array as PropType<Array<ServiceEvent>>, required: true }
  },
  setup(props) {
    const chart = shallowRef<Chart>();
    const refChart = ref<HTMLCanvasElement>();
    watch(props.events, () => {
      if (!chart.value || props.events.length === 0) return;
      const lastElem = props.events[props.events.length-1];
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
      refChart,
      chart
    };
  }
});
</script>
