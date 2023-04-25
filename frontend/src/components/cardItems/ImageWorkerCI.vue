<template>
  <p>Max image size:</p>
  <p>{{ maxImSize }}x</p>
  <p>Generated MPS:</p>
  <p>{{ mpsGenerated }}</p>
</template>

<script lang="ts">
export default {
  name: 'ImageWorkerCI'
}
</script>

<script setup lang="ts">
import { computed } from 'vue'

const props = defineProps({
  maxPixels: { type: Number, required: true },
  mpsGenerated: { type: Number, required: true }
})

const maxImSize = computed(() => {
  // max square image that can be generated
  return Math.floor(Math.sqrt(props.maxPixels))
})
</script>

<style scoped></style>