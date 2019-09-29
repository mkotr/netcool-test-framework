<template>
  <div class="about">
    <h1>This is an about page</h1>
    <Test msg="Morokot" />
  </div>
</template>

<script>
import Test from '@/components/Test.vue'
export default {
  name: 'about',
  components: {
    Test
  }
}
</script>
