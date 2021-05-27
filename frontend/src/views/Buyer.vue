<template>
  <v-container grid-list-lg>
    <v-layout row>
      <v-flex xs12 class="text-center display-1 font-weight-black my-5">
          Cliente {{ buyer.name }}
        </v-flex>
    </v-layout>

    <v-flex class="text-center">
      <h3 class="mb-5">Historial de compras</h3>
      <v-card elevation="6">
        <h3>IP: {{ buyer.transaction[page-1].ip }}</h3>
        <v-list>
          <template v-for="(item, index) in displayedProducts">
            <div :key="index">
                <h4> {{ item.name }}</h4>
                <p>${{ item.price }}</p>
            </div>
          </template>
        </v-list>
      </v-card>
    </v-flex>

    <v-layout row justify-center>
      <h3 class="mt-5 mb-5">Usuarios con la misma IP</h3>
    </v-layout>

    <v-layout d-sm-none row wrap justify-center>
      <v-flex
        xs6
        sm4
        md4
        lg2
        class="text-center"
        v-for="(user, index) in users"
        :key="index"
      >
        <v-card elevation="6">
          <v-card-text>
            <v-chip class="ma-2" color="box" label text-color="white">
              <v-icon left> mdi-account-circle-outline </v-icon>
              {{ user.buyerid.name }}
            </v-chip>
            <h4>{{ user.buyerid.age }} años</h4>
          </v-card-text>
        </v-card>
      </v-flex>
    </v-layout>

    <v-sheet d-none class="mx-auto" elevation="8" max-width="1000">
      <v-slide-group class="pa-4 d-none d-sm-flex" show-arrows>
        <v-slide-item v-for="(user, index) in users" :key="index">
          <v-card class="ma-4 text-center" height="100" width="150">
            <v-card-text>
              <v-chip class="ma-2" color="box" text-color="white" label>
                <v-icon left> mdi-account-circle-outline </v-icon>
                {{ user.buyerid.name }}
              </v-chip>
              <h4>{{ user.buyerid.age }} años</h4>
            </v-card-text>
          </v-card>
        </v-slide-item>
      </v-slide-group>
    </v-sheet>

    <v-container class="max-width">
        <v-pagination
        v-model="page"
        class="my-4"
        :length="buyer.transaction.length"
        ></v-pagination>
    </v-container>

    <v-layout row justify-center>
      <h3 class="mt-5 mb-5">Recomendaciones</h3>
    </v-layout>
    

  </v-container>
</template>

<script>
import { queryOneBuyer, queryProducts, queryIp } from "../assets/query";

export default {
  props: ["id"],
  data() {
    return {
      page: 1,
      buyer: [],
      users: [],
    };
  },
  created() {
    this.getBuyer(this.$route.params.id);
  },
  methods: {
    //Obtener el comprador
    async getBuyer(id) {
      const res = await fetch("http://localhost:8080/graphql", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: queryOneBuyer(id),
      });
      const buyer = await res.json();
      this.buyer = buyer.data.queryBuyers[0];
      this.getUserIp(this.buyer.transaction[this.page-1].ip);
    },
    //Obtener los usuarios con la misma ip del comprador
    async getUserIp(ip) {
      const res = await fetch("http://localhost:8080/graphql", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: queryIp(ip),
      });
      const users = await res.json();
      this.users = users.data.queryTransactions;
    },
    //Obtener todos los productos para hacer la recomendacion
    async getProducts() {
      fetch("http://localhost:8080/graphql", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: queryProducts(),
      })
        .then((res) => res.json())
        .then((products) => (this.products = products.data.queryProducts));
    },
  },
  computed: {
    displayedProducts: function () {
      return this.buyer.transaction[this.page-1].productids;
    }
  },
};
</script>