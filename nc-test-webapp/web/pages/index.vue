<template>
  <div class="page-container">
    <md-app>
      <md-app-toolbar class="md-primary">
        <span class="md-title">Netcool Testy McTestFace</span>
      </md-app-toolbar>

      <md-app-drawer md-permanent="full">
        <md-toolbar class="md-transparent" md-elevation="0">Navigation</md-toolbar>

        <md-list>
          <md-list-item>
            <md-icon>move_to_inbox</md-icon>
            <nuxt-link to="/">
              <span class="md-list-item-text">Run A Test</span>
            </nuxt-link>
          </md-list-item>

          <md-list-item>
            <md-icon>send</md-icon>
            <nuxt-link to="/upload">
              <span class="md-list-item-text">Bulk Upload</span>
            </nuxt-link>
          </md-list-item>

          <md-list-item>
            <md-icon>delete</md-icon>
            <span class="md-list-item-text">View/Edit Tests</span>
          </md-list-item>

          <md-list-item>
            <md-icon>error</md-icon>
            <span class="md-list-item-text">Probe configuration</span>
          </md-list-item>
        </md-list>
      </md-app-drawer>

      <md-app-content></md-app-content>
    </md-app>
  </div>
</template>

<style lang="scss" scoped>
.md-app {
  //max-height: 400px;
  border: 1px solid rgba(#000, 0.12);
}

// Demo purposes only
.md-drawer {
  width: 230px;
  max-width: calc(100vw - 125px);
}
</style>

<script>
export default {
  name: 'Normal'
}
</script>
