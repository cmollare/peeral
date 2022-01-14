<template>
  <StackLayout>
    <Label :text="'login.loginLabel'|L" />
    <TextField :text="login" hint="Login" />
    <Label :text="'login.pwdLabel'|L" />
    <TextField :text="pwd" secure />
    <Button text="Login" @tap="onButtonTap" />
    <Label class="fas" text="fa-chess-knight | fonticon"></Label>
  </StackLayout>
</template>

<script>

import Home from '@/pages/Home'

export default {
  data() {
    return {
      login: "",
      pwd: ""
    }
  },
  methods: {
    onButtonTap() {
      console.log("login")
      this.$navigateTo(Home)
    }
  }
}
</script>

<style scoped>
button {
    background-color: gray;
    border-radius: 25%;
}
</style>
