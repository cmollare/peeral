<template>
    <Page>
        <ActionBar>
            <Label text="Login"/>
        </ActionBar>

        <Login />
    </Page>
</template>

<script>
  import Login from '@/components/Login.vue'
  //import { Chat } from '@/poc.js'

  export default {
    computed: {
      message() {
        return "Home"
      }
    },
    components: {
        Login
    }
  };
</script>