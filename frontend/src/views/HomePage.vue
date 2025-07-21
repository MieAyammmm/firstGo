<template>
  <div class="home-page">
    <h1>All Recipes</h1>
    <router-link to="/add-recipe" class="btn btn-success mb-4"> Add New Recipe </router-link>

    <div v-if="loading" class="text-center">
      <div class="spinner-border" role="status">
        <span class="visually-hidden">Loading...</span>
      </div>
    </div>

    <div v-else class="recipe-list">
      <div v-for="recipe in recipes" :key="recipe.id" class="recipe-card">
        <div class="card">
          <div class="card-body">
            <h5 class="card-title">{{ recipe.title }}</h5>
            <p class="card-text">
              Difficulty: {{ recipe.difficulty }} | Duration: {{ recipe.duration }} mins
            </p>
            <router-link :to="`/recipe/${recipe.id}`" class="btn btn-primary">
              View Details
            </router-link>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted, ref } from 'vue'
import { RecipeService } from '@/api/recipeService'
import type { Recipe } from '@/types/recipe'

export default defineComponent({
  name: 'HomePage',
  setup() {
    const recipes = ref<Recipe[]>([])
    const loading = ref(true)

    const fetchRecipes = async () => {
      try {
        recipes.value = await RecipeService.getRecipes()
      } catch (error) {
        console.error('Error fetching recipes:', error)
      } finally {
        loading.value = false
      }
    }

    onMounted(fetchRecipes)

    return {
      recipes,
      loading,
    }
  },
})
</script>

<style scoped>
.recipe-list {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  gap: 1rem;
}
</style>
