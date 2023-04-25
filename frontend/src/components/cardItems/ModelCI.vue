<template>
  <p>Model:</p>
  <button @click="openModelPage()">
    <p class="model-text">{{ modelNameNoPrefix }}</p>
  </button>
</template>

<script lang="ts">
export default {
  name: 'ModelCI'
}
</script>

<script setup lang="ts">
import { BrowserOpenURL } from '@/wailsjs/runtime'
import { computed } from 'vue'

const props = defineProps({
  modelName: { type: String, required: true },
  expanded: { type: Boolean, required: true }
})

const modelNameNoPrefix = computed(() => {
  const name = props.modelName.split('/')
  if (name.length > 1) {
    return name[1]
  }
  return name[0]
})

function openModelPage() {
  BrowserOpenURL('https://huggingface.co/search/full-text?q=' + props.modelName)
}
</script>

<style scoped lang="scss">
button {
  background-color: #ffffff00;
  border: 2px;
  border-radius: 50px;
  transition: background-color 0.2s;
  align-self: start;

  max-width: 98%;
  position: relative;
  right: 7px;

  &:hover {
    background-color: #ffffff11;
  }
}

.model-text {
  color: #2978ff;
  text-overflow: ellipsis;
  text-align: center;
  align-self: start;
  overflow: hidden;
  white-space: v-bind("props.expanded ? 'break-spaces' : 'nowrap'");
  font-size: 15px;
  margin: 3px;
}
</style>