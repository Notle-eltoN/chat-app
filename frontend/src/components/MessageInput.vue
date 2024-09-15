<template>
  <div class="message-input">
    <input 
      type="text" 
      v-model="newMessage" 
      @keyup.enter="submitMessage" 
      placeholder="Type a message..." 
    />
    <button @click="submitMessage">Send</button>
  </div>
</template>

<script>
export default {
  data() {
    return {
      newMessage: ''
    };
  },
  methods: {
    submitMessage() {
      if (this.newMessage.trim() === '') return;  // Prevent sending empty messages
      this.$emit('send', this.newMessage);  // Emit the message to the parent component (ChatRoom.vue)
      this.newMessage = '';  // Clear the input field after sending
    }
  }
};
</script>

<style scoped>
.message-input {
  display: flex;
  padding: 10px;
  background-color: #fff;
  border-top: 1px solid #ccc;
}

.message-input input {
  flex-grow: 1;
  padding: 10px;
  border: 1px solid #ccc;
  border-radius: 4px;
  margin-right: 10px;
}

.message-input button {
  padding: 10px 20px;
  background-color: #1a73e8;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.message-input button:hover {
  background-color: #155abc;
}
</style>
