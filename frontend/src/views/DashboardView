<template>
  <div class="dashboard">
    <div class="page-header">
      <h2>Dashboard</h2>
      <div class="header-icon">
        <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect x="3" y="3" width="7" height="7"></rect><rect x="14" y="3" width="7" height="7"></rect><rect x="14" y="14" width="7" height="7"></rect><rect x="3" y="14" width="7" height="7"></rect></svg>
      </div>
    </div>
    
    <div class="top-cards">
      <div class="card projects-card">
        <div class="card-header">
          <h3>Projects</h3>
          <div class="file-count">
            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M22 19a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h5l2 3h9a2 2 0 0 1 2 2z"></path></svg>
            <span>12 Files</span>
          </div>
        </div>
        
        <div class="projects-preview">
          <!-- Placeholder for project preview images from template -->
          <div class="preview-item">
             <div class="mock-app-ui">
               <div class="app-header">My Plants</div>
               <div class="app-content plants-bg"></div>
             </div>
          </div>
          <div class="preview-item center-item">
             <div class="mock-app-ui dark">
               <div class="app-image flowers-bg"></div>
             </div>
          </div>
          <div class="preview-item">
             <div class="mock-app-ui">
               <div class="app-image papa-bg"></div>
               <div class="app-details">
                 <h4>Papaver Somniferum</h4>
                 <p>Description lorem ipsum</p>
               </div>
             </div>
          </div>
        </div>
      </div>
      
      <div class="card tasks-card">
        <div class="card-header">
          <h3>Tasks</h3>
          <div class="dropdown-badge">
            This Week
            <svg xmlns="http://www.w3.org/2000/svg" width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="6 9 12 15 18 9"></polyline></svg>
          </div>
        </div>
        
        <div class="chart-section">
          <PieChart />
          <div class="chart-legend">
            <div class="legend-item"><span class="dot bg-blue"></span> Completed</div>
            <div class="legend-item"><span class="dot bg-purple"></span> On Hold</div>
            <div class="legend-item"><span class="dot bg-lightblue"></span> On Progress</div>
            <div class="legend-item"><span class="dot bg-red"></span> Pending</div>
          </div>
        </div>
      </div>
    </div>
    
    <TeamGrid />
  </div>
</template>

<script setup>
import PieChart from '../components/PieChart.vue';
import TeamGrid from '../components/TeamGrid.vue';
</script>

<style scoped>
.dashboard {
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

.header-icon {
  color: var(--text-gray, #6b7280);
}

.top-cards {
  display: grid;
  grid-template-columns: 1.2fr 1fr;
  gap: 24px;
  margin-bottom: 24px;
}

.card {
  background-color: var(--card-bg, #ffffff);
  border-radius: 12px;
  padding: 24px;
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.05);
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
}

.card-header h3 {
  font-size: 18px;
  font-weight: 600;
  margin: 0;
  color: var(--text-dark, #1a1d29);
}

.file-count {
  display: flex;
  align-items: center;
  gap: 8px;
  color: var(--text-gray, #6b7280);
  font-size: 14px;
}

.dropdown-badge {
  display: flex;
  align-items: center;
  gap: 6px;
  background-color: #e0f2fe;
  color: #0369a1;
  padding: 6px 12px;
  border-radius: 16px;
  font-size: 12px;
  font-weight: 500;
  cursor: pointer;
}

/* Projects Preview Styling */
.projects-preview {
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 250px;
  gap: 16px;
  overflow: hidden;
}

.preview-item {
  flex: 1;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
}

.mock-app-ui {
  width: 140px;
  height: 240px;
  border-radius: 20px;
  background-color: #f8fafc;
  box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1);
  overflow: hidden;
  display: flex;
  flex-direction: column;
}

.mock-app-ui.dark {
  background-color: #1e293b;
}

.center-item .mock-app-ui {
  width: 160px;
  height: 260px;
  z-index: 10;
  box-shadow: 0 20px 25px -5px rgba(0, 0, 0, 0.1);
}

.app-header {
  padding: 12px;
  font-size: 12px;
  font-weight: 600;
  text-align: center;
  border-bottom: 1px solid #e2e8f0;
}

.app-content {
  flex: 1;
  background-color: #e2e8f0;
  margin: 8px;
  border-radius: 12px;
}

.app-image {
  flex: 1;
  background-size: cover;
  background-position: center;
}

.plants-bg { background-color: #dcfce7; }
.flowers-bg { background: linear-gradient(to bottom, #fecaca, #ef4444); }
.papa-bg { background: linear-gradient(to bottom, #bbf7d0, #22c55e); height: 120px; flex: none; }

.app-details {
  padding: 12px;
  font-size: 10px;
}

.app-details h4 {
  margin: 0 0 4px 0;
  font-size: 12px;
}

.app-details p {
  margin: 0;
  color: #64748b;
  font-size: 8px;
}

/* Chart Section */
.chart-section {
  display: flex;
  align-items: center;
  justify-content: space-around;
}

.chart-legend {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.legend-item {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 14px;
  color: var(--text-dark, #1a1d29);
}

.dot {
  width: 10px;
  height: 10px;
  border-radius: 2px;
  display: inline-block;
}

.bg-blue { background-color: #4a90d9; }
.bg-purple { background-color: #7b61ff; }
.bg-lightblue { background-color: #38bdf8; }
.bg-red { background-color: #f43f5e; }

@media (max-width: 1024px) {
  .top-cards {
    grid-template-columns: 1fr;
  }
}
</style>
