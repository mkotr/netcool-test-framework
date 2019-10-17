<template>
  <div>
    <div class="md-layout md-alignment-top-center">
      <div class="md-layout-item">
        <span class="md-display-1 ">{{ name }}</span>
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
        <div class="md-layout-item md-layout md-gutter">
          <div class="md-layout-item">
            <md-field>
              <label>Click here to upload.</label>
              <md-file
                v-model="singleFile"
                @md-change="handleFileInputChange()"
              />
            </md-field>
          </div>
          <div class="md-layout-item parser-button">
            <md-button
              class="md-raised md-primary"
              :disabled="!isParserButtonDisabled"
              @click="onParserButtonClick"
            >
              Parse File
            </md-button>
          </div>
        </div>
      </div>
      <div class="md-layout-item">
        <p>this is where the parsed stuff will be</p>
        <md-table>
          <md-table-row>
            <md-table-head>Test Name</md-table-head>
            <md-table-head>Sleep Time</md-table-head>
            <md-table-head>Network Domain</md-table-head>
            <md-table-head>Test Name</md-table-head>
            <md-table-head>Test Name</md-table-head>
            <md-table-head>Test Name</md-table-head>
            <md-table-head>Test Name</md-table-head>
          </md-table-row>
          <md-table-row>
            <md-table-cell>Test</md-table-cell>
          </md-table-row>
        </md-table>
      </div>
      <div class="md-layout-item" />
    </div>

    <div class="md-layout">
      <p>Hello</p>
    </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  data: function() {
    return {
      name: 'Bulk Upload Page',
      singleFile: '',
      file: null,
      isParserButtonDisabled: false,
      tests: null,
      traps: null,
      varbinds: null,
      expectedResults: null
    }
  },
  methods: {
    handleFileInputChange: function() {
      console.log(this.singleFile)
      if (this.singleFile) {
        this.isParserButtonDisabled = true
        this.file = this.singleFile.target.files[0]
      }
    },
    onParserButtonClick: function() {
      var formData = new FormData()
      console.log('the file used is ', this.file.name)
      formData.append('file', this.file)
      axios
        .post('/api/fileParser', formData, {
          header: {
            'Content-Type': 'multipart/form-data'
          }
        })
        .then((res) => {
          if (res.data.status === 0) {
            this.tests = res.data.data.tests
            this.varbinds = res.data.data.varbinds
            this.expectedResults = res.data.data.expectedResults
            this.traps = res.data.data.traps

            console.log('checking the state')
          }
          //TODO: set the data. from the response.
          else {
            alert('An error occurred')
          }
        })
        .catch((e) => {
          console.log('error?')
          console.log(e)
          alert('An error occurred: ', e)
        })
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
.md-layout-item.parser-button {
  padding-top: 10px;
}
</style>
