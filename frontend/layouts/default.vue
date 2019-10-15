<template>
  <div class="page-container">
    <md-app>
      <md-app-toolbar class="md-primary" md-elevation="0">
        <md-button
          v-if="!menuVisible"
          class="md-icon-button"
          @click="toggleMenu"
        >
          <md-icon>menu</md-icon>
        </md-button>
        <nuxt-link to="/">
          <span class="md-title">Netcool Testy McTestFace</span>
        </nuxt-link>
      </md-app-toolbar>

      <md-app-drawer :md-active.sync="menuVisible" md-persistent="mini">
        <md-toolbar class="md-transparent" md-elevation="0">
          <span>Navigation</span>

          <div class="md-toolbar-section-end">
            <md-button class="md-icon-button md-dense" @click="toggleMenu">
              <md-icon>keyboard_arrow_left</md-icon>
            </md-button>
          </div>
        </md-toolbar>

        <md-list>
          <md-list-item>
            <md-icon>move_to_inbox</md-icon>
            <span class="md-list-item-text">Run a test</span>
          </md-list-item>

          <md-list-item>
            <md-icon>send</md-icon>
            <span class="md-list-item-text">View/Edit a test</span>
          </md-list-item>
          <nuxt-link to="upload">
            <md-list-item>
              <md-icon>delete</md-icon>
              <span class="md-list-item-text">Bulk Upload</span>
            </md-list-item>
          </nuxt-link>
          <nuxt-link to="probes">
            <md-list-item>
              <md-icon>error</md-icon>
              <span class="md-list-item-text">View/Configure Probes</span>
            </md-list-item>
          </nuxt-link>
        </md-list>
      </md-app-drawer>

      <md-app-content>
        <nuxt />
      </md-app-content>
    </md-app>
  </div>
</template>

<script>
export default {
  name: 'PersistentMini',
  data: () => ({
    menuVisible: false
  }),
  methods: {
    toggleMenu() {
      this.menuVisible = !this.menuVisible
    }
  }
}
</script>
<style>
html {
  font-family: 'Source Sans Pro', -apple-system, BlinkMacSystemFont, 'Segoe UI',
    Roboto, 'Helvetica Neue', Arial, sans-serif;
  font-size: 16px;
  word-spacing: 1px;
  -ms-text-size-adjust: 100%;
  -webkit-text-size-adjust: 100%;
  -moz-osx-font-smoothing: grayscale;
  -webkit-font-smoothing: antialiased;
  box-sizing: border-box;
}

*,
*:before,
*:after {
  box-sizing: border-box;
  margin: 0;
}

.md-app {
  min-height: 350px;
  border: 1px solid rgba(#000, 0.12);
}

.md-drawer {
  width: 230px;
  max-width: calc(100vw - 125px);
}

.md-list-item-text:hover {
  color: aquamarine;
}
.md-list-icon.md-icon-font.md-theme-default:hover {
  color: aquamarine;
}
</style>
