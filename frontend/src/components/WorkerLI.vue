<template>
  <li>
    <div class='main-content'>
      <h1>heh</h1>
    </div>
  </li>
</template>

<script setup lang='ts'>

import type { Data } from '@/wailsjs/go/models'

const props = defineProps<{
  WorkerInfo: Data.WorkerInfo
}>()

</script>

<style scoped>

</style>