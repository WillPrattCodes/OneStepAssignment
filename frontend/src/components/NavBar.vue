<template>
  <nav class="flex items-center justify-between p-4 text-white bg-blue-600">
    <div>
      <router-link v-if="isAuthenticated" to="/" class="mr-4 hover:underline">Home</router-link>
      <router-link v-if="!isAuthenticated" to="/register" class="mr-4 hover:underline"
        >Register</router-link
      >
      <router-link v-if="!isAuthenticated" to="/login" class="mr-4 hover:underline"
        >Login</router-link
      >
    </div>
    <div class="absolute transform -translate-x-1/2 left-1/2">
      <h1 class="text-xl">GPS App</h1>
    </div>
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
import { computed } from 'vue'
import { useAuthStore } from '@/stores/auth'

const authStore = useAuthStore()

//computed property to check if user is authenticated
const isAuthenticated = computed(() => authStore.isAuthenticated)

//func to handle logout
const handleLogout = () => {
  authStore.logout()
}
</script>