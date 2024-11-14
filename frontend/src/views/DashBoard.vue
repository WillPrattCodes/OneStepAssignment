<template>
  <div class="flex flex-col items-center justify-center min-h-screen bg-gray-100">
    <div class="w-full max-w-3xl p-8 bg-white rounded-lg shadow-lg">
      <h2 class="mb-6 text-3xl font-bold text-center text-gray-800">Welcome, {{ user?.username }}!</h2>
      <p class="mb-8 text-center text-gray-600">Here is your dashboard:</p>

      <!--conditional gps data-->
      <div v-if="gpsStore.gpsData.length > 0" class="mb-6">
        <h3 class="mb-4 text-xl font-semibold text-gray-700">Your GPS Data:</h3>
        <ul>
          <li v-for="(data, index) in gpsStore.gpsData" :key="index" class="py-2 border-b border-gray-300">
            <p><strong>Device:</strong> {{ data.display_name }}</p>
            <p><strong>Location:</strong> Lat: {{ data.latest_device_point.lat }}, Lng: {{ data.latest_device_point.lng }}</p>
            <p><strong>Status:</strong> {{ data.online ? 'Online' : 'Offline' }}</p>
          </li>
        </ul>
      </div>
      <div v-else class="text-gray-500">No GPS data available.</div>

      <!--conditional err message-->
      <p v-if="gpsStore.error" class="text-red-500">{{ gpsStore.error }}</p>
    </div>
  </div>
</template>

<script setup>
import { onMounted } from 'vue';
import { useAuthStore } from '@/stores/auth';
import { useGPSStore } from '@/stores/gps';

//get auth and gps store
const authStore = useAuthStore();
const gpsStore = useGPSStore();
const user = authStore.user;

//fetch gps data on mounted
onMounted(() => {
  gpsStore.fetchGPSData();
});

</script>
