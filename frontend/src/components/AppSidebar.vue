<template>
  <aside class="app-sidebar">
    <h3>Navigation</h3>
    <ul>
      <li><router-link to="/">Dashboard</router-link></li>
      <li><router-link to="/profile">Profile</router-link></li>
      <li><router-link to="/settings">Settings</router-link></li>
      <!-- Add more sidebar links as needed -->
    </ul>
  </aside>
</template>

<script setup>
// No script logic needed for a simple sidebar
</script>

<style scoped>
.app-sidebar {
  width: 220px;
  background-color: #f8f9fa;
  padding: 20px;
  border-right: 1px solid #e0e0e0;
  box-shadow: 2px 0 5px rgba(0, 0, 0, 0.05);
  flex-shrink: 0; /* Prevent sidebar from shrinking */
}

.app-sidebar h3 {
  margin-top: 0;
  color: #333;
  border-bottom: 1px solid #eee;
  padding-bottom: 10px;
  margin-bottom: 15px;
}

.app-sidebar ul {
  list-style: none;
  padding: 0;
  margin: 0;
}

.app-sidebar li {
  margin-bottom: 10px;
}

.app-sidebar a {
  text-decoration: none;
  color: #333;
  display: block;
  padding: 8px 10px;
  border-radius: 4px;
  transition:
    background-color 0.3s ease,
    color 0.3s ease;
}

.app-sidebar a:hover,
.app-sidebar a.router-link-active {
  background-color: #e9ecef;
  color: #007bff; /* Bootstrap primary blue */
}
</style>
