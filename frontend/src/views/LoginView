<template>
  <div class="login-container">
    <div class="login-left">
      <div class="logo-area">
        <div class="logo-icon"></div>
        <span class="logo-text">AProjectO</span>
      </div>
      
      <div class="illustration-area">
        <div class="login-card-illustration">
          <h2 class="card-title">LOGIN ACCESS</h2>
          <div class="shield-icon">
            <svg xmlns="http://www.w3.org/2000/svg" width="48" height="48" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z"></path><rect x="9" y="11" width="6" height="6" rx="1"></rect><path d="M12 11V9a2 2 0 0 1 4 0"></path></svg>
          </div>
          <div class="form-mockup">
            <div class="input-mockup">
              <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"></path><circle cx="12" cy="7" r="4"></circle></svg>
              <span>USERNAME</span>
            </div>
            <div class="input-mockup">
              <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect x="3" y="11" width="18" height="11" rx="2" ry="2"></rect><path d="M7 11V7a5 5 0 0 1 10 0v4"></path></svg>
              <span>--------</span>
            </div>
          </div>
        </div>
        <div class="character-illustration"></div>
      </div>
    </div>
    
    <div class="login-right">
      <div class="top-brand">
        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="12" y1="2" x2="12" y2="22"></line><line x1="2" y1="12" x2="22" y2="12"></line><line x1="4.93" y1="4.93" x2="19.07" y2="19.07"></line><line x1="4.93" y1="19.07" x2="19.07" y2="4.93"></line></svg>
        <span>Asite Product System</span>
      </div>
      
      <div class="login-form-container">
        <h1 class="welcome-heading">Welcome back, Yash</h1>
        <p class="welcome-sub">Welcome back! Please enter your details.</p>
        
        <form @submit.prevent="handleLogin" class="login-form">
          <div class="form-group">
            <label>Email</label>
            <input type="email" v-model="email" required />
          </div>
          
          <div class="form-group">
            <label>Password</label>
            <div class="password-input">
              <input :type="showPassword ? 'text' : 'password'" v-model="password" required />
              <button type="button" @click="showPassword = !showPassword" class="toggle-password">
                <svg v-if="!showPassword" xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M17.94 17.94A10.07 10.07 0 0 1 12 20c-7 0-11-8-11-8a18.45 18.45 0 0 1 5.06-5.94M9.9 4.24A9.12 9.12 0 0 1 12 4c7 0 11 8 11 8a18.5 18.5 0 0 1-2.16 3.19m-6.72-1.07a3 3 0 1 1-4.24-4.24"></path><line x1="1" y1="1" x2="23" y2="23"></line></svg>
                <svg v-else xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z"></path><circle cx="12" cy="12" r="3"></circle></svg>
              </button>
            </div>
          </div>
          
          <div class="form-actions">
            <label class="checkbox-label">
              <input type="checkbox" v-model="rememberMe" />
              <span>Terms & Conditions</span>
            </label>
            <a href="#" class="forgot-link">Forgot Password</a>
          </div>
          
          <button type="submit" class="login-button" :disabled="loading">
            {{ loading ? 'Logging in...' : 'Log In' }}
          </button>
        </form>
        
        <p class="signup-prompt">
          Don't have an account? <a class="signup-link" @click.prevent="goToSignup">Sign up for free</a>
        </p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { useAuthStore } from '../stores/auth';

const router = useRouter();
const authStore = useAuthStore();

const email = ref('admin@erp.com');
const password = ref('admin123');
const showPassword = ref(false);
const rememberMe = ref(false);
const loading = ref(false);

const handleLogin = async () => {
  try {
    loading.value = true;
    const success = await authStore.login({ email: email.value, password: password.value });
    if (success) {
      router.push('/');
    } else {
      alert('Login failed. Please check your credentials.');
    }
  } catch (error) {
    console.error('Login error:', error);
    alert(error.response?.data?.error || 'An error occurred during login.');
  } finally {
    loading.value = false;
  }
};

const goToSignup = () => {
  // Clear any existing auth state so router guard allows navigation
  authStore.logout();
  router.push('/signup');
};
</script>

