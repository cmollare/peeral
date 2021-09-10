var subscribed

export function subscribe(obj) {
  subscribe = obj
}

export var JSStreamHandler =  java.lang.Object.extend({
      interfaces: [interfaces.StreamHandler],
      onReceive: function(msg) {
        console.log("received msg : " + msg)
        subscribe.lastMsg = msg
      }
    })


var peer = null

export function start(streamHandler) {
  if (peer) return

  peer = new libp2p.Peer(streamHandler)
  peer.start()
  peer.connectToPeer("")
}

export function send(msg) {
  peer.send(msg)
}
