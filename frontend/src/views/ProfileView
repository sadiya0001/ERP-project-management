<template>
  <div class="profile-page">
    <div class="profile-layout">
      <!-- Left Column: User Card & Stats -->
      <div class="left-col">
        <div class="user-card">
          <div class="avatar-container">
            <img :src="user.avatar" :alt="user.name" class="user-avatar" />
          </div>
          <h3 class="user-name">{{ user.name }}</h3>
          <p class="user-location">{{ user.location }}</p>
          
          <div class="user-stats">
            <div class="stat-row">
              <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M17 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"></path><circle cx="9" cy="7" r="4"></circle><path d="M23 21v-2a4 4 0 0 0-3-3.87"></path><path d="M16 3.13a4 4 0 0 1 0 7.75"></path></svg>
              <span>{{ user.teamMembers }} - Interns</span>
            </div>
            <div class="stat-row">
              <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"></circle><line x1="4.93" y1="4.93" x2="19.07" y2="19.07"></line></svg>
              <span>{{ user.activeTasks }} Task</span>
            </div>
          </div>
          
          <div class="user-contact">
            <div class="contact-row">
              <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M22 16.92v3a2 2 0 0 1-2.18 2 19.79 19.79 0 0 1-8.63-3.07 19.5 19.5 0 0 1-6-6 19.79 19.79 0 0 1-3.07-8.67A2 2 0 0 1 4.11 2h3a2 2 0 0 1 2 1.72 12.84 12.84 0 0 0 .7 2.81 2 2 0 0 1-.45 2.11L8.09 9.91a16 16 0 0 0 6 6l1.27-1.27a2 2 0 0 1 2.11-.45 12.84 12.84 0 0 0 2.81.7A2 2 0 0 1 22 16.92z"></path></svg>
              <span>{{ user.phone }}</span>
            </div>
            <div class="contact-row">
              <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M4 4h16c1.1 0 2 .9 2 2v12c0 1.1-.9 2-2 2H4c-1.1 0-2-.9-2-2V6c0-1.1.9-2 2-2z"></path><polyline points="22,6 12,13 2,6"></polyline></svg>
              <span>{{ user.email }}</span>
            </div>
            <div class="contact-row">
              <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect x="3" y="4" width="18" height="18" rx="2" ry="2"></rect><line x1="16" y1="2" x2="16" y2="6"></line><line x1="8" y1="2" x2="8" y2="6"></line><line x1="3" y1="10" x2="21" y2="10"></line></svg>
              <span>{{ user.skills }}</span>
            </div>
          </div>
        </div>
        
        <div class="work-done-card">
          <div class="work-done-header">
            <h4>Total work done</h4>
            <div class="badge">This Week</div>
          </div>
          <RingChart time="5w: 2d" :percentage="75" />
        </div>
      </div>
      
      <!-- Right Column: Edit Profile Form -->
      <div class="right-col">
        <div class="edit-profile-card">
          <h2>Edit Profile</h2>
          
          <form @submit.prevent="saveProfile" class="profile-form">
            <div class="form-row two-cols">
              <div class="form-group">
                <label>First Name</label>
                <input type="text" v-model="form.firstName" />
              </div>
              <div class="form-group">
                <label>Last Name</label>
                <input type="text" v-model="form.lastName" />
              </div>
            </div>
            
            <div class="form-group">
              <label>Email</label>
              <input type="email" v-model="form.email" />
            </div>
            
            <div class="form-row two-cols">
              <div class="form-group">
                <label>Password</label>
                <input type="password" placeholder="Change Password" />
              </div>
              <div class="form-group">
                <label>Nationality</label>
                <select v-model="form.nationality">
                  <option value="India">India</option>
                  <option value="USA">USA</option>
                  <option value="UK">UK</option>
                </select>
              </div>
            </div>
            
            <div class="form-row two-cols">
              <div class="form-group">
                <label>Phone Number</label>
                <div class="phone-input-group">
                  <select class="country-code">
                    <option value="+234">🇳🇬 +234</option>
                    <option value="+91">🇮🇳 +91</option>
                    <option value="+1">🇺🇸 +1</option>
                    <option value="+44">🇬🇧 +44</option>
                  </select>
                  <input type="text" v-model="form.phone" placeholder="Phone number" />
                </div>
              </div>
              <div class="form-group">
                <label>Designation</label>
                <select v-model="form.designation">
                  <option value="UI Intern">UI Intern</option>
                  <option value="Developer">Developer</option>
                  <option value="Designer">Designer</option>
                  <option value="Manager">Manager</option>
                </select>
              </div>
            </div>
            
            <div class="form-actions">
              <button type="submit" class="save-btn">Save</button>
            </div>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { reactive, computed } from 'vue';
