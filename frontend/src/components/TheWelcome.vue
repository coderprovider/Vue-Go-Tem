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
  axios.post("/api/hello", {
    message: 'Hello from Vue'
  })
  .then(function (response) {
    console.log(response);
  })
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