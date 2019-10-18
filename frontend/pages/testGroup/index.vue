<template>
  <div>
    <p>xD</p>
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

    this.$router.replace({
      path: 'testGroup'
    })
  }
}
</script>

<style></style>
