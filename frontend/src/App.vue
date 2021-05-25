<template>
  <v-app>
    <v-app-bar
      color="bar"
      app
      dark
    >
      <v-app-bar-nav-icon @click="drawer = true"></v-app-bar-nav-icon>

      <v-toolbar-title>El Paisa</v-toolbar-title>
      <v-spacer></v-spacer>
        <v-btn icon>
            <v-icon>mdi-magnify</v-icon>
        </v-btn>
    </v-app-bar>

    <v-navigation-drawer
      v-model="drawer"
      absolute
      temporary
    >
      <v-list
        nav
        dense
      >
        <v-list-item-group
          v-model="group"
          active-class="deep-purple--text text--accent-4"
        >
          <v-list-item :to="{name : 'Home'}">
            <v-list-item-icon>
              <v-icon>mdi-home</v-icon>
            </v-list-item-icon>
            <v-list-item-title>Inicio</v-list-item-title>
          </v-list-item>

          <v-list-item :to="{name : 'Upload'}">
            <v-list-item-icon>
              <v-icon>mdi-cloud-upload</v-icon>
            </v-list-item-icon>
            <v-list-item-title>Subir archivos</v-list-item-title>
          </v-list-item>

          <v-list-item :to="{name : 'Buyers'}">
            <v-list-item-icon>
              <v-icon>mdi-checkbox-marked-circle</v-icon>
            </v-list-item-icon>
            <v-list-item-title>Clientes</v-list-item-title>
          </v-list-item>

        </v-list-item-group>
      </v-list>
    </v-navigation-drawer>
    <v-main >
      <router-view/>
    </v-main>
  </v-app>
</template>

<script>
  export default {
    data: () => ({
      drawer: false,
      group: null,
    }),
  }
</script>