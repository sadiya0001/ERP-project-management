<template>
  <div class="ring-chart-container">
    <div class="chart-wrapper">
      <Doughnut :data="chartData" :options="chartOptions" />
      <div class="center-text">
        <span class="time">{{ time }}</span>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue';
import { Chart as ChartJS, ArcElement, Tooltip, Legend } from 'chart.js';
import { Doughnut } from 'vue-chartjs';

ChartJS.register(ArcElement, Tooltip, Legend);

const props = defineProps({
  time: {
    type: String,
    default: '5w: 2d'
  },
  percentage: {
    type: Number,
    default: 75
  }
});

const chartData = computed(() => ({
  datasets: [
    {
      backgroundColor: ['#4a90d9', '#e5e7eb'],
      data: [props.percentage, 100 - props.percentage],
      borderWidth: 0,
      cutout: '80%'
    }
  ]
}));

const chartOptions = {
  responsive: true,
  maintainAspectRatio: false,
  plugins: {
    legend: {
      display: false
    },
    tooltip: {
      enabled: false
    }
  }
};
</script>

<style scoped>
.ring-chart-container {
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 20px;
}

.chart-wrapper {
  position: relative;
  width: 150px;
  height: 150px;
}

.center-text {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  display: flex;
  align-items: center;
  justify-content: center;
}

.time {
  font-size: 24px;
  font-weight: 700;
  color: var(--text-dark, #1a1d29);
}
</style>
