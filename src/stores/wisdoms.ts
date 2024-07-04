import { ref, computed } from "vue";
import { defineStore } from "pinia";

export const useWisdomStore = defineStore("counter", {
    state: () => ({
        wisdoms: ref<Wisdom[]>(wisdomList),
        activeWisdom: ref(""),
    }),
    actions: {
        randomizeWisdom() {
            const i = Math.floor(Math.random() * this.wisdoms.length);
            this.activeWisdom = this.wisdoms[i].description;
        },
        async fetchWisdoms() {
            const resp = await fetch(`${import.meta.env.VITE_API_URL}/wisdoms`);
            const json: WisdomResponse = await resp.json();

            this.wisdoms = json.wisdoms;
            if (this.activeWisdom === "") {
                this.randomizeWisdom();
            }
        },
    },
});

interface WisdomResponse {
    wisdoms: Wisdom[];
}

interface Wisdom {
    description: string;
    explanation: string;
}

const wisdomList: Wisdom[] = [];
