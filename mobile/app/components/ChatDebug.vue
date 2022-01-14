<template>
  <FlexboxLayout flexDirection="column" alignContent="flex-end">
    <Label textWrap="true">
      <FormattedString>
        <Span class="dot" text=" "/>
        <Span text=" 0/1000" />
      </FormattedString>
    </Label>
    <TextView editable="false" :text=textReceived flexGrow="80%"/>
    <TextField :text=textToSend />
    <Button className="graybutton" @tap="onButtonTap" wwwstyle="background-color: gray">
      <FormattedString>
        <Label class="fas" :text="'fa-bluetooth'"></Label> 
      </FormattedString>
    </Button>
  </FlexboxLayout>
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

<style scoped>
.dot {
  height: 25px;
  width: 25px;
  background-color: #bbb;
  border-radius: 50%;
  /*display: inline-block;*/
}

.graybutton {
  background-color: gray;
}
</style>