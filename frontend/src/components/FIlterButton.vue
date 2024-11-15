<template>
  <button 
    @click="handleClick"
    class="p-2 text-gray-500 border border-gray-300 rounded-full hover:text-gray-700 focus:outline-none hover:bg-gray-100"
  >
    <svg xmlns="http://www.w3.org/2000/svg" class="w-6 h-6" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
      <path stroke-linecap="round" stroke-linejoin="round" d="M3 6h18M6 12h12m-9 6h6" />
    </svg>
  </button>
</template>

<script setup>
const props = defineProps(['onClick']);
const handleClick = () => {
  if (props.onClick) props.onClick();
};
</script>
