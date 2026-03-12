<template>
  <div id="app-wrapper">
    <AppHeader />
    <div class="main-layout-content">
      <AppSidebar />
      <main class="page-content-area">
        <!-- This is where the component for the current route will be rendered -->
        <router-view />
      </main>
    </div>
    <AppFooter />
  </div>
</template>

<script setup>
import AppHeader from './components/AppHeader.vue'
import AppSidebar from './components/AppSidebar.vue'
import AppFooter from './components/AppFooter.vue'
</script>

<style>
/* Global styles for the layout */
html,
body,
#app {
  margin: 0;
  padding: 0;
  height: 100%;
  width: 100%;
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
  display: block; /* Ensures #app is a block container, overriding potential grid styles */
}

#app-wrapper {
  display: flex;
  flex-direction: column;
  min-height: 100vh; /* Ensures the wrapper takes full viewport height */
}

.main-layout-content {
  display: flex;
  flex: 1; /* Allows this section to grow and push the footer down */
}

.page-content-area {
  flex: 1; /* Allows the content area to take up remaining space */
  padding: 20px;
  background-color: #ffffff;
  overflow-y: auto; /* Add scroll for long content */
}
</style>
