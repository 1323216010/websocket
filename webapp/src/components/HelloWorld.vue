<script setup>
import { ref } from 'vue'

const str = ref("")

const ws = new WebSocket('ws://localhost:2020');
console.log(ws)

ws.onopen = (e) => {
  console.log('e: ', e);
  console.log('连接成功');
}

ws.onmessage = e => {
  console.log('e: ', e);
  console.log('e.data: ', e.data);
  str.value = e.data
}

ws.close = e => {
  console.log('关闭连接');
}

</script>

<template>
  <h1>{{ str }}</h1>
</template>

<style scoped>
.read-the-docs {
  color: #888;
}
</style>
