import axios from 'axios'

const api = axios.create({
  baseURL: import.meta.env.VITE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
})

//on request add token to headers
api.interceptors.request.use((config) => {
  const token = localStorage.getItem('token')
  if (token) {
    config.headers.Authorization = `Bearer ${token}`
  }
  return config
})

//on response check for errors
api.interceptors.response.use(
  (response) => response,
  (error) => {
    //check if the error is related to authentication
    if (error.response && error.response.status === 401) {
      //if token is invalid remove from local storage and redirect to login
      localStorage.removeItem('token')
      window.location.href = '/login'
    }
    return Promise.reject(error)
  },
)

export default api
