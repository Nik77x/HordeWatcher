import type Filter from '@/menus/filtering/conditions/Filter'

export class TextFilter<T> implements Filter<T> {
  filterText: string
  getText: (thing: T) => string

  constructor(text: string, getText: (thing: T) => string) {
    this.getText = getText
    this.filterText = text
  }

  Pass(toFilter: T): boolean {
    if (toFilter === this.getText(toFilter)) {
      return false
    }
    return true
  }
}
