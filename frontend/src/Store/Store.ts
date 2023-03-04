import { defineStore } from 'pinia'
import { SortingMode } from '@/sorting/SortingMode'
import { ref } from 'vue'

export const useSortingModeStore = defineStore('sortingModes', () => {
  const sortingMode = ref<SortingMode>(SortingMode.Performance)
  const invert = ref<boolean>(false)
  return { sortingMode, invert }
})
