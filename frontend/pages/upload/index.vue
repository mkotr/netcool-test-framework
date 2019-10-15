<template>
  <div>
    <div class="md-layout md-alignment-top-center">
      <div class="md-layout-item">
        <span class="md-display-1 ">{{ name }}</span>
      </div>
    </div>
    <div class="md-layout">
      <div class="md-layout-item md-size-25">
        <p class="md-body-1">
          Bulk upload your test cases using an .xlsx file
        </p>
        <p class="md-subheading">
          You can download a sample .xlsx file below:
        </p>
        <md-button class="md-raised md-primary">
          Download Sample
        </md-button>
        <p class="md-subheading">
          Upload your file below:
        </p>
        <md-field>
          <label>Click here to upload.</label>
          <md-file v-model="single" />
        </md-field>
      </div>
      <div class="md-layout-item"><p>test</p></div>
    </div>
    <div class="md-layout"><p>Test</p></div>
  </div>
</template>

<script>
export default {
  data: function() {
    return {
      name: 'Bulk Upload Page'
    }
  }
}
</script>

<style scoped>
div.md-layout:not(:first-child) {
  margin-top: 1em;
}
.md-subheading {
  margin-top: 1em;
  margin-bottom: 1em;
}
</style>
