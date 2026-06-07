<template>
  <div class="project-card" @click="$emit('click')">
    <div class="card-header">
      <h3 class="card-title">{{ project.title }}</h3>
      <button class="edit-btn" @click.stop>
        <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <path d="M11 4H4a2 2 0 00-2 2v14a2 2 0 002 2h14a2 2 0 002-2v-7"/>
          <path d="M18.5 2.5a2.121 2.121 0 013 3L12 15l-4 1 1-4 9.5-9.5z"/>
        </svg>
      </button>
    </div>

    <span
      class="status-badge"
      :class="statusClass"
    >
      {{ project.status }}
    </span>

    <p class="card-description">{{ project.description }}</p>

    <div class="card-date">
      <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <rect x="3" y="4" width="18" height="18" rx="2" ry="2"/>
        <line x1="16" y1="2" x2="16" y2="6"/>
        <line x1="8" y1="2" x2="8" y2="6"/>
        <line x1="3" y1="10" x2="21" y2="10"/>
      </svg>
      <span>{{ project.date }}</span>
    </div>

    <div class="card-footer">
      <div class="avatar-group">
        <img
          v-for="(member, idx) in displayMembers"
          :key="idx"
          :src="member.avatar"
          :alt="member.name"
          class="avatar"
          :title="member.name"
        />
        <span v-if="extraMembers > 0" class="avatar-extra">+{{ extraMembers }}</span>
      </div>
      <span class="issues-count">
        <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <circle cx="12" cy="12" r="10"/>
          <line x1="12" y1="8" x2="12" y2="12"/>
          <line x1="12" y1="16" x2="12.01" y2="16"/>
        </svg>
        {{ project.issues }} Issues
      </span>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  project: {
    type: Object,
    required: true
  }
})

defineEmits(['click'])

const statusClass = computed(() => {
  const s = props.project.status?.toLowerCase()
  if (s === 'ontrack') return 'ontrack'
  if (s === 'offline') return 'offline'
  if (s === 'completed') return 'completed'
  return 'offline'
})

const displayMembers = computed(() => {
  return (props.project.members || []).slice(0, 4)
})

const extraMembers = computed(() => {
  const total = (props.project.members || []).length
  return total > 4 ? total - 4 : 0
})
</script>

<style scoped>
.project-card {
  background: var(--white);
  border-radius: var(--radius-lg);
  padding: 22px;
  box-shadow: var(--shadow-card);
  cursor: pointer;
  transition: all var(--transition-normal);
  display: flex;
  flex-direction: column;
  gap: 12px;
  animation: fadeInUp 0.4s ease forwards;
  border: 1px solid transparent;
}

.project-card:hover {
  box-shadow: var(--shadow-lg);
  transform: translateY(-3px);
  border-color: var(--primary-blue-light);
}

.card-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.card-title {
  font-size: 1.05rem;
  font-weight: 700;
  color: var(--text-dark);
}

.edit-btn {
  background: none;
  color: var(--text-muted);
  padding: 4px;
  border-radius: var(--radius-sm);
  transition: all var(--transition-fast);
}

.edit-btn:hover {
  color: var(--primary-blue);
  background: var(--primary-blue-light);
}

.card-description {
  font-size: 0.8125rem;
  color: var(--text-secondary);
  line-height: 1.6;
  display: -webkit-box;
  -webkit-line-clamp: 3;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.card-date {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 0.8125rem;
  color: var(--text-muted);
  font-weight: 500;
}

.card-footer {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-top: 4px;
  padding-top: 14px;
  border-top: 1px solid var(--border-light);
}

.avatar-group .avatar {
  width: 30px;
  height: 30px;
  border-radius: 50%;
  border: 2px solid var(--white);
  margin-left: -8px;
  object-fit: cover;
  transition: transform var(--transition-fast);
}

.avatar-group .avatar:first-child {
  margin-left: 0;
}

.avatar-group .avatar:hover {
  transform: scale(1.15);
  z-index: 2;
}

.avatar-extra {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  width: 30px;
  height: 30px;
  border-radius: 50%;
  background: var(--primary-blue-light);
  color: var(--primary-blue);
  font-size: 0.7rem;
  font-weight: 700;
  margin-left: -8px;
  border: 2px solid var(--white);
}

.issues-count {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 0.8125rem;
  color: var(--text-muted);
  font-weight: 500;
}
</style>
