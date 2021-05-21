<template>
    <v-container grid-list-lg>
        <v-layout row>
            <v-flex xs12 class="text-center display-1 font-weight-black my-5">Clientes</v-flex>
        </v-layout>
        <v-layout row wrap >
            <v-flex xs6 sm4 md4 lg2 class="text-center" v-for="(item, index) in buyers" :key="index">
                <v-card elevation="6" > 
                    <v-card-text >
                        <v-chip
                        class="ma-2"
                        color="primary"
                        label
                        >
                            <v-icon left>
                                mdi-account-circle-outline
                            </v-icon>
                            {{item.name}}
                        </v-chip>
                        <p>Age: {{item.age}}</p>
                        <p>ID: {{item.id}}</p>
                        <router-link :to="{name : 'Buyer', params: {id: item.id}}">Consultar</router-link>
                    </v-card-text>
                </v-card>
            </v-flex>
        </v-layout>
        <v-layout row justify-center>
            <v-btn @click="page--" v-if="page!=0" class="align-self-start" color="info">Anterior</v-btn>
            <v-btn @click="page++" v-if="(perPage-1)<buyers.length" class="ml-10 " color="info">Siguiente</v-btn>
        </v-layout>
    </v-container>
</template>

<script>
import {queryBuyers} from "../assets/query"

export default {
    data() {
        return {
            buyers: [],
            page: 0,
            perPage: 12
        }
    },
    created () {
        this.getBuyers(this.page*this.perPage, this.perPage)
    },
    methods:{
        getBuyers(i, e){
            fetch('http://localhost:8080/graphql', {
                method: 'POST',
                headers: {"Content-Type":"application/json" },
                body: queryBuyers(i, e)
            }).then(res => res.json())
              .then(buyers => this.buyers = buyers.data.queryBuyers)
        }
    },
    watch: {
        buyers(){
            this.getBuyers(this.page*this.perPage, this.perPage)
        }
    }
}
</script>
