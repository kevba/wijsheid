<script setup lang="ts">
import { useWisdomStore } from "./stores/wisdoms";
import Card from "./components/Card.vue";

const wisdomStore = useWisdomStore();
wisdomStore.fetchWisdoms();
</script>

<template>
    <div class="main-body">
        <header @click="() => wisdomStore.randomizeWisdom()">
            <h1 class="title-text">> Wijsheid <</h1>
        </header>

        <div class="app-body">
            <Card></Card>
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
        font-size: 2.5rem;
        font-weight: 800;
        padding: 1rem 0;
        color: yellowgreen;

        @media only screen and (min-width: 768px) {
            font-size: 4rem;
        }
    }
}

.app-body {
    height: 100%;
    display: flex;
    justify-content: center;
}
</style>
