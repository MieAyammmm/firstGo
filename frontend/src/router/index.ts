// src/router/index.ts
import { createRouter, createWebHistory } from 'vue-router'
import LandingPage from '@/views/LandingPage.vue'
import HomePage from '@/views/HomePage.vue'
import AddRecipePage from '@/views/AddRecipePage.vue'
import RecipeDetailPage from '@/views/RecipeDetailPage.vue'
import EditRecipePage from '@/views/EditRecipePage.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'landing',
      component: LandingPage,
    },
    {
      path: '/home',
      name: 'home',
      component: HomePage,
    },
    {
      path: '/add-recipe',
      name: 'add-recipe',
      component: AddRecipePage,
    },
    {
      path: '/recipe/:id',
      name: 'recipe-detail',
      component: RecipeDetailPage,
      props: true,
    },
    {
      path: '/edit-recipe/:id',
      name: 'edit-recipe',
      component: EditRecipePage,
      props: true,
    },
  ],
})

export default router
