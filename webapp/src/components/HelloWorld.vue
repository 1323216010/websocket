<script setup>
import { reactive } from 'vue'

const arr = reactive([])

const ws = new WebSocket('ws://localhost:2020');
console.log(ws)

ws.onopen = (e) => {
  ws.send("233")
}

ws.onmessage = e => {
  arr.push(e.data)
  console.log(arr)
}

ws.close = e => {
  console.log('关闭连接');
}

</script>

<template>
  <p style="text-align:center">{{ arr }}</p>
</template>

<style scoped>
.read-the-docs {
  color: #888;
}
</style>
