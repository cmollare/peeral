<template>
  <StackLayout>
    <TextView editable="false" :text=textReceived />
    <TextField :text=textToSend />
    <Button text="Send" @tap="onButtonTap" />
  </StackLayout>
</template>

<script>
import { send, connect, JSStreamHandler, lastMsg, subscribe } from '@/libp2p.js'

export default {
  data() {
    return {
      lastMsg: "",
      textReceived: "",
      textToSend: "Hello I'm nativescript node"
    }
  },
  methods: {
    onButtonTap() {
      console.log(`send ${this.textToSend}`)
      send(this.textToSend)
    }
  },
  created() {
    //var handler = new JSStreamHandler()
    connect(null)
    subscribe(this)
  },
  watch: {
    lastMsg: function(newVal) {
      this.textReceived += newVal + "\n"
    }
  }
}
</script>