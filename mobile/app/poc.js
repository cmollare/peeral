export function Test() {
  var counter = new myexportedpackage.Counter()
  counter.inc()
  return counter.getValue()
}