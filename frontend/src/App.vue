<script lang="ts" setup>

import {computed, Ref, ref} from "vue";
import {GetLetters} from "../wailsjs/go/main/App";
import {main} from "../wailsjs/go/models";
import Word = main.Word;
import Frequency = main.Frequency;
import Response = main.Response;

const words = ref('');
const wordToAnalise = ref('');
const response: Ref<Response> = ref(new Response())
const keys: Ref<string[]> = ref([]);
const letters: Ref<{ [key: string]: Word; }> = ref({});
const frequencies: Ref<{ [key: string]: Frequency; }> = ref({});

const calculateTotal = computed(()=>{
  let total = 0;
  for (let key in frequencies.value) {
    total += frequencies.value[key].value * frequencies.value[key].count;
  }
  return total;
})

const countWords = (): void => {
  wordToAnalise.value = words.value.replaceAll(' ', '');
  GetLetters(words.value).then((result: Response) => {
    response.value = result;
    keys.value = Object.keys(result.words);
    letters.value = result.words;
    frequencies.value = result.frequencies;
    console.log(frequencies.value);
  });
  words.value = '';
}

const getRedundancy = (frequency: number): number => {
  let frequenciesToAnalise = Object.entries(frequencies.value);
  let index = frequenciesToAnalise.findIndex((item: any) => item[1].value === frequency);
  return frequenciesToAnalise[index][1].count;
}

const getBgStripped = (index: number): string => {
  return index % 2 === 0 ? 'bg-indigo-100' : 'bg-indigo-200';
}

</script>

<template>
  <main class="w-full h-full bg-slate-200 ">
    <div class="bg-neutral-800 p-2">
      <h1 class="text-2xl font-bold text-center uppercase text-indigo-500">IDEAL LIMIT CALCULATOR</h1>
    </div>

    <div class="pt-2 px-4 w-full flex gap-2">
      <div class="bg-slate-50 p-2 rounded-lg shadow-lg w-5/6">
        <h3 class="font-bold">Word:</h3>
        <p class="overflow-x-scroll w-full">{{ wordToAnalise }}</p>
      </div>
      <div class="bg-slate-50 p-2 rounded-lg shadow-lg w-1/6">
        <h3 class="font-bold">Result:</h3>
        <p>{{ calculateTotal }}</p>
      </div>
    </div>
    <div class="grid grid-cols-2 gap-2 p-4 ">
      <section class="col-span-1 bg-slate-50 shadow-lg border border-slate-200 p-2 rounded-lg flex flex-col">
        <h3 class="font-bold text-lg text-center uppercase text-blue-500">Data Entry</h3>
        <div class="flex flex-col gap-2 h-full">
          <label for="words">Enter the word to calculate:</label>
          <textarea class="rounded outline-0 p-2 bg-neutral-200" id="words" v-model="words" rows="2"></textarea>
          <button
              class="bg-emerald-400 text-white uppercase text-lg font-bold p-1 rounded hover:bg-emerald-500 transition-colors mt-auto"
              v-on:click="countWords"
          >Calculate
          </button>
        </div>
      </section>
      <section class="col-span-1 bg-slate-50 shadow-lg border border-slate-200 p-2 rounded-lg">
        <h3 class="font-bold text-lg text-center uppercase text-red-500">Frequency Redundancy</h3>
        <div class="flex flex-col gap-2">
          <p><strong>Word size: </strong>{{ response.size }}</p>
        </div>
        <section class="grid grid-rows-[repeat(6,1fr)] grid-flow-col gap-x-1">
          <div v-for="(frequency,,index) in frequencies" class="flex gap-2 text-xs justify-start w-full p-1"
               :class="getBgStripped(Number(index))">
            <span class="w-full whitespace-nowrap"><strong>F:</strong>{{ frequency.value.toFixed(7) }} </span>
            <span class="w-full whitespace-nowrap"><strong>R:</strong>{{ frequency.count }}</span>
            <span class="w-full whitespace-nowrap"><strong>T:</strong>{{ frequency.total.toFixed(7) }}</span>
          </div>
        </section>
      </section>
      <section class="col-span-2 bg-slate-50 shadow-lg border border-slate-200 p-2 rounded-lg">
        <h3 class="font-bold text-lg text-center uppercase text-green-500 mb-1">Development table</h3>
        <div class="grid grid-rows-[repeat(12,1fr)] grid-flow-col gap-x-1">
          <div v-for="(key,index) in keys" class="flex gap-2 text-xs justify-start w-full p-1"
               :class="getBgStripped(index)">
            <span class="w-full whitespace-nowrap"><strong>L:</strong> {{ key }} </span>
            <span class="w-full whitespace-nowrap"><strong>Q:</strong>{{ letters[key].count }}</span>
            <span class="w-full whitespace-nowrap"><strong>F:</strong>{{ letters[key].frequency.toFixed(7) }}</span>
            <span class="w-full whitespace-nowrap"><strong>R:</strong>{{ getRedundancy(letters[key].frequency) }}</span>
          </div>
        </div>
      </section>
    </div>

  </main>

</template>

