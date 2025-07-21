<template>
  <div class="edit-recipe-page">
    <h1>Edit Recipe</h1>

    <form @submit.prevent="submitForm" v-if="recipe">
      <div class="mb-3">
        <label for="title" class="form-label">Title</label>
        <input v-model="recipe.title" type="text" class="form-control" id="title" required />
      </div>

      <div class="mb-3">
        <label class="form-label">Ingredients</label>
        <div
          v-for="(ingredient, index) in recipe.ingredients"
          :key="index"
          class="input-group mb-2"
        >
          <input v-model="recipe.ingredients[index]" type="text" class="form-control" required />
          <button @click="removeIngredient(index)" type="button" class="btn btn-outline-danger">
            Remove
          </button>
        </div>
        <button @click="addIngredient" type="button" class="btn btn-outline-secondary">
          Add Ingredient
        </button>
      </div>

      <div class="mb-3">
        <label class="form-label">Steps</label>
        <div v-for="(step, index) in recipe.steps" :key="index" class="input-group mb-2">
          <textarea v-model="recipe.steps[index]" class="form-control" rows="2" required></textarea>
          <button @click="removeStep(index)" type="button" class="btn btn-outline-danger">
            Remove
          </button>
        </div>
        <button @click="addStep" type="button" class="btn btn-outline-secondary">Add Step</button>
      </div>

      <div class="mb-3">
        <label for="duration" class="form-label">Duration (minutes)</label>
        <input
          v-model.number="recipe.duration"
          type="number"
          class="form-control"
          id="duration"
          required
        />
      </div>

      <div class="mb-3">
        <label for="difficulty" class="form-label">Difficulty</label>
        <select v-model="recipe.difficulty" class="form-select" id="difficulty" required>
          <option value="easy">Easy</option>
          <option value="medium">Medium</option>
          <option value="hard">Hard</option>
        </select>
      </div>

      <button type="submit" class="btn btn-primary" :disabled="submitting">
        <span
          v-if="submitting"
          class="spinner-border spinner-border-sm"
          role="status"
          aria-hidden="true"
        ></span>
        {{ submitting ? 'Updating...' : 'Update Recipe' }}
      </button>
    </form>

    <div v-else-if="loading" class="text-center">
      <div class="spinner-border" role="status">
        <span class="visually-hidden">Loading...</span>
      </div>
    </div>

    <div v-else class="alert alert-danger">Recipe not found</div>
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted, reactive, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { RecipeService } from '@/api/recipeService'
import type { Recipe } from '@/types/recipe'

export default defineComponent({
  name: 'EditRecipePage',
  setup() {
    const route = useRoute()
    const router = useRouter()
    const recipe = ref<Recipe | null>(null)
    const loading = ref(true)
    const submitting = ref(false)

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

    const addIngredient = () => {
      if (recipe.value) {
        recipe.value.ingredients.push('')
      }
    }

    const removeIngredient = (index: number) => {
      if (recipe.value && recipe.value.ingredients.length > 1) {
        recipe.value.ingredients.splice(index, 1)
      }
    }

    const addStep = () => {
      if (recipe.value) {
        recipe.value.steps.push('')
      }
    }

    const removeStep = (index: number) => {
      if (recipe.value && recipe.value.steps.length > 1) {
        recipe.value.steps.splice(index, 1)
      }
    }

    const submitForm = async () => {
      if (!recipe.value) return

      submitting.value = true
      try {
        await RecipeService.updateRecipe(recipe.value.id, recipe.value)
        router.push(`/recipe/${recipe.value.id}`)
      } catch (error) {
        console.error('Error updating recipe:', error)
      } finally {
        submitting.value = false
      }
    }

    onMounted(fetchRecipe)

    return {
      recipe,
      loading,
      submitting,
      addIngredient,
      removeIngredient,
      addStep,
      removeStep,
      submitForm,
    }
  },
})
</script>
