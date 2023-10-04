<script setup>
import { computed } from 'vue';
import {
  Management, Document
} from '@element-plus/icons-vue'
const props = defineProps(['item', 'basePath'])
const _item = computed(() => props.item)
const _basePath = computed(() => props.basePath)

console.log('_item', _item.value)
console.log('_basePath', _basePath.value)
function resolvePath(path) {
  if (_basePath.value === undefined || _basePath.value.length === 0) return path
  return `${_basePath.value}/${path}`
}
</script>

<template>
  <div>
    <div v-if="_item.redirect">
      <SliderItem v-for="(item, index) in _item.children" :key="index" :item="item"></SliderItem>
    </div>
    <el-menu-item v-else-if="!_item.children" :index="resolvePath(_item.path)">
      <el-icon>
        <Management />
      </el-icon>
      <span>{{ _item.name }}</span>
    </el-menu-item>
    <el-sub-menu v-else :index="resolvePath(_item.path)">
      <template #title>
        <el-icon>
          <Document />
        </el-icon>
        <span>{{ _item.name }}</span>
      </template>
      <SliderItem v-for="(item, index) in _item.children" :key="index" :item="item" :base-path="resolvePath(_item.path)"></SliderItem>
    </el-sub-menu>
  </div>
</template>