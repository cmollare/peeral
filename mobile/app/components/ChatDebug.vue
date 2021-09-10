<template>
  <StackLayout>
    <TextView editable="false" :text=textReceived />
    <TextField :text=textToSend />
    <Button text="Send" @tap="onButtonTap" />
  </StackLayout>
</template>

<script>
export default {
  data() {
    return {
      textReceived: "",
      textToSend: "Hello I'm nativescript node"
    }
  },
  methods: {
    onButtonTap() {
      console.log(`send ${this.textToSend}`)
    }
  }
}
</script>