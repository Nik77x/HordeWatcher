import { SortingMode } from '@/sorting/SortingMode'
import type { Data } from '@/wailsjs/go/models'

export class Sorter {
  sort(wis: Data.WorkerInfo[], mode: SortingMode, invert: boolean): Data.WorkerInfo[] {
    return this.NumberSort(wis, (wi) => this.getSortValue(wi, mode), invert)
  }

  NumberSort(
    arr: Data.WorkerInfo[],
    getSortValue: (wi: Data.WorkerInfo) => number,
    invert: boolean
  ): Data.WorkerInfo[] {
    return arr.sort((a, b) => {
      let val = getSortValue(b) - getSortValue(a)
      if (invert) val *= -1
      return val
    })
  }

  getSortValue(wi: Data.WorkerInfo, mode: SortingMode): number {
    switch (mode) {
      case SortingMode.Uptime:
        return wi.uptime
      case SortingMode.Performance:
        // remove useless text from the json then parse
        return parseFloat(wi.performance.replace(' tokens per second', ''))
      case SortingMode.Jobs:
        return wi.uncompleted_jobs
      case SortingMode.Kudos:
        return wi.kudos_details.generated
      case SortingMode.GenLength:
        return wi.max_length
      case SortingMode.ContextSize:
        return wi.max_context_length
    }
  }
}

export default Sorter
