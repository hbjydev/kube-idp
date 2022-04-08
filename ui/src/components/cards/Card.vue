<template>
  <div class="bg-white rounded-md shadow-md border border-gray-200" :id="id">
    <div class="p-4 border-b border-gray-200">
      <span class="font-bold">{{ title }}</span>
    </div>
    <div class="p-4">
      <slot></slot>
    </div>
  </div>
</template>

<script setup lang="ts">
import H1 from '../typography/H1.vue';
const { title, id } = defineProps([ 'title', 'id' ])
</script>
