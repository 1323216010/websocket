<script setup>
import { ref, reactive } from 'vue'

const input = ref('')
const arr = reactive([])

const ws = new WebSocket('ws://localhost:2023');
console.log(ws)

ws.onopen = (e) => {
}

ws.onmessage = e => {
  arr.push(e.data)
  console.log(e)
}

ws.close = e => {
  console.log('关闭连接');
}

function send() {
  ws.send(input.value)
}
</script>

<template>
  <p style="text-align:center">{{ arr }}</p>
  <p style="text-align:center">
    <el-input v-model="input" placeholder="Please input" clearable style="width: 250px" />
    <el-button type="primary" @click="send">发 送</el-button>
  </p>
</template>

<style scoped>
</style>