import RingChart from '../components/RingChart.vue';
import { useAuthStore } from '../stores/auth';

const authStore = useAuthStore();

const user = computed(() => {
  const currentUser = authStore.currentUser || {};
  return {
    ...currentUser,
    teamMembers: currentUser.teamMembers || 0,
    activeTasks: currentUser.activeTasks || 0
  };
});

const form = reactive({
  firstName: authStore.currentUser?.firstName || '',
  lastName: authStore.currentUser?.lastName || '',
  email: authStore.currentUser?.email || '',
  phone: authStore.currentUser?.phone || '',
  nationality: authStore.currentUser?.nationality || '',
  designation: authStore.currentUser?.designation || ''
});

const saveProfile = () => {
  alert('Profile saved successfully!');
};
</script>

<style scoped>
.profile-page {
  padding: 24px;
  max-width: 1200px;
  margin: 0 auto;
}

.profile-layout {
  display: grid;
  grid-template-columns: 320px 1fr;
  gap: 24px;
}

/* Left Column Styles */
.user-card {
  background-color: var(--card-bg, #ffffff);
  border-radius: 12px;
  padding: 30px 24px;
  text-align: center;
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.05);
  margin-bottom: 24px;
}

.avatar-container {
  width: 120px;
  height: 120px;
  margin: 0 auto 16px;
  border-radius: 50%;
  border: 4px solid #ec4899; /* Pink ring */
  padding: 4px;
}

.user-avatar {
  width: 100%;
  height: 100%;
  border-radius: 50%;
  object-fit: cover;
}

.user-name {
  font-size: 18px;
  font-weight: 700;
  color: var(--text-dark, #1a1d29);
  margin: 0 0 8px 0;
}

.user-location {
  font-size: 14px;
  color: var(--text-gray, #6b7280);
  margin: 0 0 24px 0;
  white-space: pre-line;
}

.user-stats, .user-contact {
  border-top: 1px solid #f3f4f6;
  padding-top: 16px;
  margin-top: 16px;
  text-align: left;
}

.stat-row, .contact-row {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 12px;
  font-size: 14px;
  color: var(--text-dark, #1a1d29);
}

.stat-row svg, .contact-row svg {
  color: var(--text-gray, #6b7280);
}

.work-done-card {
  background-color: var(--card-bg, #ffffff);
  border-radius: 12px;
  padding: 24px;
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.05);
}

.work-done-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
}

.work-done-header h4 {
  font-size: 16px;
  font-weight: 600;
  margin: 0;
  color: var(--text-dark, #1a1d29);
}

.badge {
  background-color: #e0f2fe;
  color: #0369a1;
  padding: 4px 10px;
  border-radius: 12px;
  font-size: 12px;
  font-weight: 500;
}

/* Right Column Styles */
.edit-profile-card {
  background-color: var(--card-bg, #ffffff);
  border-radius: 12px;
  padding: 32px;
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.05);
  min-height: 100%;
}

.edit-profile-card h2 {
  font-size: 20px;
  font-weight: 600;
  color: var(--text-dark, #1a1d29);
  margin: 0 0 32px 0;
}

.profile-form {
  display: flex;
  flex-direction: column;
  gap: 24px;
}

.form-row.two-cols {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 24px;
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.form-group.half-width {
  width: 50%;
  padding-right: 12px;
}

.form-group label {
  font-size: 14px;
  font-weight: 500;
  color: var(--text-gray, #6b7280);
}

.form-group input, .form-group select {
  padding: 12px 16px;
  border: 1px solid #e5e7eb;
  border-radius: 8px;
  font-size: 15px;
  color: var(--text-dark, #1a1d29);
  outline: none;
  transition: border-color 0.2s;
  background-color: #fff;
}

.form-group input:focus, .form-group select:focus {
  border-color: var(--primary-color, #4a90d9);
}

.phone-input-group {
  display: flex;
  gap: 8px;
}

.country-code {
  width: 100px;
}

.phone-input-group input {
  flex: 1;
}

.form-actions {
  margin-top: 16px;
  display: flex;
  justify-content: center;
}

.save-btn {
  background-color: var(--primary-color, #4a90d9);
  color: white;
  border: none;
  border-radius: 8px;
  padding: 12px 48px;
  font-size: 16px;
  font-weight: 500;
  cursor: pointer;
  transition: background-color 0.2s;
}

.save-btn:hover {
  background-color: #3b76b8;
}

@media (max-width: 1024px) {
  .profile-layout {
    grid-template-columns: 1fr;
  }
  
  .form-group.half-width {
    width: 100%;
    padding-right: 0;
  }
}

@media (max-width: 768px) {
  .form-row.two-cols {
    grid-template-columns: 1fr;
  }
}
</style>
