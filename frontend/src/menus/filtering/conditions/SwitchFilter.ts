import type Filter from '@/menus/filtering/conditions/Filter'

class SwitchFilter<T> implements Filter<T> {
  switchVal: boolean
  filterOne: Filter<T>
  filterTwo: Filter<T>

  constructor(switchVal: boolean, filterOne: Filter<T>, filterTwo: Filter<T>) {
    this.switchVal = switchVal
    this.filterOne = filterOne
    this.filterTwo = filterTwo
  }

  Pass(toFilter: T): boolean {
    if (this.switchVal) {
      return this.filterOne.Pass(toFilter)
    } else {
      return this.filterTwo.Pass(toFilter)
    }
  }

  Toggle(): boolean {
    this.switchVal = !this.switchVal
    return this.switchVal
  }

  SetSwitch(val: boolean) {
    this.switchVal = val
  }
}
