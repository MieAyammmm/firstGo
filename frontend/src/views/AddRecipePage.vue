<template>
  <div class="add-recipe-page">
    <h1>Add New Recipe</h1>

    <form @submit.prevent="submitForm">
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
        {{ submitting ? 'Submitting...' : 'Submit' }}
      </button>
    </form>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, ref } from 'vue'
import { RecipeService } from '@/api/recipeService'
import { useRouter } from 'vue-router'
import type { Recipe } from '@/types/recipe'

export default defineComponent({
  name: 'AddRecipePage',
  setup() {
    const router = useRouter()
    const submitting = ref(false)

    const recipe = reactive<Omit<Recipe, 'id'>>({
      title: '',
      ingredients: [''],
      steps: [''],
      duration: 0,
      difficulty: 'easy',
    })

    const addIngredient = () => {
      recipe.ingredients.push('')
    }

    const removeIngredient = (index: number) => {
      if (recipe.ingredients.length > 1) {
        recipe.ingredients.splice(index, 1)
      }
    }

    const addStep = () => {
      recipe.steps.push('')
    }

    const removeStep = (index: number) => {
      if (recipe.steps.length > 1) {
        recipe.steps.splice(index, 1)
      }
    }

    const submitForm = async () => {
      submitting.value = true
      try {
        await RecipeService.createRecipe(recipe)
        router.push('/home')
      } catch (error) {
        console.error('Error creating recipe:', error)
      } finally {
        submitting.value = false
      }
    }

    return {
      recipe,
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
