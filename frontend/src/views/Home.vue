<script setup lang="ts">
import { ref, onMounted } from 'vue'

interface SystemInfo {
  hostname: string
  os: string
  arch: string
  numCPU: number
}

const systemInfo = ref<SystemInfo | null>(null)
const loading = ref(true)

onMounted(async () => {
  try {
    // 通过 Wails v3 生成的绑定调用后端 Service 方法
    // 绑定文件由 `wails3 generate bindings` 自动生成
    const { GetSystemInfo } = await import('../../bindings/cleanC/services/system/systemservice.js')
    systemInfo.value = await GetSystemInfo()
  } catch (e) {
    console.error('获取系统信息失败:', e)
  } finally {
    loading.value = false
  }
})
</script>

<template>
  <div class="home">
    <h2>系统信息</h2>
    <div v-if="loading" class="loading">加载中...</div>
    <div v-else-if="systemInfo" class="info-grid">
      <div class="info-item">
        <span class="label">主机名</span>
        <span class="value">{{ systemInfo.hostname }}</span>
      </div>
      <div class="info-item">
        <span class="label">操作系统</span>
        <span class="value">{{ systemInfo.os }}</span>
      </div>
      <div class="info-item">
        <span class="label">架构</span>
        <span class="value">{{ systemInfo.arch }}</span>
      </div>
      <div class="info-item">
        <span class="label">CPU 核心数</span>
        <span class="value">{{ systemInfo.numCPU }}</span>
      </div>
    </div>
    <div v-else class="error">无法获取系统信息</div>
  </div>
</template>

<style scoped>
.home h2 {
  margin-bottom: 1rem;
  color: #2c3e50;
}

.info-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 1rem;
}

.info-item {
  padding: 1rem;
  background: #f8f9fa;
  border-radius: 6px;
}

.info-item .label {
  display: block;
  font-size: 0.8rem;
  color: #7f8c8d;
  margin-bottom: 0.25rem;
}

.info-item .value {
  font-size: 1.1rem;
  font-weight: 600;
  color: #2c3e50;
}

.loading {
  color: #7f8c8d;
}

.error {
  color: #e74c3c;
}
</style>
