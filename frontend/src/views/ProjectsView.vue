<template>
  <div class="projects-page">
    <div class="page-header">
      <h2>Projects</h2>
      <button class="create-btn" @click="$router.push('/projects/create')">Create</button>
    </div>
    
    <div class="projects-grid">
      <ProjectCard 
        v-for="project in projects" 
        :key="project.id" 
        :project="project" 
      />
    </div>
    
    <div class="pagination">
      <button class="page-btn prev">Previous</button>
      <button class="page-btn active">1</button>
      <button class="page-btn">2</button>
      <button class="page-btn">3</button>
      <button class="page-btn next">Next</button>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import ProjectCard from '../components/ProjectCard.vue';

// Mock data until API is connected
const projects = ref([
  {
    id: 1,
    title: 'Adoddle',
    description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.',
    date: '05 APRIL 2023',
    status: 'OffTrack',
    statusColor: '#ef4444',
    statusBg: '#fee2e2',
    issues: 14,
    team: [
      { id: 1, avatar: 'https://i.pravatar.cc/150?img=11' },
      { id: 2, avatar: 'https://i.pravatar.cc/150?img=12' },
      { id: 3, avatar: 'https://i.pravatar.cc/150?img=13' },
      { id: 4, avatar: 'https://i.pravatar.cc/150?img=14' }
    ]
  },
  {
    id: 2,
    title: 'Adoddle',
    description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.',
    date: '05 APRIL 2023',
    status: 'OffTrack',
    statusColor: '#ef4444',
    statusBg: '#fee2e2',
    issues: 14,
    team: [
      { id: 1, avatar: 'https://i.pravatar.cc/150?img=21' },
      { id: 2, avatar: 'https://i.pravatar.cc/150?img=22' },
      { id: 3, avatar: 'https://i.pravatar.cc/150?img=23' },
      { id: 4, avatar: 'https://i.pravatar.cc/150?img=24' }
    ]
  },
  {
    id: 3,
    title: 'Adoddle',
    description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.',
    date: '05 APRIL 2023',
    status: 'OffTrack',
    statusColor: '#ef4444',
    statusBg: '#fee2e2',
    issues: 14,
    team: [
      { id: 1, avatar: 'https://i.pravatar.cc/150?img=31' },
      { id: 2, avatar: 'https://i.pravatar.cc/150?img=32' },
      { id: 3, avatar: 'https://i.pravatar.cc/150?img=33' },
      { id: 4, avatar: 'https://i.pravatar.cc/150?img=34' }
    ]
  },
  {
    id: 4,
    title: 'Adoddle',
    description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.',
    date: '05 APRIL 2023',
    status: 'OffTrack',
    statusColor: '#ef4444',
    statusBg: '#fee2e2',
    issues: 14,
    team: [
      { id: 1, avatar: 'https://i.pravatar.cc/150?img=41' },
      { id: 2, avatar: 'https://i.pravatar.cc/150?img=42' },
      { id: 3, avatar: 'https://i.pravatar.cc/150?img=43' },
      { id: 4, avatar: 'https://i.pravatar.cc/150?img=44' }
    ]
  },
  {
    id: 5,
    title: 'Adoddle',
    description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.',
    date: '05 APRIL 2023',
    status: 'OffTrack',
    statusColor: '#ef4444',
    statusBg: '#fee2e2',
    issues: 14,
    team: [
      { id: 1, avatar: 'https://i.pravatar.cc/150?img=51' },
      { id: 2, avatar: 'https://i.pravatar.cc/150?img=52' },
      { id: 3, avatar: 'https://i.pravatar.cc/150?img=53' },
      { id: 4, avatar: 'https://i.pravatar.cc/150?img=54' }
    ]
  },
  {
    id: 6,
    title: 'Adoddle',
    description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.',
    date: '05 APRIL 2023',
    status: 'OffTrack',
    statusColor: '#ef4444',
    statusBg: '#fee2e2',
    issues: 14,
    team: [
      { id: 1, avatar: 'https://i.pravatar.cc/150?img=61' },
      { id: 2, avatar: 'https://i.pravatar.cc/150?img=62' },
      { id: 3, avatar: 'https://i.pravatar.cc/150?img=63' },
      { id: 4, avatar: 'https://i.pravatar.cc/150?img=64' }
    ]
  }
]);
</script>

<style scoped>
.projects-page {
  padding: 24px;
  max-width: 1200px;
  margin: 0 auto;
}

.page-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
}

.page-header h2 {
  font-size: 24px;
  font-weight: 700;
  color: var(--text-dark, #1a1d29);
  margin: 0;
}

.create-btn {
  background-color: var(--primary-color, #4a90d9);
  color: white;
  border: none;
  border-radius: 8px;
  padding: 10px 24px;
  font-size: 14px;
  font-weight: 500;
  cursor: pointer;
  transition: background-color 0.2s;
}

.create-btn:hover {
  background-color: #3b76b8;
}

.projects-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 24px;
  margin-bottom: 32px;
}

.pagination {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 8px;
}

.page-btn {
  background: transparent;
  border: none;
  color: var(--text-gray, #6b7280);
  font-size: 14px;
  font-weight: 500;
  cursor: pointer;
  padding: 6px 12px;
  border-radius: 6px;
}

.page-btn.active {
  background-color: var(--primary-color, #4a90d9);
  color: white;
}

.page-btn:hover:not(.active) {
  background-color: rgba(0, 0, 0, 0.05);
}

@media (max-width: 1024px) {
  .projects-grid {
    grid-template-columns: repeat(2, 1fr);
  }
}

@media (max-width: 768px) {
  .projects-grid {
    grid-template-columns: 1fr;
  }
}
</style>
