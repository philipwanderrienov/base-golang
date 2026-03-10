<script setup>
import { RouterLink, RouterView } from 'vue-router'
import HelloWorld from './components/HelloWorld.vue'
import { ref } from 'vue'
import DataTable from 'primevue/datatable'
import Column from 'primevue/column'
import Button from 'primevue/button'
import Tag from 'primevue/tag'

// Mock data (This will eventually come from your Golang API)
const users = ref([
  { id: 1, name: '', location: '' },
  { id: 2, name: '', location: '' },
  { id: 3, name: '', location: '' },
])

const getSeverity = (role) => {
  switch (role) {
    case 'Admin':
      return 'success'
    case 'Editor':
      return 'info'
    case 'User':
      return 'secondary'
    default:
      return null
  }
}

const fetchFromGo = async () => {
  console.log('Calling your Golang backend...')
  // Example: const response = await fetch('http://localhost:8080/api/users');

  try {
    const response = await fetch('http://localhost:8080/api/v1/congregations')
    const result = await response.json()

    // We use result.data because your Go struct says `json:"data"`
    if (result && result.data) {
        users.value = result.data; 
    } else {
        users.value = [];
    }
  } catch (error) {
    console.error('API Error:', error)
    users.value = [] // Fallback to empty array so the UI doesn't crash
  }
}
</script>

<!-- <template>
  <header>
    <img alt="Vue logo" class="logo" src="@/assets/logo.svg" width="125" height="125" />

    <div class="wrapper">
      <HelloWorld msg="You did it!" />

      <nav>
        <RouterLink to="/">Home</RouterLink>
        <RouterLink to="/about">About</RouterLink>
      </nav>
    </div>
  </header>

  <RouterView />
</template> -->

<template>
  <div class="p-4 md:p-10 bg-surface-50 dark:bg-surface-950 min-h-screen">
    <div class="max-w-6xl mx-auto">
      <header class="mb-8 flex justify-between items-center">
        <h1 class="text-2xl font-bold text-surface-900 dark:text-surface-0">User Management</h1>
        <Button label="Sync with Go" icon="pi pi-sync" @click="fetchFromGo" />
      </header>

      <DataTable
        :value="users"
        paginator
        :rows="5"
        dataKey="id"
        responsiveLayout="stack"
        breakpoint="960px"
        class="shadow-xl rounded-lg overflow-hidden"
      >
        <Column field="id" header="ID" sortable></Column>

        <Column field="name" header="Name" sortable></Column>

        <!-- <Column field="role" header="Role">
          <template #body="slotProps">
            <Tag :value="slotProps.data.role" :severity="getSeverity(slotProps.data.role)" />
          </template>
        </Column> -->

        <!-- <Column header="Actions">
          <template #body>
            <div class="flex gap-2">
              <Button icon="pi pi-pencil" text rounded />
              <Button icon="pi pi-trash" text rounded severity="danger" />
            </div>
          </template>
        </Column> -->

        <Column field="location" header="Location" sortable> </Column>
      </DataTable>
    </div>
  </div>
</template>

<style scoped>
header {
  line-height: 1.5;
  max-height: 100vh;
}

.logo {
  display: block;
  margin: 0 auto 2rem;
}

nav {
  width: 100%;
  font-size: 12px;
  text-align: center;
  margin-top: 2rem;
}

nav a.router-link-exact-active {
  color: var(--color-text);
}

nav a.router-link-exact-active:hover {
  background-color: transparent;
}

nav a {
  display: inline-block;
  padding: 0 1rem;
  border-left: 1px solid var(--color-border);
}

nav a:first-of-type {
  border: 0;
}

@media (min-width: 1024px) {
  header {
    display: flex;
    place-items: center;
    padding-right: calc(var(--section-gap) / 2);
  }

  .logo {
    margin: 0 2rem 0 0;
  }

  header .wrapper {
    display: flex;
    place-items: flex-start;
    flex-wrap: wrap;
  }

  nav {
    text-align: left;
    margin-left: -1rem;
    font-size: 1rem;

    padding: 1rem 0;
    margin-top: 1rem;
  }
}
</style>
