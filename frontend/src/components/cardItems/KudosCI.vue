<template>
  <p>Kudos earned:</p>
  <p>{{ kudos }}</p>
</template>

<script lang="ts">
export default {
  name: 'KudosCI'
}
</script>

<script setup lang="ts">
const props = defineProps({
  kudos: { type: Number, required: true }
})
</script>

<style scoped></style>