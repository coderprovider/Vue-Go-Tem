<template>
  <div class="wrapper">
    <button @click="hello" class="hello">
      Say hello
    </button>
  </div>
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
.wrapper {
  display: flex;
  width: 100vw;
  justify-content: center;
}
.hello {
  background-color: aqua;
  border-radius: 20px;
  border: 0px;
  padding: 10px;
  height: 3rem;
}
</style>