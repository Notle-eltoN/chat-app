<template>
    <div class="message-list">
      <div v-for="message in messages" :key="message.id" class="message">
        <strong>{{ message.username }}</strong>: {{ message.content }}
      </div>
    </div>
  </template>
  
  <script>
  export default {
    props: {
      messages: Array,
    },
  };
  </script>
  