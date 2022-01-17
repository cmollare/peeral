var subscribed

export function subscribe(obj) {
  subscribe = obj
}

export var EventCallback = java.lang.Object.extend({
  interfaces: [api.EventCallback],
  Event: function(msg) {
    console.log("received event "+msg)
    subscribe.lastMsg = msg
  }
})

var peer = null

export function connect(streamHandler) {
  if (peer) return

  peer = new libp2p.Peer(streamHandler)
  peer.connect()
}

export function start(streamHandler) {
  if (peer) return

  peer = new libp2p.Peer(streamHandler)
  peer.start()
  peer.connectToPeer("")
}

export function send(msg) {
  peer.send(msg)
}
