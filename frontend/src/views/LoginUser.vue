<template>
  <div class="flex items-center justify-center bg-gray-100">
    <div class="w-full max-w-md p-8 bg-white rounded-lg shadow-lg">
      <h2 class="mb-6 text-2xl font-bold text-center text-gray-800">Login</h2>
      <form @submit.prevent="handleLogin" class="space-y-4">
        <!--email input-->
        <input
          v-model="email"
          type="email"
          placeholder="Email"
          class="w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
          required
        />

        <!--pass input-->
        <input
          v-model="password"
          type="password"
          placeholder="Password"
          class="w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
          required
        />

        <!--submit button-->
        <button
          :disabled="loading"
          type="submit"
          class="w-full p-3 font-semibold text-white bg-blue-600 rounded-lg hover:bg-blue-700 disabled:opacity-50 disabled:cursor-not-allowed"
        >
          {{ loading ? 'Logging in...' : 'Login' }}
        </button>
      </form>

      <!--err/success message-->
      <p
        v-if="message"
        :class="messageType === 'error' ? 'text-red-500' : 'text-green-500'"
        class="mt-4 text-center"
      >
        {{ message }}
      </p>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useAuthStore } from '@/stores/auth'

//refs for email and pass inputs
const email = ref('')
const password = ref('')

//refs for messages
const message = ref('')
const messageType = ref('') //to display success or error message
const loading = ref(false) //handle button state to prevent multiple clicks

//initialize auth store
const authStore = useAuthStore()

//function to handle login
const handleLogin = async () => {
  message.value = ''
  messageType.value = ''
  loading.value = true //set loading to true to disable button

  //use auth store login action
  const result = await authStore.login(email.value, password.value)

  //handle response from auth store action
  if (result.success) {
    message.value = result.message
    messageType.value = 'success'

    //redirect to dash after 1 second delay to display success message
    setTimeout(() => {
      window.location.href = '/'
    }, 500)
  } else {
    //display error message
    message.value = result.message
    messageType.value = 'error'
    password.value = '' //clear pass field after failed login
  }

  loading.value = false //set loading to false to enable button
}
</script>
