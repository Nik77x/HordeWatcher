<template>
  <p>Performance:</p>
  <p class="perf-text">{{ performance }}</p>
</template>

<script lang="ts">
export default {
  name: 'PerfCI'
}
</script>

<script setup lang="ts">
import { computed, ref } from 'vue'

const props = defineProps({
  perfValue: { type: String, required: true },
  workerType: { type: String, required: true }
})

const performance = computed(() => {
  const match = /[0-9.]/
  const perfOut = match.exec(props.perfValue)
  if (!perfOut) return -1

  const num = parseFloat(perfOut[0])
  if (isNaN(num)) {
    return perfOut
  }

  SetPerfColor(num)
  const perfText = props.workerType === 'text' ? ' tokens / sec' : ' MPS / sec'
  return perfOut + perfText
})

let perfValColor = ref('')

function SetPerfColor(val: number) {
  if (val < 1) perfValColor.value = '#bd0d0d'
  else if (val < 7) perfValColor.value = '#e0e000'
  else perfValColor.value = '#00ad00'
}
</script>

<style scoped>
.perf-text {
  color: v-bind(perfValColor);
}
</style>