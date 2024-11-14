import { defineStore } from 'pinia';
import api from '@/utils/axios';
import { useAuthStore } from './auth';

export const useGPSStore = defineStore('gps', {
  state: () => ({
    gpsData: [],
    loading: false,
    error: null,
  }),

  actions: {
    async fetchGPSData() {
      this.loading = true;
      this.error = null;
      const authStore = useAuthStore();

      try {
        // Fetch GPS data from the server using the token
        const response = await api.get('/api/gps', {
          headers: {
            Authorization: `Bearer ${authStore.token}`,
          },
        });
        this.gpsData = response.data;
        console.log('Fetched GPS Data:', response.data);
      } catch (error) {
        console.error('Failed to fetch GPS data:', error);
        this.error = 'Failed to fetch GPS data';
      } finally {
        this.loading = false;
      }
    },
  },

  getters: {
    getGPSData: (state) => state.gpsData,
    
  },
});