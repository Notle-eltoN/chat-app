<template>
    <div class="signup">
      <h1>Sign Up</h1>
      <form @submit.prevent="signup">
        <input type="text" v-model="username" placeholder="Username" required />
        <input type="password" v-model="password" placeholder="Password" required />
        <button type="submit">Sign Up</button>
      </form>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        username: '',
        password: '',
      };
    },
    methods: {
      signup() {
        this.$store.dispatch('signup', {
          username: this.username,
          password: this.password,
        })
        .then(() => {
          this.$router.push('/');
        })
        .catch(() => {
          alert('Error signing up');
        });
      },
    },
  };
  </script>
  