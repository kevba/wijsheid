<script setup lang="ts">
import { ref } from 'vue';
import { useWisdomStore } from './stores/wisdoms';

const wisdomStore = useWisdomStore()

const randomWisdom = ref(getRandomWisdom())

function randomizeWisdom() {
  randomWisdom.value =getRandomWisdom()
}

function getRandomWisdom() {
  const i = Math.floor(Math.random() * wisdomStore.wisdoms.length)
  return wisdomStore.wisdoms[i]
}

</script>

<template>
  <div class="main-body">
    <header @click="randomizeWisdom">
      <h1 class="title-text">> Wijsheid <</h1>
    </header>

    <div class="app-body">
      <div class="card" >
        <p class="wisdom-text">{{ randomWisdom }}</p>
      </div>
    </div>
  </div>
</template>

<style scoped>
.main-body {
  display: flex;
  flex-direction: column;
  height: 100%;
  width: 100%;
}

header {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: center;
  cursor: pointer;

  .title-text {
    font-size: 4rem;
    font-weight: 800;
    padding: 1rem 0;
    color: yellowgreen;
  }
}

.app-body {
  height: 100%;
  display: flex;
  justify-content: center;
}

.card {
  min-height: 30vh;
  height: fit-content;
  width: 80%;
  max-width: 800px;
  background-color: darkcyan;
  padding: 4rem;
  border-radius: 4rem;
  border: solid 1.5rem whitesmoke;
}

.wisdom-text {
  font-size: 2.5rem;
  font-weight: 500;
  color: yellowgreen;
}

</style>
