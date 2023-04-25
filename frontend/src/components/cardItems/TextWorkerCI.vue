<template>
  <p>Max generation length:</p>
  <p>{{ maxGenLength }} tokens</p>
  <p>Max context length:</p>
  <p>{{ maxCtxLength }} tokens</p>
</template>

<script lang="ts">
export default {
  name: 'TextWorkerCI'
}
</script>

<script setup lang="ts">
const props = defineProps({
  maxGenLength: { type: Number, required: true },
  maxCtxLength: { type: Number, required: true }
})
</script>

<style scoped></style>