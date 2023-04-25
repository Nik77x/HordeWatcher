<template>
  <div class="card" :style="{ width: cardWidth + 'px', height: cardHeight + 'px' }">
    <h1 class="title">{{ props.workerData.name }}</h1>

    <DynamicGrid
      :columns="gridColumns / 2"
      :rows="gridRows"
      row-css="minmax(20px, 1fr)"
      column-css="minmax(130px, 1fr) minmax(0, 2fr)"
      class="grid-content"
    >
      <ModelCI :model-name="workerData.models[0]" :expanded="expanded" />
      <QueueCI :queue="workerData.queue" :worker-type="workerData.type" />
      <PerfCI :perf-value="workerData.performance" :worker-type="workerData.type" />
      <UptimeCI :uptime="workerData.uptime" />

      <template v-if="expanded">
        <ReqCI :req-fulfilled="workerData.requests_fulfilled" />
        <KudosCI :kudos="workerData.kudos_rewards" />
        <template v-if="workerData.type === 'text'">
          <TextWorkerCI
            :max-ctx-length="getTextWorker().max_context_length"
            :max-gen-length="getTextWorker().max_length"
          />
        </template>
        <template v-else>
          <ImageWorkerCI
            :mps-generated="getImageWorker().megapixelsteps_generated"
            :max-pixels="getImageWorker().max_pixels"
          />
        </template>
        <InfoCI :maybe-info="workerData.info ? workerData.info : ''" />
      </template>
    </DynamicGrid>

    <button class="btn-expand" @click="toggleExpand()">
      <font-awesome-icon
        class="arrow"
        icon="fa-solid fa-chevron-down"
        inverse
        transform="rotate--45 grow-1"
      />
    </button>
  </div>
</template>

<script lang="ts">
export default {
  name: 'WorkerCard'
}
</script>

<script setup lang="ts">
import { ref } from 'vue'
import type { AIWorker } from '@/common/AIWorker'
import DynamicGrid from '@/components/basic/DynamicGrid.vue'
import ModelCI from '@/components/cardItems/ModelCI.vue'
import PerfCI from '@/components/cardItems/PerfCI.vue'
import QueueCI from '@/components/cardItems/QueueCI.vue'
import UptimeCI from '@/components/cardItems/UptimeCI.vue'
import ReqCI from '@/components/cardItems/ReqCI.vue'
import InfoCI from '@/components/cardItems/InfoCI.vue'
import TextWorkerCI from '@/components/cardItems/TextWorkerCI.vue'
import type { data } from '@/wailsjs/go/models'
import ImageWorkerCI from '@/components/cardItems/ImageWorkerCI.vue'
import KudosCI from '@/components/cardItems/KudosCI.vue'

const props = defineProps<{
  workerData: AIWorker
}>()

let expanded = false
let cardWidth = ref(300)
let cardHeight = ref(170)

const gridRows = ref(4)
const gridColumns = ref(2)

function updateGrid() {
  if (expanded) {
    gridRows.value -= 1
    gridColumns.value /= 2

    setTimeout(() => {
      // compensate for the margin
      cardWidth.value = (cardWidth.value - 10) / 2
      cardHeight.value /= 2
    }, 20)
  } else {
    gridRows.value += 1
    gridColumns.value *= 2
    setTimeout(() => {
      cardWidth.value = cardWidth.value * 2 + 10
      cardHeight.value *= 2
    }, 20)
  }
}

function toggleExpand() {
  updateGrid()
  expanded = !expanded
}

function getTextWorker(): data.TextWorker {
  return props.workerData as data.TextWorker
}

function getImageWorker(): data.ImageWorker {
  return props.workerData as data.ImageWorker
}
</script>

<style scoped lang="scss">
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

.btn-expand {
  border: 2px;
  border-radius: 50px;
  position: absolute;
  right: 5px;
  bottom: 5px;

  width: 20px;
  aspect-ratio: 1;
  background-color: transparent;
  justify-content: center;
  display: flex;
  transition: background-color 0.2s;

  &:hover {
    background-color: #ffffff11;
  }

  .arrow {
    align-self: center;
  }
}
</style>