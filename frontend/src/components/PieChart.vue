<template>
  <div class="pie-chart-container">
    <Pie :data="chartData" :options="chartOptions" />
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Pie } from 'vue-chartjs'
import { Chart as ChartJS, ArcElement, Tooltip, Legend } from 'chart.js'

ChartJS.register(ArcElement, Tooltip, Legend)

const props = defineProps({
  completed: { type: Number, default: 25 },
  onHold: { type: Number, default: 25 },
  inProgress: { type: Number, default: 18 },
  pending: { type: Number, default: 32 }
})

const chartData = computed(() => ({
  labels: ['Completed', 'On Hold', 'In Progress', 'Pending'],
  datasets: [
    {
      data: [props.completed, props.onHold, props.inProgress, props.pending],
      backgroundColor: ['#4a7cff', '#8b5cf6', '#67b7f7', '#f87171'],
      borderColor: ['#4a7cff', '#8b5cf6', '#67b7f7', '#f87171'],
      borderWidth: 0,
      hoverOffset: 8,
      spacing: 2
    }
  ]
}))

const chartOptions = {
  responsive: true,
  maintainAspectRatio: true,
  plugins: {
    legend: {
      position: 'right',
      labels: {
        padding: 16,
        usePointStyle: true,
        pointStyle: 'circle',
        font: {
          family: 'Inter',
          size: 12,
          weight: '500'
        },
        color: '#6b7280'
      }
    },
    tooltip: {
      backgroundColor: '#1e2139',
      titleFont: { family: 'Inter', size: 13, weight: '600' },
      bodyFont: { family: 'Inter', size: 12 },
      padding: 12,
      cornerRadius: 8,
      callbacks: {
        label: function(context) {
          return `${context.label}: ${context.parsed}%`
        }
      }
    }
  },
  animation: {
    animateRotate: true,
    duration: 1000
  }
}
</script>

<style scoped>
.pie-chart-container {
  width: 100%;
  max-width: 320px;
  margin: 0 auto;
}
</style>
