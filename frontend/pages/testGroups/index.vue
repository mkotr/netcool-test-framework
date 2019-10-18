<template>
  <div>
    <div class="md-layout md-alignment-top-center">
      <div class="md-layout-item">
        <md-table
          v-model="searched"
          md-sort="name"
          md-sort-order="asc"
          md-fixed-header
          md-card
        >
          <md-table-toolbar>
            <div class="md-toolbar-section-start">
              <h1 class="md-title">
                Test Groups
              </h1>
            </div>
            <md-button
              class="md-primary md-raised top-button"
              @click="newTestGroup"
            >
              Create New Test Group
            </md-button>
            <md-field md-clearable class="md-toolbar-section-end">
              <md-input
                v-model="search"
                placeholder="Search by name..."
                @input="searchOnTable"
              />
            </md-field>
          </md-table-toolbar>
          <md-table-empty-state
            md-label="No test groups found"
            :md-description="
              `No test group found for this '${search}' query. Try a different search term or create a new test group.`
            "
          >
            <md-button class="md-primary md-raised" @click="newTestGroup">
              Create New Test Group
            </md-button>
          </md-table-empty-state>
          <md-table-row
            slot="md-table-row"
            slot-scope="{ item }"
            @click="onRowClick(item)"
          >
            <md-table-cell md-label="Name" md-sort-by="name">
              {{ item.name }}
            </md-table-cell>
            <md-table-cell md-label="Network Domain" md-sort-by="networkDomain">
              {{ item.networkDomain }}
            </md-table-cell>
            <md-table-cell md-label="Equipment Type" md-sort-by="equipmentType">
              {{ item.equipmentType }}
            </md-table-cell>
            <md-table-cell md-label="Equipment Role" md-sort-by="equipmentRole">
              {{ item.equipmentRole }}
            </md-table-cell>
            <md-table-cell md-label="Probe Type" md-sort-by="probeType">
              {{ item.probeType }}
            </md-table-cell>
          </md-table-row>
        </md-table>
      </div>
    </div>
  </div>
</template>
<script>
import axios from 'axios'

const toLower = (text) => {
  return text.toString().toLowerCase()
}

const searchByName = (items, term) => {
  if (term) {
    return items.filter((item) => toLower(item.name).includes(toLower(term)))
  }

  return items
}

export default {
  data: function() {
    return {
      search: null,
      searched: [],
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
  created: function() {
    this.getTestGroups()
    this.searched = this.testGroups

    console.log(this.$route)
  },
  methods: {
    onRowClick: function(testGroup) {
      console.log(testGroup.id)
      this.$router.push({
        name: `testGroup`,
        params: { testGroup }
      })
    },
    getTestGroups: function() {
      axios
        .get('/api/testGroups')
        .then((res) => {
          //we do this because as a workaround during dev otherwise it returns a html template.
          if (Array.isArray(res.data)) this.testGroups = res.data
          else console.log('there was no response')
        })
        .catch((e) => {
          console.log('an error occured. ', e)
        })
    },
    newTestGroup: function() {
      this.testGroups.push({
        id: 3,
        name: 'Test Group Name3',
        networkDomain: 'networkDomain',
        equipmentRole: 'equipmentRole',
        equipmentType: 'equipmentType',
        probeType: 'SNMP'
      })
    },
    searchOnTable: function() {
      this.searched = searchByName(this.testGroups, this.search)
    }
  }
}
</script>

<style>
button.top-button {
  margin-right: 10px;
}
.md-field {
  max-width: 300px;
}
</style>
