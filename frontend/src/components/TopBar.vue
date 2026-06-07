<template>
  <header class="topbar">
    <div class="topbar-left">
      <button class="mobile-menu-btn" @click="$emit('toggle-sidebar')">
        <svg width="22" height="22" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round">
          <line x1="3" y1="6" x2="21" y2="6"/>
          <line x1="3" y1="12" x2="21" y2="12"/>
          <line x1="3" y1="18" x2="21" y2="18"/>
        </svg>
      </button>
      <div class="search-box">
        <svg class="search-icon" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round">
          <circle cx="11" cy="11" r="8"/>
          <line x1="21" y1="21" x2="16.65" y2="16.65"/>
        </svg>
        <input
          type="text"
          class="search-input"
          placeholder="Search..."
          v-model="searchQuery"
        />
      </div>
    </div>

    <div class="topbar-right">
      <!-- Notification bell -->
      <button class="icon-btn">
        <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <path d="M18 8A6 6 0 006 8c0 7-3 9-3 9h18s-3-2-3-9"/>
          <path d="M13.73 21a2 2 0 01-3.46 0"/>
        </svg>
        <span class="notification-dot"></span>
      </button>

      <!-- User info -->
      <router-link to="/profile" class="user-info">
        <img
          :src="authStore.currentUser.avatar"
          :alt="authStore.currentUser.firstName"
          class="user-avatar"
        />
        <div class="user-details">
          <span class="user-name">{{ authStore.currentUser.firstName }} {{ authStore.currentUser.lastName }}</span>
          <span class="user-role">{{ authStore.currentUser.role }}</span>
        </div>
        <svg class="chevron" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round">
          <polyline points="6,9 12,15 18,9"/>
        </svg>
      </router-link>
    </div>
  </header>
</template>

<script setup>
import { ref } from 'vue'
import { useAuthStore } from '../stores/auth'

const authStore = useAuthStore()
const searchQuery = ref('')

defineEmits(['toggle-sidebar'])
</script>

<style scoped>
.topbar {
  position: fixed;
  top: 0;
  left: var(--sidebar-width);
  right: 0;
  height: var(--topbar-height);
  background: var(--white);
  border-bottom: 1px solid var(--border-light);
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 32px;
  z-index: 50;
  transition: left var(--transition-normal);
}

.topbar-left {
  display: flex;
  align-items: center;
  gap: 16px;
}

.mobile-menu-btn {
  display: none;
  background: none;
  color: var(--text-secondary);
  padding: 4px;
}

.search-box {
  position: relative;
  display: flex;
  align-items: center;
}

.search-icon {
  position: absolute;
  left: 14px;
  color: var(--text-muted);
  pointer-events: none;
}

.search-input {
  width: 320px;
  padding: 10px 14px 10px 42px;
  border: 1.5px solid var(--border-color);
  border-radius: var(--radius-full);
  font-size: 0.875rem;
  color: var(--text-primary);
  background: var(--main-bg);
  transition: all var(--transition-fast);
}

.search-input:focus {
  border-color: var(--primary-blue);
  background: var(--white);
  box-shadow: 0 0 0 3px rgba(74, 124, 255, 0.1);
  width: 380px;
}

.search-input::placeholder {
  color: var(--text-muted);
}

.topbar-right {
  display: flex;
  align-items: center;
  gap: 20px;
}

.icon-btn {
  position: relative;
  background: none;
  color: var(--text-secondary);
  padding: 8px;
  border-radius: var(--radius-md);
  transition: all var(--transition-fast);
}

.icon-btn:hover {
  background: var(--main-bg);
  color: var(--text-primary);
}

.notification-dot {
  position: absolute;
  top: 6px;
  right: 6px;
  width: 8px;
  height: 8px;
  background: var(--status-onhold);
  border-radius: 50%;
  border: 2px solid var(--white);
}

.user-info {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 6px 12px;
  border-radius: var(--radius-md);
  cursor: pointer;
  transition: background var(--transition-fast);
  text-decoration: none;
}

.user-info:hover {
  background: var(--main-bg);
}

.user-avatar {
  width: 38px;
  height: 38px;
  border-radius: 50%;
  object-fit: cover;
  border: 2px solid var(--primary-blue-light);
}

.user-details {
  display: flex;
  flex-direction: column;
}

.user-name {
  font-size: 0.875rem;
  font-weight: 600;
  color: var(--text-dark);
  line-height: 1.2;
}

.user-role {
  font-size: 0.75rem;
  color: var(--text-muted);
}

.chevron {
  color: var(--text-muted);
}

@media (max-width: 768px) {
  .topbar {
    left: 0;
    padding: 0 16px;
  }

  .mobile-menu-btn {
    display: flex;
  }

  .search-input {
    width: 200px;
  }

  .search-input:focus {
    width: 200px;
  }

  .user-details {
    display: none;
  }

  .chevron {
    display: none;
  }
}
</style>
