<template>
  <div class="wrapper">
    <h1 class="message">
      Hello, Vue!
    </h1>
    <div class="content">
      <p>
        This is my first vue project
      </p>
    </div>
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
  color: darkgreen;
  margin-top: 0;
  padding-top: 1rem;
  font-weight: bold;
}

.content {
  width: 100vw;
  font-size: larger;
  color: salmon;
  padding-top: 24px;
  font-weight: 700;
}
</style>