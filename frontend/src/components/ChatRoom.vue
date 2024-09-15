<template>
    <div class="chat-room">
      <h1>Chat Room</h1>
      
      <!-- Display all chat messages -->
      <MessageList :messages="messages" />
      
      <!-- Input for sending new messages -->
      <MessageInput @send="sendMessage" />
  
      <!-- Leave Room Button -->
      <div class="leave-room">
        <button @click="leaveRoom">Leave Room</button>
      </div>
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
        return this.$store.state.messages;  // Get messages from Vuex store
      },
    },
    created() {
      this.$store.dispatch('listenForMessages');  // Start listening for WebSocket messages
    },
    methods: {
      sendMessage(message) {
        this.$store.dispatch('sendMessage', message);  // Dispatch message sending action to Vuex
      },
      leaveRoom() {
        this.$router.push('/');  // Redirect to another page (e.g., home) when leaving the chat room
        this.$store.commit('clearMessages');  // Clear the messages when leaving
      },
    },
  };
  </script>
  
  <style scoped>
  .chat-room {
    display: flex;
    flex-direction: column;
    height: 100vh;
    justify-content: space-between;
  }
  
  .leave-room {
    text-align: center;
    margin-top: 20px;
  }
  
  button {
    padding: 10px;
    background-color: #ff4b4b;
    color: white;
    border: none;
    border-radius: 5px;
    cursor: pointer;
  }
  
  button:hover {
    background-color: #ff3b3b;
  }
  </style>
  