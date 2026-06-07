<template>
  <div class="create-project-page">
    <div class="breadcrumb">
      <router-link to="/projects">Projects</router-link>
      <span class="separator">/</span>
      <span class="current">Create Project</span>
    </div>
    
    <div class="form-card">
      <form @submit.prevent="createProject">
        <div class="form-row">
          <div class="form-group">
            <label>Project Title</label>
            <input type="text" v-model="form.title" required />
          </div>
          <div class="form-group">
            <label>Project Type</label>
            <input type="text" v-model="form.type" />
          </div>
          <div class="form-group">
            <label>Start Date</label>
            <div class="date-input">
              <input type="date" v-model="form.startDate" />
            </div>
          </div>
          <div class="form-group">
            <label>End Date</label>
            <div class="date-input">
              <input type="date" v-model="form.endDate" />
            </div>
          </div>
        </div>
        
        <div class="form-group full-width">
          <label>Project Description</label>
          <textarea v-model="form.description" rows="4"></textarea>
        </div>
        
        <div class="roles-section">
          <h3>Project Roles</h3>
          
          <div class="role-selector">
            <select v-model="selectedRole">
              <option value="Team Lead">Team Lead</option>
              <option value="Developer">Developer</option>
              <option value="Designer">Designer</option>
            </select>
            <button type="button" class="icon-btn remove-btn">
              <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="5" y1="12" x2="19" y2="12"></line></svg>
            </button>
          </div>
          
          <div class="roles-list">
            <div class="role-row" v-for="(member, index) in rolesList" :key="index">
              <div class="member-name">{{ member.name }}</div>
              <div class="member-role">{{ member.role }}</div>
              <div class="member-check">
                <input type="checkbox" v-model="member.selected" />
              </div>
            </div>
          </div>
        </div>
        
        <div class="form-actions">
          <button type="submit" class="create-btn">Create</button>
          <button type="button" class="delete-btn" @click="$router.push('/projects')">Cancel</button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup>
import { reactive, ref } from 'vue';

const form = reactive({
  title: '',
  type: '',
  startDate: '',
  endDate: '',
  description: ''
});

const selectedRole = ref('Team Lead');

const rolesList = ref([
  { name: 'Yash', role: 'Team lead', selected: true },
  { name: 'Yash', role: 'Team lead', selected: false },
  { name: 'Yash', role: 'Team lead', selected: false },
  { name: 'Yash', role: 'Team lead', selected: false },
  { name: 'Yash', role: 'Team lead', selected: false },
  { name: 'Yash', role: 'Team lead', selected: false }
]);

const createProject = () => {
  alert('Project creation logic here');
};
</script>

<style scoped>
.create-project-page {
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
  color: var(--text-gray, #6b7280);
}

.form-card {
  background-color: var(--card-bg, #ffffff);
  border-radius: 12px;
  padding: 32px;
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.05);
}

.form-row {
  display: grid;
  grid-template-columns: 2fr 2fr 1fr 1fr;
  gap: 24px;
  margin-bottom: 24px;
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.form-group.full-width {
  margin-bottom: 32px;
}

.form-group label {
  font-size: 14px;
  font-weight: 500;
  color: var(--text-dark, #1a1d29);
}

.form-group input, .form-group textarea, .form-group select {
  padding: 12px 16px;
  border: 1px solid #e5e7eb;
  border-radius: 8px;
  font-size: 14px;
  color: var(--text-dark, #1a1d29);
  outline: none;
  transition: border-color 0.2s;
  background-color: #fff;
}

.form-group textarea {
  resize: vertical;
}

.form-group input:focus, .form-group textarea:focus, .form-group select:focus {
  border-color: var(--primary-color, #4a90d9);
}

.roles-section {
  margin-bottom: 32px;
}

.roles-section h3 {
  font-size: 16px;
  font-weight: 600;
  color: var(--text-dark, #1a1d29);
  margin: 0 0 16px 0;
}

.role-selector {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 16px;
  max-width: 300px;
}

.role-selector select {
  flex: 1;
  padding: 10px 16px;
  border: 1px solid #e5e7eb;
  border-radius: 8px;
  background-color: #fff;
  outline: none;
}

.icon-btn {
  background: transparent;
  border: none;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 8px;
  border-radius: 4px;
  color: var(--text-gray, #6b7280);
}

.icon-btn:hover {
  background-color: #f3f4f6;
}

.roles-list {
  border: 1px solid #e5e7eb;
  border-radius: 8px;
  max-width: 400px;
  overflow: hidden;
}

.role-row {
  display: grid;
  grid-template-columns: 1fr 1fr auto;
  padding: 12px 16px;
  border-bottom: 1px solid #e5e7eb;
  align-items: center;
  font-size: 14px;
}

.role-row:last-child {
  border-bottom: none;
}

.member-name {
  color: var(--text-dark, #1a1d29);
}

.member-role {
  color: var(--text-gray, #6b7280);
}

.form-actions {
  display: flex;
  justify-content: flex-end;
  gap: 16px;
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

.delete-btn {
  background-color: #eef2ff;
  color: var(--primary-color, #4a90d9);
  border: none;
  border-radius: 8px;
  padding: 10px 24px;
  font-size: 14px;
  font-weight: 500;
  cursor: pointer;
}

@media (max-width: 1024px) {
  .form-row {
    grid-template-columns: 1fr 1fr;
  }
}

@media (max-width: 640px) {
  .form-row {
    grid-template-columns: 1fr;
  }
}
</style>
