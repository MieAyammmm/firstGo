<template>
  <div class="landing-page">
    <h1>Welcome to Recipe App</h1>
    <p>Discover and share your favorite recipes</p>
    <router-link to="/home" class="btn btn-primary"> Explore Recipes </router-link>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'LandingPage',
})
</script>

<style scoped>
.landing-page {
  text-align: center;
  padding: 2rem;
  min-height: 80vh;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}
</style>
