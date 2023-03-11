import { SortingMode } from '@/sorting/SortingMode'
import type { AIWorker } from '@/common/AIWorker'

export class Sorter {
  sort(wis: AIWorker[], mode: SortingMode, invert: boolean): AIWorker[] {
    return this.NumberSort(wis, (wi) => this.getSortValue(wi, mode), invert)
  }

  NumberSort(arr: AIWorker[], getSortValue: (wi: AIWorker) => number, invert: boolean): AIWorker[] {
    return arr.sort((a, b) => {
      let val = getSortValue(b) - getSortValue(a)
      if (invert) val *= -1
      return val
    })
  }

  getSortValue(wi: AIWorker, mode: SortingMode): number {
    switch (mode) {
      case SortingMode.Uptime:
        return wi.uptime
      case SortingMode.Performance: {
        // remove useless text from the json then parse
        const searchPattern = RegExp('[0-9.]')
        const val = searchPattern.exec(wi.performance)
        if (val) return parseFloat(val[0])
        else return -1
      }

      case SortingMode.Kudos:
        return wi.kudos_rewards
      default:
        return -1
    }
  }
}

export default Sorter
