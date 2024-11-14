<template>
  <div class="flex flex-col min-h-screen bg-gray-50">
    <Navbar />
    <router-view class="flex-1 p-8" />
  </div>
</template>

<script setup>
import Navbar from '@/components/NavBar.vue';
</script>
