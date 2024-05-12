<template>
  <div class="welcome">
    This is a Vue frontend.
  </div>
  <button @click="hello">
    Say hello
  </button>
</template>

<script setup lang="ts">
import axios from 'axios';

function hello() {
  axios.post("http://localhost:3000/api/hello", "Hello from Vue")
  .catch((error) => {
    window.alert(`The API returned an error: ${error}`);
  })
}
</script>

<style scoped>
.welcome {
  font-size: 36;
  color: aquamarine;
}
</style>