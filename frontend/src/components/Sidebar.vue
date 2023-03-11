<template>

  <div class='sb-content'>
    <sorting-mode-selector></sorting-mode-selector>
  </div>
</template>

<script lang='ts'>
export default {
  name: 'Sidebar'
}
</script>

<script lang='ts' setup>


import SortingModeSelector from '@/components/SortingModeSelector.vue'
</script>

<style scoped lang='scss'>

:deep(button) {

  width: 100%;
  background-color: transparent;
  border: transparent;
  aspect-ratio: 1;

  transition: background-color 0.2s;

  &.btn-settings {
    margin: auto 0 0;
  }

  &:hover {

    background-color: #ffffff11;

  }
}


.icon {
  width: 100%;
  aspect-ratio: 1;
  margin: 0;
}

.sb-content {
  background-color: #303030;

  display: flex;
  flex-flow: column;
  justify-content: start;
  align-items: start;


  min-width: 70px;
  width: 70px;
  max-width: 70px;

  z-index: 5;

}

</style>