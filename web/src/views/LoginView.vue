<script setup>
import { ref } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { useCounterStore } from '@/stores/counter'
const formState = ref({
  username: 'admin',
  password: '123456',
});

const router = useRouter();
const route = useRoute();
const redirect = route.query.redirect || '/';
console.log(router, route)
async function Login() {
  const payload = {
    'username': formState.value.username,
    'password': formState.value.password
  }
  const auth = useAuthStore()
  await auth.userLogin(payload)
  const counter = useCounterStore()
  counter.increment()
  router.replace(redirect)
}


</script>

<template>
  <div style="margin: 20px auto;width: 300px;">
    <el-form :model="formState" name="basic" autocomplete="off">
      <el-form-item :rules="[{ required: true, message: 'Please input your username!' }]">
        <el-input v-model:value="formState.username" />
      </el-form-item>

      <el-form-item :rules="[{ required: true, message: 'Please input your password!' }]">
        <el-input type="password" v-model:value="formState.password" />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="Login">login</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<style></style>
