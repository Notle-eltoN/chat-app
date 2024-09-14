<template>
    <div class="message-input">
      <input type="text" v-model="newMessage" @keyup.enter="submitMessage" placeholder="Type a message..." />
      <button @click="submitMessage">Send</button>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        newMessage: '',
      };
    },
    methods: {
      submitMessage() {
        if (this.newMessage.trim() === '') return;
        this.$emit('send', this.newMessage);
        this.newMessage = '';
      },
    },
  };
  </script>
  