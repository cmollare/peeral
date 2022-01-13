<template>
  <StackLayout>
    <Label :text="'login.loginLabel'|L" />
    <TextField :text="login" hint="Login" />
    <Label :text="'login.pwdLabel'|L" />
    <TextField :text="pwd" secure />
  </StackLayout>
</template>

<script>

export default {
  data() {
    return {
      login: "",
      pwd: ""
    }
  }
}
</script>