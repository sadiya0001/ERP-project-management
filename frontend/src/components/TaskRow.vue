<template>
  <div class="task-row" :class="{ checked: isChecked }">
    <div class="task-checkbox">
      <input
        type="checkbox"
        :id="`task-${task.id}`"
        v-model="isChecked"
        class="checkbox"
      />
    </div>

    <div class="task-details">
      <div class="task-main">
        <h4 class="task-title">{{ task.title }}</h4>
        <div class="task-meta">
          <span class="task-id">{{ task.taskId }}</span>
          <span class="task-created">
            Started {{ task.daysAgo }} days ago by
            <strong>{{ task.createdBy }}</strong>
          </span>
        </div>
      </div>
      <span
        class="status-badge"
        :class="statusClass"
      >
        {{ task.status }}
      </span>
    </div>

    <div class="task-time">
      <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <circle cx="12" cy="12" r="10"/>
        <polyline points="12,6 12,12 16,14"/>
      </svg>
      <span>{{ task.timeSpent }}</span>
    </div>

    <div class="task-deadline">
      <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <rect x="3" y="4" width="18" height="18" rx="2" ry="2"/>
        <line x1="16" y1="2" x2="16" y2="6"/>
        <line x1="8" y1="2" x2="8" y2="6"/>
        <line x1="3" y1="10" x2="21" y2="10"/>
      </svg>
      <span>{{ formatDate(task.deadline) }}</span>
    </div>

    <div class="task-assignee">
      <img
        :src="task.assignee.avatar"
        :alt="task.assignee.name"
        class="assignee-avatar"
        :title="task.assignee.name"
      />
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'

const props = defineProps({
  task: {
    type: Object,
    required: true
  }
})

const isChecked = ref(props.task.status === 'Completed')

const statusClass = computed(() => {
  const s = props.task.status
  if (s === 'Completed') return 'completed'
  if (s === 'On-hold') return 'on-hold'
  if (s === 'In Progress') return 'in-progress'
  if (s === 'Pending') return 'pending'
  return ''
})

function formatDate(dateStr) {
  if (!dateStr) return '-'
  const d = new Date(dateStr)
  const options = { day: '2-digit', month: 'short', year: 'numeric' }
  return d.toLocaleDateString('en-GB', options)
}
</script>

<style scoped>
.task-row {
  display: grid;
  grid-template-columns: 40px 1fr 120px 140px 48px;
  align-items: center;
  gap: 16px;
  padding: 16px 20px;
  background: var(--white);
  border-bottom: 1px solid var(--border-light);
  transition: all var(--transition-fast);
  animation: fadeIn 0.3s ease forwards;
}

.task-row:hover {
  background: #fafbff;
}

.task-row.checked {
  opacity: 0.7;
}

.task-checkbox {
  display: flex;
  align-items: center;
  justify-content: center;
}

.checkbox {
  width: 18px;
  height: 18px;
  border: 2px solid var(--border-color);
  border-radius: 4px;
  cursor: pointer;
  accent-color: var(--primary-blue);
}

.task-details {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 16px;
  min-width: 0;
}

.task-main {
  min-width: 0;
}

.task-title {
  font-size: 0.875rem;
  font-weight: 600;
  color: var(--text-dark);
  margin-bottom: 4px;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.task-meta {
  display: flex;
  align-items: center;
  gap: 10px;
  font-size: 0.75rem;
  color: var(--text-muted);
}

.task-id {
  font-weight: 600;
  color: var(--primary-blue);
}

.task-created strong {
  color: var(--text-secondary);
}

.task-time,
.task-deadline {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 0.8125rem;
  color: var(--text-secondary);
  font-weight: 500;
}

.task-assignee {
  display: flex;
  justify-content: center;
}

.assignee-avatar {
  width: 34px;
  height: 34px;
  border-radius: 50%;
  object-fit: cover;
  border: 2px solid var(--primary-blue-light);
  transition: transform var(--transition-fast);
}

.assignee-avatar:hover {
  transform: scale(1.15);
}

@media (max-width: 1024px) {
  .task-row {
    grid-template-columns: 40px 1fr 48px;
  }

  .task-time,
  .task-deadline {
    display: none;
  }
}

@media (max-width: 768px) {
  .task-row {
    grid-template-columns: 32px 1fr 40px;
    padding: 12px 14px;
    gap: 8px;
  }

  .task-meta {
    flex-direction: column;
    align-items: flex-start;
    gap: 2px;
  }
}
</style>
