<template>
  <div class="md-layout">
    <div class="md-layout-item">
      <p>Probe Config Page</p>
      <md-table class="probe-table">
        <md-table-row>
          <md-table-head>Name</md-table-head>
          <md-table-head>Desc</md-table-head>
          <md-table-head>Hostname</md-table-head>
          <md-table-head>Port</md-table-head>
        </md-table-row>
        <md-table-row v-for="probe in probes" :key="probe.name">
          <md-table-cell>{{ probe.name }}</md-table-cell>
          <md-table-cell>{{ probe.desc }}</md-table-cell>
          <md-table-cell>{{ probe.hostname }}</md-table-cell>
          <md-table-cell>{{ probe.port }}</md-table-cell>
          <mt-table-cell>
            <md-button class="md-icon-button md-raised md-accent">
              <md-icon>edit</md-icon>
              <p>Edit</p>
            </md-button>
            <md-button class="md-icon-button md-raised md-accent">
              <md-icon>call</md-icon>
              <p>Test</p>
            </md-button>
          </mt-table-cell>
        </md-table-row>
      </md-table>
    </div>
  </div>
</template>

<script>
import axios from 'axios'
export default {
  data: function() {
    return {
      probes: [
        {
          name: 'OMI',
          desc: 'omi probe',
          hostname: 'hello',
          port: '4000'
        },
        {
          name: 'MessageBus',
          desc: 'messagebus probe',
          hostname: 'http://lxapp6661.dc.corp.telstra.com',
          port: '4000'
        }
      ]
    }
  },
  methods: {
    getProbes: function() {
      axios.get('/api/probes/')
    .then((res) => {
      return { probes: res.probes}
    })
    .catch((e) => { 
      console.log("an error occured. ",e)
    })
    }
  },
  created: function() {
    this.getProbes()
  }
}
</script>

<style>
.probe-table {
  padding-left: 200px;
  padding-right: 200px;
}
</style>
