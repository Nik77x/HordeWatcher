<template>

  <div class='sb-content'>
    <sorting-mode-selector></sorting-mode-selector>
    <button class='btn-settings'>
      <font-awesome-icon class='icon' icon='fa-solid fa-gear' fixed-width inverse transform='grow-40 rotate-30' />
    </button>
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

}

</style>