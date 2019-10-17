<template>
  <div>
    <div class="md-layout md-alignment-top-center">
      <p>Test Groups</p>
      <md-table>
        <md-table-row>
          <md-table-head>Name</md-table-head>
          <md-table-head>Network Domain</md-table-head>
          <md-table-head>Equipment Role</md-table-head>
          <md-table-head>Equipment Type</md-table-head>
          <md-table-head>Probe Type</md-table-head>
        </md-table-row>
        <md-table-row
          v-for="item in testGroups"
          :key="item.id"
          @click="alert(item.name)"
        >
          <md-table-cell v>
            {{ item.name }}
          </md-table-cell>
          <md-table-cell v>
            {{ item.networkDomain }}
          </md-table-cell>
          <md-table-cell v>
            {{ item.equipmentRole }}
          </md-table-cell>
          <md-table-cell v>
            {{ item.equipmentType }}
          </md-table-cell>
          <md-table-cell v>
            {{ item.probeType }}
          </md-table-cell>
        </md-table-row>
      </md-table>
    </div>
  </div>
</template>
<script>
export default {
  data: function() {
    return {
      testGroups: [
        {
          id: 0,
          name: 'Test Group Name',
          networkDomain: 'IP',
          equipmentRole: 'equipmentRole',
          equipmentType: 'equipmentType',
          probeType: 'OMi'
        },
        {
          id: 1,
          name: 'Test Group Name2',
          networkDomain: 'Mobiles',
          equipmentRole: 'equipmentRole',
          equipmentType: 'equipmentType',
          probeType: 'MessageBus'
        }
      ]
    }
  },
  methods: {
    alert: function(name) {
      alert(name)
    }
  }
}
</script>

<style></style>
