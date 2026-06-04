import { defineStore } from 'pinia'
import api from '../services/api'

export const useAuthStore = defineStore('auth', {
  state: () => ({
    user: JSON.parse(localStorage.getItem('user')) || null,
    token: localStorage.getItem('token') || null,
    loading: false,
    error: null
  }),

  getters: {
    isAuthenticated: (state) => !!state.token,
    currentUser: (state) => state.user || {
      id: 1,
      firstName: 'Yash',
      lastName: 'Ghori',
      email: 'yash.ghori@asite.com',
      avatar: 'https://i.pravatar.cc/150?img=11',
      role: 'Admin',
      designation: 'UI Intern',
      location: 'Ahmedabad, Gujarat, India',
      phone: '+91 9876543210',
      skills: ['PGT'],
      nationality: 'Indian'
    }
  },

  actions: {
    async login(email, password) {
      this.loading = true
      this.error = null
      try {
        const response = await api.post('/auth/login', { email, password })
        this.token = response.data.token
        this.user = response.data.user
        localStorage.setItem('token', this.token)
        localStorage.setItem('user', JSON.stringify(this.user))
        return true
      } catch (err) {
        // For demo: allow login with any credentials
        this.token = 'demo-jwt-token-group13'
        this.user = {
          id: 1,
          firstName: 'Yash',
          lastName: 'Ghori',
          email: email || 'yash.ghori@asite.com',
          avatar: 'https://i.pravatar.cc/150?img=11',
          role: 'Admin',
          designation: 'UI Intern',
          location: 'Ahmedabad, Gujarat, India',
          phone: '+91 9876543210',
          skills: ['PGT'],
          nationality: 'Indian'
        }
        localStorage.setItem('token', this.token)
        localStorage.setItem('user', JSON.stringify(this.user))
        return true
      } finally {
        this.loading = false
      }
    },

    logout() {
      this.token = null
      this.user = null
      localStorage.removeItem('token')
      localStorage.removeItem('user')
    },

    async updateProfile(profileData) {
      this.loading = true
      try {
        const response = await api.put('/auth/profile', profileData)
        this.user = { ...this.user, ...response.data }
        localStorage.setItem('user', JSON.stringify(this.user))
      } catch (err) {
        // For demo: update locally
        this.user = { ...this.user, ...profileData }
        localStorage.setItem('user', JSON.stringify(this.user))
      } finally {
        this.loading = false
      }
    }
  }
})
