<template>
    <div class="chat">
      <h1>Chat Room</h1>
      <MessageList :messages="messages" />
      <MessageInput @send="sendMessage" />
    </div>
  </template>
  
  <script>
  import MessageList from '@/components/MessageList.vue';
  import MessageInput from '@/components/MessageInput.vue';
  
  export default {
    components: {
      MessageList,
      MessageInput,
    },
    computed: {
      messages() {
        return this.$store.state.messages;
      },
    },
    created() {
      this.$store.dispatch('listenForMessages');
    },
    methods: {
      sendMessage(message) {
        this.$store.dispatch('sendMessage', message);
      },
    },
  };
  </script>
  