import { defineStore } from 'pinia'
import api from '@/utils/axios'

export const useAuthStore = defineStore('auth', {
  state: () => ({
    token: localStorage.getItem('token') || '',
    user: JSON.parse(localStorage.getItem('user')) || null,
  }),
  actions: {
    //login function
    async login(email, password) {
      try {
        const response = await api.post('/api/login', { email, password })
        this.token = response.data.token
        this.user = response.data.user
        localStorage.setItem('token', this.token)
        localStorage.setItem('user', JSON.stringify(this.user))
        return { success: true, message: 'Login successful' }
      } catch (error) {
        if (error.response && error.response.status === 401) {
          return { success: false, message: 'Invalid email or password. Try again.' }
        }
        return { success: false, message: error.response?.data || 'Login failed. Try again later.' }
      }
    },
    //register function
    async register(username, email, password) {
      try {
        const response = await api.post('/api/register', { username, email, password })

        //save token and user to localStorage and state
        this.token = response.data.token
        this.user = response.data.user
        localStorage.setItem('token', this.token)
        localStorage.setItem('user', JSON.stringify(this.user))

        return { success: true, message: 'Registration successful' }
      } catch (error) {
        if (error.response && error.response.status === 409) {
          return { success: false, message: 'Email or username already exists.' }
        }
        return {
          success: false,
          message: error.response?.data || 'Registration failed. Please try again later.',
        }
      }
    },
    //logout function
    logout() {
      this.token = ''
      this.user = null
      localStorage.removeItem('token')
      window.location.href = '/login'
    },
  },
  getters: {
    isAuthenticated: (state) => !!state.token,
  },
})
