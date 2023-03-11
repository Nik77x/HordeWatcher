<template>
  <div class='card' :style='{width: width + "px", height: height + "px"}'>
    <h1 class='title'> {{ props.workerData.name }}</h1>

    <div class='content'>
      <p class='content-text'>model: </p>

      <button @click='openModelPage()' class='btn-model'>
        <p class='model-text'> {{ modelName }} </p>
      </button>


      <p class='content-text'>Queue: </p>
      <p class='content-text'> {{ props.workerData.queue }}</p>
      <p class='content-text'>tokens / sec: </p>
      <p class='performance-value'> {{ performance }} </p>
      <p class='content-text'>Uptime: </p>
      <p class='content-text'> {{ uptime }} </p>


      <template v-if='expanded'>
        <p>Max Generation Length: </p>
        <p>{{ 'no' }} tokens</p>
        <p>Context Size: </p>
        <p>{{ 'no' }} tokens</p>
        <p>Type: </p>
        <p>{{ props.workerData.type }}</p>
        <p>Kudos: </p>
        <p>{{ props.workerData.kudos_rewards }}</p>
        <p>Online: </p>
        <p>{{ props.workerData.online ? 'Yes' : 'No' }}</p>
        <p>Requests fullfilled: </p>
        <p>{{ props.workerData.requests_fulfilled }}</p>
        <p>MMode: </p>
        <p>{{ props.workerData.maintenance_mode }}</p>
        <p></p>
        <p></p>
      </template>


    </div>

    <button class='btn-expand' @click='toggleExpand()'>
      <font-awesome-icon class='arrow' icon='fa-solid fa-chevron-down' inverse transform='rotate--45 grow-1' />
    </button>

  </div>
</template>

<script lang='ts'>
export default {
  name: 'WorkerCard'
}
</script>

<script setup lang='ts'>

import { computed, ref } from 'vue'
import type { AIWorker } from '@/common/AIWorker'
import { BrowserOpenURL } from '@/wailsjs/runtime'

const props = defineProps<{
  workerData: AIWorker
}>()

const modelName = computed(() => {

  const name = props.workerData.models[0].split('/')
  if (name.length > 1) {
    return name[1]
  }
  return name[0]
})

const performance = computed(() => {
  const match = /[0-9.]/
  const perfOut = match.exec(props.workerData.performance)
  if (!perfOut) return -1

  const num = parseFloat(perfOut[0])
  if (isNaN(num)) {
    return perfOut
  }


  SetPerfColor(num)
  const perfValue = props.workerData.type == 'text' ? ' tokens / sec' : 'MPS / sec'
  return perfOut + perfValue


})

const UncompletedJobs = computed(() => {
  if (!props.workerData) return 0
  let jobsNum = props.workerData.uncompleted_jobs
  setJobsColor(jobsNum)
  return jobsNum
})

const uptime = computed(() => {
  if (!props.workerData) return 0
  let uptimeMinutes = (props.workerData.uptime / 60)
  if (uptimeMinutes > 60) {
    return (uptimeMinutes / 60).toFixed(1) + ' h'
  }
  return uptimeMinutes.toFixed(1) + ' m'
})

function openModelPage() {
  if (!props.workerData) return

  else BrowserOpenURL('https://www.google.com/search?q=site:https://huggingface.co/ ' + props.workerData.models[0])
}


let expanded = false
let width = ref(300)
let height = ref(170)

const rows = ref(4)
const columns = ref(2)

function updateGrid() {
  if (expanded) {
    rows.value -= 1
    columns.value /= 2

    setTimeout(() => {
      // compensate for the margin
      width.value = (width.value - 10) / 2
      height.value /= 2
    }, 20)


  } else {

    rows.value += 1
    columns.value *= 2
    setTimeout(() => {
      width.value = (width.value * 2) + 10
      height.value *= 2
    }, 20)
  }
}

function toggleExpand() {
  updateGrid()
  expanded = !expanded
}

// styling functions

let perfValColor = ref('')
let jobsColor = ref('')

// TODO: the colors should be variables placed elsewhere

function SetPerfColor(val: number) {

  if (val < 1) perfValColor.value = '#bd0d0d'
  else if (val < 7) perfValColor.value = '#e0e000'
  else perfValColor.value = '#00ad00'

}

function setJobsColor(jobsNum: number) {
  if (jobsNum <= 2) jobsColor.value = '#b1b1c2'
  else if (jobsNum <= 10) jobsColor.value = '#e0e000'
  else if (jobsNum <= 20) jobsColor.value = '#d97900'
  else jobsColor.value = '#bd0d0d'
}


</script>

<style scoped lang='scss'>

.card {
  margin: 10px 0 0 10px;
  padding: 3px 2px 3px 5px;
  border-radius: 5px;
  background-color: #282828;
  display: flex;
  flex-flow: column;
  position: relative;
  transition-property: height, width;
  transition-duration: 0.3s;
  overflow: hidden;
  box-shadow: 0 3px 7px black;

}

.title {
  width: 100%;
  min-height: 35px;
  color: #e1e1e1;
  margin: 2px;
  vertical-align: top;
  font-size: 20px;
  text-overflow: ellipsis;
  overflow: hidden;
  white-space: nowrap;

}

.content {

  display: grid;
  margin: 3px;
  height: 100%;

  grid-template-columns: repeat(v-bind("columns / 2"), minmax(130px, 1fr) minmax(0, 1.5fr));
  grid-template-rows: repeat(v-bind(rows), minmax(20px, 1fr)); // 1fr 1fr 1fr 1fr; //;
  align-content: center;
  justify-items: left;
  overflow: hidden;

  font-size: 15px;

}

.btn-model {
  max-width: 98%;
  position: relative;
  right: 7px;

}

p {
  align-content: center;
  align-self: center;

  &.model-text {

    color: #2978ff;
    text-overflow: ellipsis;
    text-align: start;
    align-self: start;
    overflow: hidden;
    white-space: v-bind("expanded ? 'break-spaces' : 'nowrap'");
    font-size: 15px;
  }

  &.performance-value {
    color: v-bind(perfValColor);
  }

  &.jobs-text {
    color: v-bind(jobsColor);
  }

  &.content-text {
    overflow: clip;
  }


}

button {
  background-color: #ffffff00;
  border: 2px;
  border-radius: 50px;
  transition: background-color 0.2s;
  align-self: center;


  p {
    margin: 3px;
  }

  &:hover {
    background-color: #ffffff11;
  }
}

.btn-expand {
  position: absolute;

  width: 20px;
  aspect-ratio: 1;
  background-color: transparent;
  right: 5px;
  bottom: 5px;
  align-items: center;
  justify-content: center;
  display: flex;

  .arrow {
    align-self: center;
  }
}

</style>