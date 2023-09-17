<script setup>
import { ref } from 'vue'
import { login } from "@/api/auth.js"
const username = ref('admin')
const password = ref('123456')


async function Login() {
  const payload = {
    'username': username.value,
    'password': password.value
  }
  const res = await login(payload)
  console.log(res)
}
</script>

<template>
  <div style="display: flex;justify-items: center;flex-direction: column;width: 100vw;">
    <div style="width: 200px;">
      <a-input v-model:value="username" placeholder="username" />
      <a-input-password v-model:value="password" placeholder="input password" />
      <a-button type="primary" @click="Login">login</a-button>
    </div>
  </div>
</template>

<style></style>
