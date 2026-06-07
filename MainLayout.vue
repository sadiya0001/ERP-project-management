<template>
  <div class="main-layout">
    <Sidebar />
    <div class="main-content-area">
      <TopBar />
      <main class="page-content">
        <transition name="page" mode="out-in">
          <router-view />
        </transition>
      </main>
    </div>
  </div>
</template>

<script setup>
import Sidebar from '../components/Sidebar.vue'
import TopBar from '../components/TopBar.vue'
</script>

<style scoped>
.main-layout {
  display: flex;
  min-height: 100vh;
  background: var(--main-bg);
}

.main-content-area {
  flex: 1;
  margin-left: var(--sidebar-width);
  display: flex;
  flex-direction: column;
  min-height: 100vh;
  transition: margin-left var(--transition-normal);
}

.page-content {
  flex: 1;
  padding: 28px 32px;
  margin-top: var(--topbar-height);
  overflow-y: auto;
}

@media (max-width: 768px) {
  .main-content-area {
    margin-left: 0;
  }

  .page-content {
    padding: 20px 16px;
  }
}
</style>
