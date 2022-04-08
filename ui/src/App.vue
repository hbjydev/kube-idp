<script setup lang="ts">
// This starter template is using Vue 3 <script setup> SFCs
// Check out https://v3.vuejs.org/api/sfc-script-setup.html#sfc-script-setup
import HelloWorld from './components/HelloWorld.vue'
import Navbar from './components/layout/Navbar.vue';
import Card from './components/cards/Card.vue';
import Counter from './components/cards/Counter.vue';
import { CodeIcon, ShareIcon } from '@heroicons/vue/outline';
</script>

<template>
  <div class="bg-gray-100 flex-grow">
    <Navbar />
    <main class="p-4 flex flex-col gap-4 max-w-screen-lg mx-auto">
      <div class="flex gap-4">
        <Counter title="Deployments (Active)" class="w-full" count="2" success>
          <template v-slot:icon>
            <CodeIcon class="w-full h-full" />
          </template>
        </Counter>
        <Counter title="Deployments (Failing)" class="w-full" count="4" danger>
          <template v-slot:icon>
            <CodeIcon class="w-full h-full" />
          </template>
        </Counter>
        <Counter title="Services (Up)" class="w-full" count="5" success>
          <template v-slot:icon>
            <ShareIcon class="w-full h-full" />
          </template>
        </Counter>
        <Counter title="Services (Down)" class="w-full" count="2" danger>
          <template v-slot:icon>
            <ShareIcon class="w-full h-full" />
          </template>
        </Counter>
      </div>
      <Card title="Welcome!">
        <HelloWorld msg="Hello Vue 3 + TypeScript + Vite" />
      </Card>
    </main>
  </div>
</template>

<style>
@import './assets/style.css';
</style>
