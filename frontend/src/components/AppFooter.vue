<template>
  <footer class="app-footer">
    <p>&copy; {{ currentYear }} My SPA. All rights reserved.</p>
  </footer>
</template>

<script setup>
import { ref } from 'vue'

const currentYear = ref(new Date().getFullYear())
</script>

<style scoped>
.app-footer {
  background-color: #2c3e50;
  color: white;
  padding: 15px 20px;
  text-align: center;
  margin-top: auto; /* Pushes footer to the bottom if content is short */
  box-shadow: 0 -2px 4px rgba(0, 0, 0, 0.1);
}

.app-footer p {
  margin: 0;
  font-size: 0.9em;
}
</style>
