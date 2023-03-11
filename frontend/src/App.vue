<template>
  <main class='root'>
    <Sidebar></Sidebar>
    <TransitionGroup class='main-content' name='list' tag='div'>
      <WorkerCard v-for='wi in sortedWorkers' :key='wi.id' :worker-data='wi'></WorkerCard>
    </TransitionGroup>
  </main>
</template>


<script lang='ts' setup>
import type { AIWorker } from '@/common/AIWorker'
import { computed, ref } from 'vue'
import { EventsOn, LogDebug } from '@/wailsjs/runtime'
import Sidebar from '@/components/Sidebar.vue'
import WorkerCard from '@/components/WorkerCard.vue'
import Sorter from '@/sorting/Sorter'
import { useSortingModeStore } from '@/Store/Store'
import type { data } from '@/wailsjs/go/models'

const workers = ref<AIWorker[]>([])
const sortingStore = useSortingModeStore()
const sorter = new Sorter()

const sortedWorkers = computed(() => {
  if (workers.value)
    return sorter.sort([...workers.value], sortingStore.sortingMode, sortingStore.invert)
})

EventsOn('update_list', (textWorkers: data.TextWorker[], imageWorkers: data.ImageWorker[]) => {
  let time = Date.now()
  workers.value = [] as AIWorker[]
  workers.value.push(...textWorkers)
  workers.value.push(...imageWorkers)
  time = Date.now() - time
  LogDebug(`List refresh took: ${time}ms`)

})


</script>


<style scoped>


.list-move, /* apply transition to moving elements */
.list-enter-active,
.list-leave-active {
  transition: all 0.5s ease;
}

.list-enter-from,
.list-leave-to {
  opacity: 0;
  transform: translateX(30px);
}

/* ensure leaving items are taken out of layout flow so that moving
   animations can be calculated correctly. */
.list-leave-active {
  position: absolute;
}

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
