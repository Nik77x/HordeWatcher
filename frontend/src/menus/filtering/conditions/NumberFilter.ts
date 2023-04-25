import type Filter from '@/menus/filtering/conditions/Filter'

export class NumberFilter<T> implements Filter<T> {
  op: Comparison
  num: number
  getNumber: (thing: T) => number

  constructor(num: number, op: Comparison, getNumber: (thing: T) => number) {
    this.getNumber = getNumber
    this.num = num
    this.op = op
  }

  Pass(toFilter: T): boolean {
    switch (this.op) {
      case Comparison.LessThan:
        if (this.getNumber(toFilter) < this.num) return true
        else return false
      case Comparison.Equal:
        if (this.getNumber(toFilter) === this.num) return true
        else return false
      case Comparison.GreaterThan:
        if (this.getNumber(toFilter) > this.num) return true
        else return false
    }
  }
}

export enum Comparison {
  LessThan,
  Equal,
  GreaterThan
}
