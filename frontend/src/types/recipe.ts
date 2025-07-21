export interface Recipe {
  id: string
  title: string
  ingredients: string[]
  steps: string[]
  duration: number
  difficulty: 'easy' | 'medium' | 'hard'
}
