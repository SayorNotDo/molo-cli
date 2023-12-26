import { createRouter, createWebHistory } from 'vue-router';

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      component: () => import('../layout/login/LoginLayout.vue'),
      redirect: '/login',
      children: [
        {
          path: 'login',
          name: 'login',
          component: () => import('../pages/Login.vue')
        }
      ]
    }
  ]
});

export default router;