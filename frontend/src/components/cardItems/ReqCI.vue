<template>
  <p>Requests Fulfilled:</p>
  <p>{{ reqFulfilled }}</p>
</template>

<script lang="ts">
export default {
  name: 'ReqCI'
}
</script>

<script setup lang="ts">
const props = defineProps({
  reqFulfilled: { type: Number, required: true }
})
</script>

<style scoped></style>