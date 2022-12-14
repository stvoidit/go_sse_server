<template>
  <div>auth: {{inited}}</div>
  <HelloWorld v-if="inited"/>

</template>

<script setup lang="ts">
import HelloWorld from './components/HelloWorld.vue'
import { ref} from "vue"
const data = ref(123)
const inited = ref(false)
fetch("/api/init").then(response => {
  if (response.status === 200) {
    inited.value = true
  } else {
    alert("not auth")
  }
})

</script>


<style scoped>

</style>
