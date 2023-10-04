<script setup>
import AppMainView from '@/layout/AppMainView.vue';
import BreadcrumbView from './BreadcrumbView.vue';
import SliderItem from './SliderItem.vue';
import { useRouter, useRoute } from 'vue-router';
import { usePermissionStore } from '@/stores/permission';
import { computed } from 'vue';
const permission = usePermissionStore();
// console.log(permission.router)
const router = useRouter();
const route = useRoute();
const currentPath = computed(function () {
  console.log(route)
  return route.fullPath
});
const routerPath = permission.router.filter((item) => { return !item.hidden });
function handleSelect(key, keyPath) {
  router.push(keyPath[keyPath.length - 1]);
}
</script>

<template>
  <div style="display: flex;">
    <div style="width: 220px;">
      <el-menu :default-active="route.path" class="el-menu-vertical-demo" @select="handleSelect" active-text-color="#ffd04b" background-color="#545c64"
        text-color="#fff" style="height: 100vh;">
        <div v-for="(item, index) in routerPath" :key="index">
          <SliderItem :item="item"></SliderItem>
        </div>
      </el-menu>
    </div>
    <div style="flex: 1;">
      <BreadcrumbView :path="currentPath"> </BreadcrumbView>
      <AppMainView></AppMainView>
    </div>
  </div>
</template>
