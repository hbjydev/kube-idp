<template>
  <Card :title="title">
    <div
      :class="{
        'text-green-500': success,
        'text-red-500': danger,
        'text-orange-500': meh,
      }"
      class="flex items-center gap-4">
      <div class="h-8 w-8">
        <slot name="icon"></slot>
      </div>
      <span class="text-4xl font-bold">
        {{ count }}
      </span>
    </div>
  </Card>
</template>

<script setup lang="ts">
import Card from './Card.vue';

interface Props {
  title: string;
  count: string;
  success?: boolean;
  danger?: boolean;
  meh?: boolean;
}

const { title, count, success, danger, meh } = defineProps<Props>();
</script>
