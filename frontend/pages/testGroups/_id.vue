<template>
  <div>
    <p>{{ $route.params.id }}</p>
    <p>hello</p>
    <p>{{ $route.query.testGroup.name }}</p>
  </div>
</template>

<script>
export default {
  created: function() {
    console.log('we have routed here.')
    console.log(this)
    console.log(this.$route)
  }
}
</script>

<style></style>
