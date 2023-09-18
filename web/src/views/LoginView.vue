<script setup>
import { ref } from 'vue'
import { login } from "@/api/auth"

import { userAuthStore } from '@/stores/auth'
const formState = ref({
  username: 'admin',
  password: '123456',
});

async function Login() {
  const payload = {
    'username': formState.value.username,
    'password': formState.value.password
  }
  const res = await login(payload)
  console.log(res)
  const auth = userAuthStore()
  if (res.err_code == 0) {
    auth.token = res.data.token
  }
  localStorage.setItem('token', res.data.token)
}
</script>

<template>
  <div style="margin: 20px auto;width: 300px;">
    <a-form :model="formState" name="basic" autocomplete="off" @finish="onFinish" @finishFailed="onFinishFailed">
      <a-form-item :rules="[{ required: true, message: 'Please input your username!' }]">
        <a-input v-model:value="formState.username" />
      </a-form-item>

      <a-form-item :rules="[{ required: true, message: 'Please input your password!' }]">
        <a-input-password v-model:value="formState.password" />
      </a-form-item>
      <a-form-item :wrapper-col="{ offset: 8, span: 16 }">
        <a-button type="primary" html-type="submit" @click="Login">login</a-button>
      </a-form-item>
    </a-form>
  </div>
</template>

<style></style>
