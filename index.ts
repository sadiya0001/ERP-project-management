import { createRouter, createWebHistory } from 'vue-router'
import ProjectPage from '../pages/ProjectPage.vue'

const routes = [
  {
    path: '/',
    component: ProjectPage,
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router