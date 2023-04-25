<template>
  <p>Queue:</p>
  <p class="queue-value">{{ Queue }}</p>
</template>

<script lang="ts">
export default {
  name: 'QueueCI'
}
</script>

<script setup lang="ts">
import { computed, ref } from 'vue'

const props = defineProps({
  queue: { type: Number, required: true },
  workerType: { type: String, required: true }
})

const Queue = computed(() => {
  setQueueColor(props.queue)
  if (props.workerType == 'text') {
    return props.queue.toFixed() + ' tokens'
  }
  return props.queue.toFixed(2) + ' MPS'
})

const queueColor = ref('')

function setQueueColor(queue: number) {
  if (queue <= 20) queueColor.value = '#b1b1c2'
  else if (queue <= 80) queueColor.value = '#e0e000'
  else if (queue <= 200) queueColor.value = '#d97900'
  else queueColor.value = '#dc1010'
}
</script>

<style scoped>
.queue-value {
  color: v-bind(queueColor);
}
</style>