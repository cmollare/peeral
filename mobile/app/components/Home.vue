<template>
    <Page>
        <ActionBar>
            <Label text="Home"/>
        </ActionBar>

        <ChatDebug />
    </Page>
</template>

<script>
  import ChatDebug from '@/components/ChatDebug.vue'
  //import { Chat } from '@/poc.js'

  export default {
    computed: {
      message() {
        return "Home"
      }
    },
    components: {
        ChatDebug
    }
  };
</script>

<style scoped lang="scss">
    @import '~@nativescript/theme/scss/variables/blue';

    // Custom styles
    .fas {
        @include colorize($color: accent);
    }

    .info {
        font-size: 20;
        horizontal-align: center;
        vertical-align: center;
    }
</style>
