<template>
  <div class="recipe-detail-page">
    <div v-if="loading" class="text-center">
      <div class="spinner-border" role="status">
        <span class="visually-hidden">Loading...</span>
      </div>
    </div>

    <div v-else-if="recipe" class="recipe-container">
      <h1>{{ recipe.title }}</h1>
      <p class="recipe-meta">
        Difficulty: {{ recipe.difficulty }} | Duration: {{ recipe.duration }} minutes
      </p>

      <div class="recipe-section">
        <h2>Ingredients</h2>
        <ul>
          <li v-for="(ingredient, index) in recipe.ingredients" :key="index">
            {{ ingredient }}
          </li>
        </ul>
      </div>

      <div class="recipe-section">
        <h2>Steps</h2>
        <ol>
          <li v-for="(step, index) in recipe.steps" :key="index">
            {{ step }}
          </li>
        </ol>
      </div>

      <div class="recipe-actions">
        <button @click="editRecipe" class="btn btn-warning me-2">Edit Recipe</button>
        <button @click="deleteRecipe" class="btn btn-danger">Delete Recipe</button>
      </div>
    </div>

    <div v-else class="alert alert-danger">Recipe not found</div>
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { RecipeService } from '@/api/recipeService'
import type { Recipe } from '@/types/recipe'

export default defineComponent({
  name: 'RecipeDetailPage',
  setup() {
    const route = useRoute()
    const router = useRouter()
    const recipe = ref<Recipe | null>(null)
    const loading = ref(true)

    const fetchRecipe = async () => {
      try {
        const recipeData = await RecipeService.getRecipeById(route.params.id as string)
        recipe.value = recipeData
      } catch (error) {
        console.error('Error fetching recipe:', error)
      } finally {
        loading.value = false
      }
    }

    const editRecipe = () => {
      if (recipe.value) {
        router.push(`/edit-recipe/${recipe.value.id}`)
      }
    }

    const deleteRecipe = async () => {
      if (recipe.value && confirm('Are you sure you want to delete this recipe?')) {
        try {
          await RecipeService.deleteRecipe(recipe.value.id)
          router.push('/home')
        } catch (error) {
          console.error('Error deleting recipe:', error)
        }
      }
    }

    onMounted(fetchRecipe)

    return {
      recipe,
      loading,
      editRecipe,
      deleteRecipe,
    }
  },
})
</script>

<style scoped>
.recipe-meta {
  font-style: italic;
  color: #666;
}

.recipe-section {
  margin-bottom: 2rem;
}

.recipe-actions {
  margin-top: 2rem;
}
</style>
