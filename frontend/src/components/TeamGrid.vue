<template>
  <div class="team-grid">
    <div class="grid-header">
      <h3 class="grid-title">UI Developers (39)</h3>
      <div class="header-actions">
        <div class="search-box">
          <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="search-icon"><circle cx="11" cy="11" r="8"></circle><line x1="21" y1="21" x2="16.65" y2="16.65"></line></svg>
          <input type="text" placeholder="Search for anything..." />
        </div>
        <a href="#" class="view-all">View all</a>
      </div>
    </div>
    
    <div class="avatars-container">
      <div v-for="(member, index) in members" :key="index" class="avatar-wrapper" :style="{ borderColor: member.color }">
        <img :src="member.avatar" :alt="member.name" class="avatar-img" />
      </div>
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
import { ref } from 'vue';

const colors = ['#f59e0b', '#ef4444', '#10b981', '#3b82f6', '#8b5cf6', '#ec4899', '#14b8a6'];

// Generate mock team members
const members = ref(Array.from({ length: 18 }, (_, i) => ({
  id: i,
  name: `Developer ${i + 1}`,
  avatar: `https://i.pravatar.cc/150?img=${i + 20}`,
  color: colors[i % colors.length]
})));
</script>

<style scoped>
.team-grid {
  background-color: var(--card-bg, #ffffff);
  border-radius: 12px;
  padding: 24px;
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.05);
}

.grid-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
}

.grid-title {
  font-size: 16px;
  font-weight: 600;
  color: var(--text-dark, #1a1d29);
  margin: 0;
}

.header-actions {
  display: flex;
  align-items: center;
  gap: 16px;
}

.search-box {
  display: flex;
  align-items: center;
  background-color: var(--bg-color, #f5f7fa);
  border-radius: 20px;
  padding: 6px 16px;
  width: 250px;
}

.search-icon {
  color: var(--text-gray, #6b7280);
  margin-right: 8px;
}

.search-box input {
  border: none;
  background: transparent;
  outline: none;
  font-size: 13px;
  width: 100%;
  color: var(--text-dark, #1a1d29);
}

.view-all {
  color: var(--primary-color, #4a90d9);
  font-size: 14px;
  font-weight: 500;
  text-decoration: none;
}

.avatars-container {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(60px, 1fr));
  gap: 20px;
  justify-items: center;
  margin-bottom: 30px;
}

.avatar-wrapper {
  width: 56px;
  height: 56px;
  border-radius: 50%;
  border: 3px solid;
  padding: 2px;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: transform 0.2s;
}

.avatar-wrapper:hover {
  transform: translateY(-2px) scale(1.05);
}

.avatar-img {
  width: 100%;
  height: 100%;
  border-radius: 50%;
  object-fit: cover;
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
  padding: 4px 8px;
  border-radius: 4px;
}

.page-btn.active {
  background-color: var(--primary-color, #4a90d9);
  color: white;
}

.page-btn:hover:not(.active) {
  background-color: rgba(0, 0, 0, 0.05);
}
</style>
