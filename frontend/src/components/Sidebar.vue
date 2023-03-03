<template>

  <div class='sb-content'>
    <button>
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


</script>

<style scoped lang='scss'>

button {

  width: 100%;
  background-color: transparent;
  border: transparent;
  aspect-ratio: 1;
  margin: auto 0 0;
  transition: background-color 0.2s;

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
  align-items: start;


  min-width: 70px;
  width: 70px;
  max-width: 70px;

}

</style>