<script setup lang="ts">
import H1 from './typography/H1.vue';
import P from './typography/P.vue';
import Code from './typography/Code.vue';
import { ref } from 'vue'

defineProps<{ msg: string }>()

const count = ref(0)
</script>

<template>
  <P>
    Recommended IDE setup:
    <a href="https://code.visualstudio.com/" target="_blank">VSCode</a>
    +
    <a href="https://github.com/johnsoncodehk/volar" target="_blank">Volar</a>
  </P>

  <P>See <Code>README.md</Code> for more information.</P>

  <P>
    <a href="https://vitejs.dev/guide/features.html" target="_blank">
      Vite Docs
    </a>
    |
    <a href="https://v3.vuejs.org/" target="_blank">Vue 3 Docs</a>
  </P>

  <button type="button" @click="count++">count is: {{ count }}</button>
  <P>
    Edit
    <Code>components/HelloWorld.vue</Code> to test hot module replacement.
  </P>
</template>
