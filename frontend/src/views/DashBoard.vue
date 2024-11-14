<template>
  <div class="flex bg-gray-100">
    <!--sidebar containing devices-->
    <div class="w-1/3 h-screen p-4 overflow-y-auto scrollbar-thin scrollbar-thumb-gray-400 scrollbar-track-gray-200">
      <h2 class="mb-6 text-2xl font-bold text-gray-800">Devices</h2>
      <ul class="space-y-4">
        <!--render list using loop containing all device info-->
        <li 
          v-for="(device, index) in gpsStore.gpsData" 
          :key="index" 
          class="p-4 transition rounded-lg shadow-sm bg-gray-50 hover:bg-gray-100"
        >
          <h3 class="font-semibold text-gray-700">{{ device.display_name }}</h3>
          <p class="text-sm text-gray-600">
            <strong>Location:</strong> Lat: {{ device.latest_device_point?.lat }}, Lng: {{ device.latest_device_point?.lng }}
          </p>
          <p class="text-sm text-gray-600">
            <strong>Status:</strong> {{ device.online ? 'Online' : 'Offline' }}
          </p>
        </li>
      </ul>
    </div>

    <!--google maps-->
    <div class="flex-1 h-screen">
      <GoogleMap />
    </div>
  </div>
</template>

<script setup>
import { onMounted, onBeforeUnmount } from 'vue';
import { useGPSStore } from '@/stores/gps';
import GoogleMap from '@/components/GoogleMap.vue';

const gpsStore = useGPSStore();
let fetchInterval;

//start polling to fetch gps data
const startPolling = () => {
  gpsStore.fetchGPSData(); //fetch initial gps data

  //fetch new data every 10 seconds
  fetchInterval = setInterval(() => {
    console.log('Fetching updated GPS data...');
    gpsStore.fetchGPSData();
  }, 10000);
};

//stop olling to clean up gps data
const stopPolling = () => {
  if (fetchInterval) {
    clearInterval(fetchInterval);
  }
};

onMounted(startPolling); //start polling when component is mounted
onBeforeUnmount(stopPolling); //stop polling when component is unmounted
</script>
