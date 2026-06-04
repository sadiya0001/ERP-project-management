<template>
  <div class="tasks-page">
    <div class="breadcrumb">
      <router-link to="/projects">Projects</router-link>
      <span class="separator">/</span>
      <span class="current">Adoddle</span>
    </div>
    
    <div class="project-header">
      <div class="header-left">
        <h2 class="project-title">Adoddle <span class="emoji">👨‍💻</span></h2>
        
        <div class="team-avatars">
          <img src="https://i.pravatar.cc/150?img=11" alt="Team" class="avatar" />
          <img src="https://i.pravatar.cc/150?img=12" alt="Team" class="avatar" />
          <img src="https://i.pravatar.cc/150?img=13" alt="Team" class="avatar" />
          <img src="https://i.pravatar.cc/150?img=14" alt="Team" class="avatar" />
          <div class="avatar-more">+2</div>
        </div>
        
        <div class="status-badge on-track">OnTrack</div>
      </div>
      
      <div class="header-right">
        <button class="assign-btn">Assign Task</button>
        
        <div class="metrics">
          <div class="metric">
            <span class="metric-label">Time spent</span>
            <div class="metric-value">
              <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"></circle><polyline points="12 6 12 12 16 14"></polyline></svg>
              2W : 0W : 0D
            </div>
          </div>
          <div class="metric">
            <span class="metric-label">Deadline</span>
            <div class="metric-value success">
              <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect x="3" y="4" width="18" height="18" rx="2" ry="2"></rect><line x1="16" y1="2" x2="16" y2="6"></line><line x1="8" y1="2" x2="8" y2="6"></line><line x1="3" y1="10" x2="21" y2="10"></line></svg>
              5M : 0W : 0D
            </div>
          </div>
        </div>
      </div>
    </div>
    
    <div class="tasks-list-container">
      <div class="tasks-list">
        <TaskRow 
          v-for="task in tasks" 
          :key="task.id" 
          :task="task" 
        />
      </div>
      
      <div class="list-footer">
        <div class="footer-stats">
          <span class="stat-item">
            <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"></path></svg>
            56 tasks
          </span>
          <span class="stat-item">
            <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"></path><polyline points="14 2 14 8 20 8"></polyline><line x1="16" y1="13" x2="8" y2="13"></line><line x1="16" y1="17" x2="8" y2="17"></line><polyline points="10 9 9 9 8 9"></polyline></svg>
            15 files
          </span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import TaskRow from '../components/TaskRow.vue';

// Mock data
const tasks = ref([
  {
    id: '#NV221',
    title: 'Make an Automatic Payment System that enable the design',
    creator: 'Yash Ghori',
    daysAgo: 14,
    status1: 'Canceled',
    status2: 'Completed',
    timeSpent: '00:30:00',
    assigneeAvatar: 'https://i.pravatar.cc/150?img=11'
  },
  {
    id: '#NV222',
    title: 'Make an Automatic Payment System that enable the design',
    creator: 'Yash Ghori',
    daysAgo: 12,
    status1: 'Canceled',
    status2: 'Completed',
    timeSpent: '00:30:00',
    assigneeAvatar: 'https://i.pravatar.cc/150?img=12'
  },
  {
    id: '#NV223',
    title: 'Make an Automatic Payment System that enable the design',
    creator: 'Yash Ghori',
    daysAgo: 12,
    status1: 'Canceled',
    status2: 'Completed',
    timeSpent: '00:30:00',
    assigneeAvatar: 'https://i.pravatar.cc/150?img=13'
  },
  {
    id: '#NV224',
    title: 'Make an Automatic Payment System that enable the design',
    creator: 'Yash Ghori',
    daysAgo: 10,
    status1: 'Canceled',
    status2: 'Completed',
    timeSpent: '00:30:00',
    assigneeAvatar: 'https://i.pravatar.cc/150?img=14'
  },
  {
    id: '#NV225',
    title: 'Make an Automatic Payment System that enable the design',
    creator: 'Yash Ghori',
    daysAgo: 10,
    status1: 'Canceled',
    status2: 'Completed',
    timeSpent: '00:30:00',
    assigneeAvatar: 'https://i.pravatar.cc/150?img=15'
  },
  {
    id: '#NV226',
    title: 'Make an Automatic Payment System that enable the design',
    creator: 'Yash Ghori',
    daysAgo: 15,
    status1: 'Canceled',
    status2: 'Completed',
    timeSpent: '00:30:00',
    assigneeAvatar: 'https://i.pravatar.cc/150?img=16'
  }
]);
</script>