<style scoped>
.login-container {
  display: flex;
  height: 100vh;
  width: 100vw;
  background-color: var(--bg-color, #f5f7fa);
}

.login-left {
  flex: 1;
  background: linear-gradient(135deg, #f5f7fa 0%, #e0e7ff 100%);
  position: relative;
  display: flex;
  flex-direction: column;
  padding: 40px;
  overflow: hidden;
}

.logo-area {
  display: flex;
  align-items: center;
  gap: 12px;
  z-index: 10;
}

.logo-icon {
  width: 24px;
  height: 24px;
  background-color: var(--primary-color, #4a90d9);
  border-radius: 50%;
  position: relative;
}

.logo-icon::before {
  content: '';
  position: absolute;
  top: 0;
  left: -8px;
  width: 24px;
  height: 24px;
  background-color: rgba(74, 144, 217, 0.6);
  border-radius: 50%;
}

.logo-text {
  font-size: 20px;
  font-weight: 700;
  color: var(--text-dark, #1a1d29);
}

.illustration-area {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  position: relative;
}

.login-card-illustration {
  background: linear-gradient(180deg, #4da4f0 0%, #2974d6 100%);
  width: 280px;
  height: 400px;
  border-radius: 16px;
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.2);
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 30px 20px;
  color: white;
  z-index: 2;
  position: relative;
}

.card-title {
  font-size: 20px;
  font-weight: 700;
  margin-bottom: 30px;
  letter-spacing: 1px;
}

.shield-icon {
  margin-bottom: 40px;
  opacity: 0.9;
}

.form-mockup {
  width: 100%;
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.input-mockup {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 12px;
  background-color: rgba(255, 255, 255, 0.1);
  border-radius: 8px;
  font-size: 12px;
  letter-spacing: 1px;
}

.login-right {
  flex: 1;
  background-color: white;
  display: flex;
  flex-direction: column;
  padding: 40px;
  position: relative;
}

.top-brand {
  position: absolute;
  top: 40px;
  right: 40px;
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 14px;
  color: var(--text-dark, #1a1d29);
  font-weight: 500;
}

.login-form-container {
  max-width: 400px;
  margin: auto;
  width: 100%;
}

.welcome-heading {
  font-size: 32px;
  font-weight: 700;
  color: var(--text-dark, #1a1d29);
  margin-bottom: 8px;
}

.welcome-sub {
  font-size: 15px;
  color: var(--text-gray, #6b7280);
  margin-bottom: 40px;
}

.form-group {
  margin-bottom: 24px;
}

.form-group label {
  display: block;
  font-size: 14px;
  font-weight: 500;
  color: var(--text-dark, #1a1d29);
  margin-bottom: 8px;
}

.form-group input[type="email"],
.form-group input[type="text"],
.form-group input[type="password"] {
  width: 100%;
  padding: 12px 0;
  border: none;
  border-bottom: 1px solid #e5e7eb;
  font-size: 15px;
  outline: none;
  transition: border-color 0.2s;
}

.form-group input:focus {
  border-bottom-color: var(--primary-color, #4a90d9);
}

.password-input {
  position: relative;
}

.toggle-password {
  position: absolute;
  right: 0;
  top: 50%;
  transform: translateY(-50%);
  background: none;
  border: none;
  color: var(--text-gray, #6b7280);
  cursor: pointer;
}

.form-actions {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 30px;
  font-size: 13px;
}

.checkbox-label {
  display: flex;
  align-items: center;
  gap: 8px;
  color: var(--text-gray, #6b7280);
  cursor: pointer;
}

.forgot-link {
  color: var(--text-gray, #6b7280);
  text-decoration: none;
}

.login-button {
  width: 100%;
  padding: 14px;
  background-color: var(--text-dark, #1a1d29);
  color: white;
  border: none;
  border-radius: 8px;
  font-size: 16px;
  font-weight: 500;
  cursor: pointer;
  transition: background-color 0.2s;
  margin-bottom: 24px;
}

.login-button:hover {
  background-color: #000;
}

.signup-prompt {
  text-align: center;
  font-size: 14px;
  color: var(--text-gray, #6b7280);
}

.signup-link {
  color: var(--text-dark, #1a1d29);
  font-weight: 600;
  text-decoration: none;
  cursor: pointer;
}

.signup-link:hover {
  text-decoration: underline;
}

@media (max-width: 900px) {
  .login-left {
    display: none;
  }
}
</style>
