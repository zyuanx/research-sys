<script setup>
import { useAuthStore } from '@/stores/auth'
import { computed } from 'vue';
import { useRouter } from 'vue-router'
const props = defineProps(['path'])
const _path = computed(() => props.path.split('/'))

const router = useRouter();
const auth = useAuthStore()
const logout = async () => {
  await auth.userLogout()
  router.replace('/login')
}
</script>


<template>
  <div class="navbar">
    <el-breadcrumb separator="/">
      <el-breadcrumb-item v-for="(item, index) in _path" :key="index">{{ item }}</el-breadcrumb-item>
    </el-breadcrumb>
    <div>
      <el-button link type="primary" icon="el-icon-plus" @click="logout">Logout</el-button>
    </div>
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
  justify-content: space-between;

  padding: 10px;
}
</style>