<style scoped>
.tasks-page {
  padding: 24px;
  max-width: 1200px;
  margin: 0 auto;
}

.breadcrumb {
  font-size: 14px;
  margin-bottom: 24px;
  color: var(--text-gray, #6b7280);
}

.breadcrumb a {
  color: var(--primary-color, #4a90d9);
  text-decoration: none;
}

.separator {
  margin: 0 8px;
}

.current {
  color: var(--text-dark, #1a1d29);
  font-weight: 500;
}

.project-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 32px;
  flex-wrap: wrap;
  gap: 20px;
}

.header-left {
  display: flex;
  align-items: center;
  gap: 16px;
}

.project-title {
  font-size: 24px;
  font-weight: 700;
  color: var(--text-dark, #1a1d29);
  margin: 0;
  display: flex;
  align-items: center;
  gap: 8px;
}

.emoji {
  font-size: 20px;
}

.team-avatars {
  display: flex;
  align-items: center;
}

.avatar {
  width: 32px;
  height: 32px;
  border-radius: 50%;
  border: 2px solid #f5f7fa;
  margin-left: -8px;
}

.avatar:first-child {
  margin-left: 0;
}

.avatar-more {
  width: 32px;
  height: 32px;
  border-radius: 50%;
  background-color: #fee2e2;
  color: #ef4444;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 12px;
  font-weight: 600;
  border: 2px solid #f5f7fa;
  margin-left: -8px;
  z-index: 10;
}

.status-badge {
  padding: 6px 16px;
  border-radius: 20px;
  font-size: 13px;
  font-weight: 500;
}

.on-track {
  background-color: #dcfce7;
  color: #16a34a;
}

.header-right {
  display: flex;
  align-items: center;
  gap: 24px;
}

.assign-btn {
  background-color: var(--primary-color, #4a90d9);
  color: white;
  border: none;
  border-radius: 8px;
  padding: 10px 24px;
  font-size: 14px;
  font-weight: 500;
  cursor: pointer;
}

.metrics {
  display: flex;
  gap: 24px;
}

.metric {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.metric-label {
  font-size: 12px;
  color: var(--text-gray, #6b7280);
}

.metric-value {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 13px;
  font-weight: 600;
  color: var(--text-dark, #1a1d29);
  background-color: #e0f2fe;
  color: #0369a1;
  padding: 4px 10px;
  border-radius: 12px;
}

.metric-value.success {
  background-color: #dcfce7;
  color: #16a34a;
}

.tasks-list-container {
  background-color: var(--card-bg, #ffffff);
  border-radius: 12px;
  padding: 24px;
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.05);
}

.tasks-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
  margin-bottom: 24px;
}

.list-footer {
  display: flex;
  justify-content: flex-end;
  border-top: 1px solid #f3f4f6;
  padding-top: 16px;
}

.footer-stats {
  display: flex;
  gap: 16px;
  color: var(--text-gray, #6b7280);
  font-size: 13px;
}

.stat-item {
  display: flex;
  align-items: center;
  gap: 6px;
}

@media (max-width: 768px) {
  .project-header {
    flex-direction: column;
    align-items: flex-start;
  }
  
  .header-right {
    width: 100%;
    flex-direction: column;
    align-items: stretch;
  }
  
  .metrics {
    justify-content: space-between;
  }
}
</style>
