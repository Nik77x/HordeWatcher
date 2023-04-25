<template>
  <p>Uptime:</p>
  <p>{{ uptime }}</p>
</template>

<script lang="ts">
export default {
  name: 'UptimeCI'
}
</script>

<script setup lang="ts">
import { computed } from 'vue'

const props = defineProps({
  uptime: { type: Number, required: true }
})

const uptime = computed(() => {
  let uptimeMinutes = props.uptime / 60
  if (uptimeMinutes > 60) {
    return (uptimeMinutes / 60).toFixed(1) + ' h'
  }
  return uptimeMinutes.toFixed(1) + ' m'
})
</script>

<style scoped lang="scss"></style>