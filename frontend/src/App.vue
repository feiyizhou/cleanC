<script setup lang="ts">
import Home from './views/Home.vue'
</script>

<template>
  <div class="app-container">
    <header class="app-header">
      <h1>CleanC</h1>
      <p>Windows 系统工具</p>
    </header>
    <main class="app-main">
      <Home />
    </main>
  </div>
</template>

<style scoped>
.app-container {
  max-width: 960px;
  margin: 0 auto;
  padding: 2rem;
}

.app-header {
  text-align: center;
  margin-bottom: 2rem;
}

.app-header h1 {
  font-size: 2rem;
  color: #2c3e50;
  margin-bottom: 0.5rem;
}

.app-header p {
  color: #7f8c8d;
}

.app-main {
  background: #fff;
  border-radius: 8px;
  padding: 2rem;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.08);
}
</style>
