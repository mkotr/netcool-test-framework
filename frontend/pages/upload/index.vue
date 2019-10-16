<template>
  <div>
    <div class="md-layout md-alignment-top-center">
      <div class="md-layout-item">
        <span class="md-display-1 ">{{ name }}</span>
      </div>
    </div>
    <div class="md-layout">
      <div class="md-layout-item md-size-50">
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
      isParserButtonDisabled: false
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
          if (res.data.status !== 1) alert('An error occurred')
          //TODO: set the data. from the response.
          console.log(res)
        })
        .catch((e) => {
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
