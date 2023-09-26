<script setup>
import { computed } from 'vue';

computed(() => {
  return {
    key: '1',
  };
});
</script>

<template>
  <transition name="fade-transform" mode="out-in">
    <router-view :key="key" />
  </transition>
</template>