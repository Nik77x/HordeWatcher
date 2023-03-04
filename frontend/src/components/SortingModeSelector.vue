<template>
  <button class='open-button' @click='toggleMenu'>
    <font-awesome-icon icon='fa-solid fa-bars-staggered' inverse transform='grow-40' />
  </button>

  <div class='sort-menu' v-if='isMenuOpen' v-on:mouseleave='toggleMenu'>
    <p>Sorting Mode: </p>
    <select class='selector' v-model='sortingStore.sortingMode'>
      <option :value='SortingMode.Performance'>Performance</option>
      <option :value='SortingMode.Uptime'>Uptime</option>
      <option :value='SortingMode.GenLength'>GenLength</option>
      <option :value='SortingMode.ContextSize'>ContextSize</option>
      <option :value='SortingMode.Jobs'>Jobs</option>
      <option :value='SortingMode.Kudos'>Kudos</option>
    </select>
    <span>
      <input type='checkbox' v-model='sortingStore.invert'>
      <p>Invert</p>
    </span>
  </div>

</template>

<script lang='ts'>
export default {
  name: 'SortingModeSelector'
}
</script>

<script setup lang='ts'>
import { useSortingModeStore } from '@/Store/Store'
import { SortingMode } from '@/sorting/SortingMode'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import { ref } from 'vue'

const sortingStore = useSortingModeStore()

const isMenuOpen = ref(false)

function toggleMenu() {
  isMenuOpen.value = !isMenuOpen.value
}


</script>

<style scoped lang='scss'>

.sort-menu {
  position: absolute;
  left: 75px;
  top: 5px;

  z-index: 10;
  width: 156px;
  border-radius: 12px;
  height: 110px;
  background-color: #181818;
  box-shadow: 0 0 20px black;
  border: solid 2px #1d4086;
  vertical-align: center;

  padding: 10px;

  display: flex;
  flex-flow: column;

}

p {
  font-size: 15px;
}

span {
  display: flex;
  flex-flow: row;

  input[type="checkbox"] {
    margin-right: 10px;

    &:checked {
      background-color: #3473ef;
    }
  }

}

.selector {
  width: 100%;
  height: 30px;
  background-color: #282828;
  border: solid 1px #3473ef;
  border-radius: 5px;
  color: #e1e1e1;
  padding: 5px;
  margin-top: 5px;
  margin-bottom: 5px;


  :nth-child(even) {
    background-color: #212121;
  }

  :nth-child(odd) {
    background-color: #252525;
  }


}

option {
  color: #919191;
  font-size: 15px;
  border-radius: 0;
  border: solid 1px #3473ef;

}

select {

}

.select-items {
  padding: 10px;

}

.select-selected {
  background-color: #3473ef;

}
</style>