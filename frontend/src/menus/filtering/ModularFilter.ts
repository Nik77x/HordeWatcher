import type Filter from '@/menus/filtering/conditions/Filter'

export default class ModularFilter<T> {
  conditions: Filter<T>[] = [] as Filter<T>[]

  filter(toFilter: T[]): T[] {
    return toFilter.filter((f) => {
      for (const item of this.conditions) {
        // if item doesn't satisfy one filter remove it
        if (!item.Pass(f)) return false
      }
    })
  }

  addCondition(cond: Filter<T>): ModularFilter<T> {
    this.conditions.push(cond)
    return this
  }

  AddCondition(cond: Filter<T>) {
    this.conditions.push(cond)
  }

  RemoveCondition(condToRemove: Filter<T>) {
    for (let i = 0; i < this.conditions.length; i++) {
      if (this.conditions[i] === condToRemove) {
        this.conditions.splice(i, 1)
      }
    }
  }
}
