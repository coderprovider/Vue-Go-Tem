<template>
  <div class="wrapper">
    <h1 class="message">
      Hello to Vue
    </h1>
    <HelloButton />
  </div>
</template>

<script setup lang="ts">
import HelloButton from './HelloButton.vue'
</script>

<style scoped>
.wrapper {
  background-color: aliceblue;
  width: 100vw;
  height: 100vh;
}
.message {
  text-align: center;
  color: aquamarine;
  margin-top: 0;
  padding-top: 1rem;
}
</style>