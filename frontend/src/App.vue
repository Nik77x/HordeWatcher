<template>
  <main class='root'>
    <Sidebar></Sidebar>
    <div class='main-content'>
      <WorkerCard v-for='wi in workerInfos' :key='wi.id' :worker-info='wi'></WorkerCard>


    </div>
  </main>
</template>


<script lang='ts' setup>
import { ref } from 'vue'
import { EventsOn } from '@/wailsjs/runtime'
import type { Data } from './wailsjs/go/models'
import Sidebar from '@/components/Sidebar.vue'
import WorkerCard from '@/components/WorkerCard.vue'

let workerInfos = ref<Data.WorkerInfo[]>([])

EventsOn('update_list', (wi: Data.WorkerInfo[]) => {

  workerInfos.value = wi

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
  padding-bottom: 15px;
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
