<template>
  <header class="app-header">
    <div class="header-content">
      <h1>My Awesome SPA</h1>
      <nav>
        <router-link to="/">Home</router-link> |
        <router-link to="/about">About</router-link>
        <!-- Add more navigation links here -->
      </nav>
    </div>
  </header>
</template>

<script setup>
// No script logic needed for a simple header
</script>

<style scoped>
.app-header {
  background-color: #2c3e50;
  color: white;
  padding: 15px 20px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.header-content {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.app-header h1 {
  margin: 0;
  font-size: 1.5em;
}

.app-header nav a {
  color: white;
  margin-left: 20px;
  text-decoration: none;
  font-weight: bold;
  transition: color 0.3s ease;
}

.app-header nav a:hover,
.app-header nav a.router-link-active {
  color: #42b983; /* Vue green */
}
</style>
