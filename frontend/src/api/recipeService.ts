import apiClient from './index'
import type { Recipe } from '@/types/recipe'

export const RecipeService = {
  async createRecipe(recipe: Omit<Recipe, 'id'>): Promise<Recipe> {
    const response = await apiClient.post('/recipes', recipe)
    return response.data
  },

  async getRecipes(): Promise<Recipe[]> {
    const response = await apiClient.get('/recipes')
    return response.data
  },

  async getRecipeById(id: string): Promise<Recipe> {
    const response = await apiClient.get(`/recipes/${id}`)
    return response.data
  },

  async updateRecipe(id: string, recipe: Partial<Recipe>): Promise<void> {
    await apiClient.put(`/recipes/${id}`, recipe)
  },

  async deleteRecipe(id: string): Promise<void> {
    await apiClient.delete(`/recipes/${id}`)
  },
}
