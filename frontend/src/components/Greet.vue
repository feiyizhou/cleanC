<script setup lang="ts">
import { ref } from 'vue'

const name = ref('')
const greeting = ref('')

function greet() {
  if (name.value) {
    greeting.value = `你好, ${name.value}!`
  }
}
</script>

<template>
  <div class="greet">
    <input v-model="name" placeholder="输入你的名字" @keyup.enter="greet" />
    <button @click="greet">打招呼</button>
    <p v-if="greeting" class="result">{{ greeting }}</p>
  </div>
</template>

<style scoped>
.greet {
  display: flex;
  gap: 0.5rem;
  align-items: center;
  flex-wrap: wrap;
}

input {
  padding: 0.5rem;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 0.9rem;
}

button {
  padding: 0.5rem 1rem;
  background: #3498db;
  color: #fff;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

button:hover {
  background: #2980b9;
}

.result {
  width: 100%;
  margin-top: 0.5rem;
  color: #2c3e50;
}
</style>
