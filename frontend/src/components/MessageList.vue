<template>
  <div class="message-list">
    <div v-for="(message, index) in messages" :key="index" class="message-item">
      <strong>{{ message.username }}</strong>: {{ message.content }}
    </div>
  </div>
</template>

<script>
export default {
  props: {
    messages: {
      type: Array,
      required: true
    }
  }
};
</script>

<style scoped>
.message-list {
  flex-grow: 1;
  padding: 10px;
  overflow-y: auto;
  background-color: #f9f9f9;
  border: 1px solid #ccc;
}

.message-item {
  margin-bottom: 10px;
}

.message-item strong {
  color: #1a73e8;
}
</style>
