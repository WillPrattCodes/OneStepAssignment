<template>
  <div class="flex-col h-screen bg-gray-50">
    <Navbar class="h-[7vh]" />
    <router-view class="flex-1 h-[93vh] overflow-hidden" />
  </div>
</template>

<script setup>
import Navbar from '@/components/NavBar.vue';
</script>
