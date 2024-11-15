<template>
  <div class="flex bg-gray-100">
    <!--sidebar with devices-->
    <div class="w-1/3 h-screen p-4 overflow-y-auto scrollbar-thin scrollbar-thumb-gray-400 scrollbar-track-gray-200">
      <div class="flex items-center justify-between mb-6">
        <h2 class="text-2xl font-bold text-gray-800">Devices</h2>
        <!--filter button component-->
        <FilterButton @click="openPreferences" />
      </div>

      <!--preferences menu-->
      <div v-if="showPreferences" class="p-4 mb-6 bg-white border rounded shadow-lg">
        <h3 class="mb-4 text-lg font-semibold text-gray-800">Sort Devices</h3>
        <button
          @click="setSortOrder('asc')"
          :class="['w-full p-2 mt-2 rounded', preferences.sortOrder === 'asc' ? 'bg-blue-600 text-white' : 'bg-gray-200']"
        >
          Ascending
        </button>
        <button
          @click="setSortOrder('desc')"
          :class="['w-full p-2 mt-2 rounded', preferences.sortOrder === 'desc' ? 'bg-blue-600 text-white' : 'bg-gray-200']"
        >
          Descending
        </button>
      </div>

      <!--render dvice list using loop through device data-->
      <ul class="space-y-4">
        <li 
          v-for="(device, index) in sortedDevices" 
          :key="index" 
          class="p-4 transition rounded-lg shadow-sm bg-gray-50 hover:bg-gray-100"
        >
          <h3 class="font-semibold text-gray-700">{{ device.display_name }}</h3>
          <p class="text-sm text-gray-600">
            <strong>Location:</strong> Lat: {{ device.latest_device_point?.lat }}, Lng: {{ device.latest_device_point?.lng }}
          </p>
          <p class="text-sm text-gray-600">
            <strong>Speed:</strong> {{ Math.round(device.latest_device_point.speed) }}
          </p>
          <p class="text-sm text-gray-600">
            <strong>Status:</strong> {{ device.online ? 'Online' : 'Offline' }}
          </p>
        </li>
      </ul>
    </div>

    <!--google map-->
    <div class="flex-1 h-screen">
      <GoogleMap />
    </div>
  </div>
</template>

<script setup>
//importing necessary functions, components and stores
import { ref, onMounted, onBeforeUnmount, computed } from 'vue';
import { useGPSStore } from '@/stores/gps';
import { usePrefStore } from '@/stores/pref';
import GoogleMap from '@/components/GoogleMap.vue';
import FilterButton from '@/components/FilterButton.vue';

const gpsStore = useGPSStore();
const prefStore = usePrefStore();
const showPreferences = ref(false);
const preferences = ref({ sortOrder: 'asc' });
let fetchInterval;

//toggle preferences menu
const openPreferences = () => {
  showPreferences.value = !showPreferences.value;
};

//get user preferences on mount
const fetchPreferences = async () => {
  await prefStore.fetchPreferences();
  preferences.value.sortOrder = prefStore.preferences.sortOrder || 'asc';
};

//save user sort order preference
const setSortOrder = (order) => {
  preferences.value.sortOrder = order;
  prefStore.setPreferences(preferences.value);
};

//create sorted devices list based on user preference
const sortedDevices = computed(() => {
  const devices = [...gpsStore.gpsData];
  if (preferences.value.sortOrder === 'asc') {
    return devices.sort((a, b) => a.display_name.localeCompare(b.display_name));
  } else if (preferences.value.sortOrder === 'desc') {
    return devices.sort((a, b) => b.display_name.localeCompare(a.display_name));
  }
  return devices;
});

//function to start polling GPS data
const startPolling = () => {
  gpsStore.fetchGPSData(); 

  //poll every 10 seconds
  fetchInterval = setInterval(() => {
    console.log('fetching updated GPS data');
    gpsStore.fetchGPSData();
  }, 10000);
};

//stop polling GPS data when component unmounted
const stopPolling = () => {
  if (fetchInterval) clearInterval(fetchInterval);
};

//fetch user preferences and start polling GPS data on mount
onMounted(async () => {
  await fetchPreferences();
  startPolling();
});

//stop polling GPS data on unmount
onBeforeUnmount(stopPolling);
</script>
