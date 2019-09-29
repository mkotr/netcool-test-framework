<template>
  <div>
    <h1>Hello! {{ msg }}</h1>
  </div>
</template>

<script>
export default {
  name: 'Test',
  props: {
    msg: String
  }
}
</script>

<style></style>
