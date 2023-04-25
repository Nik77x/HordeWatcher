<template>
  <p>Info:</p>
  <p class="info-text">{{ info }}</p>
</template>

<script lang="ts">
export default {
  name: 'InfoCI'
}
</script>

<script setup lang="ts">
import { computed, ref } from 'vue'

const props = defineProps({
  maybeInfo: { type: String, required: true }
})

const noInfo = ref(false)
const info = computed(() => {
  if (props.maybeInfo == '') {
    noInfo.value = true
    return '*No info provided*'
  }
  noInfo.value = false

  return props.maybeInfo
})
</script>

<style scoped lang="scss">
.info-text {
  font-style: v-bind('noInfo ? "italic" : "normal"');
  color: v-bind('noInfo ? "#777" : "#a4a4a4"');
  grid-column: span 3;
  background-color: #212121;
  border: 1px;
  border-radius: 5px;
  width: 100%;
  padding: 3px 5px;
}
</style>