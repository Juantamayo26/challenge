<template>
  <v-container grid-list-lg>
    <v-layout row>
      <v-flex xs12 class="text-center display-1 font-weight-black my-5"
        >Clientes</v-flex
      >
    </v-layout>
    <v-layout row wrap>
      <v-flex
        xs6
        sm4
        md4
        lg2
        class="text-center"
        v-for="(item, index) in buyers"
        :key="index"
      >
        <v-card elevation="6">
          <v-card-text>
            <v-chip class="ma-2" color="box" label text-color="white">
              <v-icon left> mdi-account-circle-outline </v-icon>
              {{ item.name }}
            </v-chip>
            <h4>{{ item.age }} años</h4>
            <h4>ID: {{ item.id }}</h4>
            <router-link :to="{ name: 'Buyer', params: { id: item.id } }"
              >Consultar</router-link
            >
          </v-card-text>
        </v-card>
      </v-flex>
    </v-layout>
    <v-layout class="mt-5" row justify-center>
      <v-row align="center" justify="space-around">
        <v-btn @click="page--" v-if="page != 0" color="info">Anterior</v-btn>
        <v-btn @click="page++" v-if="perPage - 1 < buyers.length" color="info"
          >Siguiente</v-btn
        >
      </v-row>
    </v-layout>
  </v-container>
</template>

<script>
import { queryBuyers } from "../assets/query";

export default {
  data() {
    return {
      buyers: [],
      page: 0,
      perPage: 12,
    };
  },
  created() {
    this.getBuyers(this.page * this.perPage, this.perPage);
  },
  methods: {
    getBuyers(i, e) {
      fetch("http://localhost:8080/graphql", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: queryBuyers(i, e),
      })
        .then((res) => res.json())
        .then((buyers) => (this.buyers = buyers.data.queryBuyers));
    },
  },
  watch: {
    buyers() {
      this.getBuyers(this.page * this.perPage, this.perPage);
    },
  },
};
</script>