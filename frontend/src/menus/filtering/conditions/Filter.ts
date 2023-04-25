export default interface Filter<T> {
  Pass(toFilter: T): boolean
}
