<template>

  <button class='open-button' @click='toggleMenu'>
    <font-awesome-icon icon='fa-solid fa-bars-staggered' inverse transform='grow-40' />
  </button>
  <transition>
    <div class='sort-menu' v-if='isMenuOpen' v-on:mouseleave='closeMenu' v-on:mouseenter='keepOpen'>
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
        <label class='chk-label'>
          <input type='checkbox' v-model='sortingStore.invert'>
          Invert
        </label>
    </span>
    </div>
  </transition>

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

let doClose: boolean = false

function toggleMenu() {
  isMenuOpen.value = !isMenuOpen.value
}

function closeMenu() {
  doClose = true
  setTimeout(() => {
    if (doClose)
      isMenuOpen.value = false
  }, 600)

}

function keepOpen() {
  doClose = false
}


</script>

<style scoped lang='scss'>


input[type="checkbox"]:checked {
  background-color: red;
}


.sort-menu {
  position: absolute;
  left: 75px;
  top: 5px;

  width: 156px;
  border-radius: 12px;
  height: 110px;
  background-color: #181818;
  box-shadow: 0 0 20px black;
  border: solid 2px #1d4086;
  vertical-align: center;

  z-index: 0;
  padding: 10px;

  display: flex;
  flex-flow: column;

}

.v-enter-active,
.v-leave-active {
  transition: all 0.3s ease;
  pointer-events: none
}

.v-enter-from,
.v-leave-to {

  opacity: 0;
  transform: translateX(-50px);

}

p {
  font-size: 15px;
}

span {
  display: flex;
  flex-flow: row;

  input[type="checkbox"] {
    margin-right: 5px;

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