<template>
  <div class="button-wrapper">
    <button @click="hello" class="hello">
      Say Hello
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
.button-wrapper {
  display: flex;
  justify-content: center;
}

.hello {
  background-color: aqua;
  border-radius: 20px;
  border: 1px solid gray;
  padding: 0.8rem 1rem;
  transition: .3s;
}

.hello:hover {
  cursor: pointer;
  background-color: aquamarine;
}
</style>