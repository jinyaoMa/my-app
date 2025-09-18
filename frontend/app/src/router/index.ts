import { createRouter, createWebHashHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import SystrayView from '../views/SystrayView.vue'

const router = createRouter({
  history: createWebHashHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/systray',
      name: 'systray',
      component: SystrayView
    }
    // {
    //   path: '/tmpl',
    //   name: 'tmpl',
    //   // route level code-splitting
    //   // this generates a separate chunk (About.[hash].js) for this route
    //   // which is lazy-loaded when the route is visited.
    //   // component: () => import('../views/TemplateView.vue')
    //   component: TemplateView
    // }
  ]
})

export default router
