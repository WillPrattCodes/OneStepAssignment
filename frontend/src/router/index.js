//import router
import { createRouter, createWebHistory } from 'vue-router';
//import views
import RegisterUser from '@/views/RegisterUser.vue';
import LoginUser from '@/views/LoginUser.vue';
import DashBoard from '@/views/DashBoard.vue';
import GoogleMap from '@/components/GoogleMap.vue';

//declare routes
const routes = [
  { path: '/', component: DashBoard, meta: { requiresAuth: true }},
  { path: '/register', name: 'Register', component: RegisterUser },
  { path: '/login', name: 'Login', component: LoginUser },
  { path: '/map', name: 'Map', component: GoogleMap, meta: { requiresAuth: true }},
];

//create router
const router = createRouter({
  history: createWebHistory(),
  routes,
});

//nav guard for protected routes
router.beforeEach((to, from, next) => {
  //check for auth token from local storage
  const isAuthenticated = localStorage.getItem('token');
  //if route requires auth and user is not authenticated redirect to login
  if ((to.path === '/login' || to.path === '/register') && isAuthenticated) {
    next('/');
  } else if (to.meta.requiresAuth && !isAuthenticated) {
    next('/login');
  } else {
    next();
  }
});

export default router;
