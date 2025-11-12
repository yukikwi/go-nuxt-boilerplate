<template>
  <div>
    OK
  </div>
</template>

<script setup lang="ts">
import { onMounted } from 'vue';
import { homePostSpeakClient } from '@/clients/axios';

// test backend call on mounted
console.log('Calling backend before mount API...');
onBeforeMount(async () => {
  console.log('Calling backend API...');
  const response = await homePostSpeakClient({
      message: 'Hello from Nuxt 3 frontend!',
  });
  console.log('Response from backend:', response);
});
</script>