export interface Recipe {
  _id: string
  title: string
  ingredients: string[]
  steps: string[]
  duration: number
  difficulty: 'easy' | 'medium' | 'hard'
}
