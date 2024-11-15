<template>
  <div class="flex items-center justify-center bg-gray-100">
    <div class="w-full max-w-md p-8 bg-white rounded-lg shadow-lg">
      <h2 class="mb-6 text-2xl font-bold text-center text-gray-800">Register</h2>
      <form @submit.prevent="registerUser" class="space-y-4">
        <!--username input-->
        <input
          v-model="username"
          type="text"
          placeholder="Username"
          class="w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
          required
        />

        <!--email input-->
        <input
          v-model="email"
          type="email"
          placeholder="Email"
          class="w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
          required
        />

        <!--password input-->
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
          {{ loading ? 'Registering...' : 'Register' }}
        </button>
      </form>

      <!--err or success message-->
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

//refs for inputs
const username = ref('')
const email = ref('')
const password = ref('')

//refs for messages and button state
const message = ref('')
const messageType = ref('')
const loading = ref(false)

//initialize auth store
const authStore = useAuthStore()

//function to handle registration
const registerUser = async () => {
  message.value = ''
  messageType.value = ''
  loading.value = true

  //use the auth store's register action
  const result = await authStore.register(username.value, email.value, password.value)

  //handle response from store action
  if (result.success) {
    message.value = result.message
    messageType.value = 'success'

    //redirect to dashboard page after successful registration
    setTimeout(() => {
      window.location.href = '/'
    }, 500)
  } else {
    // handle errors and display message
    message.value = result.message
    messageType.value = 'error'
    password.value = '' //clear password field
  }

  //set loading to false to enable button
  loading.value = false
}
</script>
