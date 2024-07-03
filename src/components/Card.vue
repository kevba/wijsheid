<script setup lang="ts">
import { useWisdomStore } from "../stores/wisdoms";

const wisdomStore = useWisdomStore();
</script>

<template>
    <div class="card">
        <p class="wisdom-text">{{ wisdomStore.activeWisdom }}</p>
    </div>
</template>

<style scoped>
.card {
    min-height: 30vh;
    height: fit-content;
    max-width: 800px;
    background-color: darkcyan;
    border-radius: 1rem;
    border: solid 0.5rem whitesmoke;
    padding: 1rem;

    @media only screen and (min-width: 768px) {
        border-radius: 4rem;
        border: solid 1.5rem whitesmoke;
        padding: 4rem;
    }
}

.wisdom-text {
    font-size: 1.5rem;
    font-weight: 500;
    color: yellowgreen;

    @media only screen and (min-width: 768px) {
        font-size: 2.5rem;
        font-weight: 500;
    }
}
</style>
