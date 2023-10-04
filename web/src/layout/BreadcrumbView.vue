<script setup>
import { computed } from 'vue';
const props = defineProps(['path'])
const _path = computed(() => props.path.split('/'))
</script>
<template>
  <div class="navbar">
    <el-breadcrumb separator="/">
      <el-breadcrumb-item v-for="(item, index) in _path" :key="index">{{ item }}</el-breadcrumb-item>
    </el-breadcrumb>
  </div>
</template>

<style lang="scss" scoped>
.navbar {
  height: 50px;
  overflow: hidden;
  position: relative;
  background: #fff;
  box-shadow: 0 1px 4px rgba(0, 21, 41, 0.08);

  display: flex;
  align-items: center;

  padding: 10px;
}
</style>