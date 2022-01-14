<template>
  <StackLayout>
    <Label :text="'login.loginLabel'|L" />
    <TextField :text="login" hint="Login" />
    <Label :text="'login.pwdLabel'|L" />
    <TextField :text="pwd" secure />
    <Button text="Login" @tap="onButtonTap" style="background-color: gray" />
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
