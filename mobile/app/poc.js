export function Test() {
  var counter = new myexportedpackage.Counter()
  counter.inc()
  return counter.getValue()
}

export function Chat() {
  var peer = new libp2p.Peer()
  peer.start()
  peer.connectToPeer("")
  return ""
}