<template>
  <nav class="flex items-center justify-between p-4 text-white bg-blue-600">
    <div>
      <router-link to="/" class="mr-4 hover:underline">Home</router-link>
      <router-link to="/register" class="mr-4 hover:underline">Register</router-link>
      <router-link to="/login" class="mr-4 hover:underline">Login</router-link>
      <router-link v-if="isAuthenticated" to="/dashboard" class="mr-4 hover:underline">Dashboard</router-link>
    </div>
    <!-- Logout Button -->
    <button 
      v-if="isAuthenticated" 
      @click="handleLogout" 
      class="px-4 py-2 text-white bg-red-500 rounded hover:bg-red-600"
    >
      Logout
    </button>
  </nav>
</template>

<script setup>
import { computed } from 'vue';
import { useRouter } from 'vue-router';
import { useAuthStore } from '@/stores/auth'; // Import the Pinia store

const authStore = useAuthStore();
const router = useRouter();

// Computed property to check if user is authenticated
const isAuthenticated = computed(() => authStore.isAuthenticated);

// Function to handle logout
const handleLogout = () => {
  authStore.logout();
  router.push('/login'); // Redirect to login page after logout
};
</script>
