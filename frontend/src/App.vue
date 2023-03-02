<template>
  <main class='root'>
    <Sidebar></Sidebar>
    <div class='main-content'>
      <div class='block'></div>
      <div class='block'></div>
      <div class='block'></div>
      <div class='block'></div>
      <div class='block'></div>
      <div class='block'></div>
      <div class='block'></div>
      <div class='block'></div>
      <div class='block'></div>
      <div class='block'></div>
      <div class='block'></div>
      <div class='block'></div>


    </div>
  </main>
</template>


<script lang='ts' setup>
import { ref } from 'vue'
import { EventsOn } from '@/wailsjs/runtime'
import type { Data } from './wailsjs/go/models'
import Sidebar from '@/components/Sidebar.vue'

let name = ref<string>()

EventsOn('update_list', (wi: Data.WorkerInfo[]) => {

  name.value = wi[0].name

})


</script>


<style scoped>

.root {
  display: flex;
  flex-flow: row;
  background-color: black;
  margin: 0;
  padding: 0;
}

.main-content {

  height: 100%;
  background-color: #181818;
  display: flex;
  flex-flow: row;
  flex-flow: row;
  flex-wrap: wrap;
  align-content: start;
  overflow: auto;
}

.block {
  width: 250px;
  height: 170px;
  margin: 5px;
  border-radius: 5px;
  background-color: #282828;
}

header {
  line-height: 1.5;
}

.logo {
  display: block;
  margin: 0 auto 2rem;
}


@media (min-width: 1024px) {
  header {
    display: flex;
    place-items: center;
    padding-right: calc(var(--section-gap) / 2);
  }

  .logo {
    margin: 0 2rem 0 0;
  }

  header .wrapper {
    display: flex;
    place-items: flex-start;
    flex-wrap: wrap;
  }
}
</style>
