import { defineStore } from 'pinia';
import api from '@/utils/axios';

export const usePrefStore = defineStore('pref', {
  state: () => ({
    preferences: {
      sortOrder: 'asc',
    },
  }),
  actions: {
    async fetchPreferences() {
      const response = await api.get('/api/getpreferences');
      this.preferences = response.data;
    },
    async setPreferences(prefs) {
      await api.post('/api/setpreferences', prefs);
      this.preferences = prefs;
      console.log('Preferences set:', prefs);
    },
  },
});
