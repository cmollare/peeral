<template>
    <Page>
        <ActionBar>
            <Label text="Login"/>
        </ActionBar>
        <FlexboxLayout flexDirection="column" justifyContent="space-around">
          <Login alignSelf="center" width="150" height="200" />
        </FlexboxLayout>
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