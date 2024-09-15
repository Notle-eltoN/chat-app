<template>
  <div id="app">
    <!-- Vue Router will inject the matched component here -->
    <router-view></router-view>
  </div>
</template>

<script>
export default {
  name: "App",
};
</script>

<style>
#app {
  text-align: center;
  margin: 0;
  padding: 0;
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  background-color: #f4f4f4;
}

body {
  margin: 0;
  padding: 0;
}
</style>
