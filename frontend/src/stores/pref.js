import { defineStore } from 'pinia';
import eventBus from '@/utils/eventBus';
import api from '@/utils/axios';

export const usePrefStore = defineStore('pref', {
  state: () => ({
    preferences: {
      sortOrder: 'asc',
      hiddenDevices: []
    },
  }),
  actions: {
    //fetch preferences from the server
    async fetchPreferences() {
      try {
        const response = await api.get('/api/getpreferences');
        this.preferences = response.data;

      } catch (error) {
        console.error('Failed to fetch preferences:', error);
      }
    },
    //set preferences func
    async setPreferences(prefs) {
      await api.post('/api/setpreferences', prefs);
      this.preferences = prefs;

    },
    async toggleDeviceVisibility(deviceId) {
      //make sure hiddenDevices is an array (was having issues with it being a string)
      if (!Array.isArray(this.preferences.hiddenDevices)) {
        this.preferences.hiddenDevices = [];
      }

      const isHidden = this.preferences.hiddenDevices.includes(deviceId);
      if (isHidden) {
        //remove device from hiddenDevices if hidden
        this.preferences.hiddenDevices = this.preferences.hiddenDevices.filter(id => id !== deviceId);
      } else {
        //add device to hiddenDevices if not hidden
        this.preferences.hiddenDevices.push(deviceId);
      }
      //call setPreferences to save the new preferences
      await this.setPreferences(this.preferences);

      //emit event to update trigger a marker update in GoogleMap.vue
      eventBus.emit('toggle-visibility');
    },
  },
  getters: {
    //check if device is hidden. used in VisibilityButton.vue
    isDeviceHidden: (state) => (deviceId) => Array.isArray(state.preferences.hiddenDevices) && state.preferences.hiddenDevices.includes(deviceId),
  },
});
