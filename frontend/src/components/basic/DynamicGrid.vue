<template>
  <div class="content-grid">
    <slot></slot>
  </div>
</template>

<script lang="ts">
export default {
  name: 'DynamicGrid'
}
</script>

<script setup lang="ts">
import { ref } from 'vue'

const props = defineProps({
  rows: { type: Number, required: true },
  columns: { type: Number, required: true },
  rowCss: { type: String, required: false },
  columnCss: { type: String, required: false }
})

const rowCss = ref('1fr')

function getCustomCss(css: string | undefined, defaultCss: string): string {
  if (css) {
    return css
  }
  return defaultCss
}
</script>

<style scoped lang="scss">
.content-grid {
  display: grid;
  margin: 3px;
  height: 100%;

  grid-template-columns: repeat(
    v-bind(columns),
    v-bind('getCustomCss(columnCss, "minmax(130px, 1fr) minmax(0, 1.5fr)")')
  );
  grid-template-rows: repeat(v-bind(rows), v-bind('getCustomCss(rowCss, "minmax(20px, 1fr)")'));
  align-content: center;
  justify-items: left;
  overflow: hidden;

  font-size: 15px;
}
</style>