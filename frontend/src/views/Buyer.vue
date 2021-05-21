<template>
    <v-container grid-list-lg>
        <v-layout row>
            <v-flex xs12 class="text-center display-1 font-weight-black my-5">Cliente {{buyer.name}}</v-flex>
        </v-layout>
        <v-flex class="text-center">
            <h3 class="mb-5">Historial de compras</h3>
            <v-card elevation="6" > 
                <h3> {{buyer.transaction[this.page].ip}} </h3>
                <v-list-item v-for="(item, index) in displayedProducts" :key="index">
                    <v-list-item-content>
                        <v-list-item-title>{{item.name}} || {{item.price}}</v-list-item-title>
                    </v-list-item-content>
                </v-list-item>
            </v-card>
        </v-flex>
        <v-layout row justify-center>
            <h3 class="mt-5 mb-5">Usuarios con la misma IP </h3>
        </v-layout>
        <v-layout row wrap justify-center>
            <v-flex xs6 sm4 md4 lg2 class="text-center" v-for="(user, index) in users" :key="index">
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
                            {{user.buyerid.name}}
                        </v-chip>
                        <p>Age: {{user.buyerid.age}}</p>
                    </v-card-text>
                </v-card>
            </v-flex>
        </v-layout>
        <v-layout row justify-center>
            <v-btn @click="page--" v-if="page > 0" class="align-self-start" color="info">Anterior</v-btn>
            <v-btn @click="page++" v-if="page < (this.buyer.transaction.length-1)" class="ml-10 " color="info">Siguiente</v-btn>
        </v-layout>
    </v-container>
</template>

<script>
import {queryOneBuyer, queryProducts, queryIp} from "../assets/query"

export default {
    props: [
        'id'
    ],
    data() {
        return {
            page: 0,
            buyer : [],
            users : [],
            userProducts: [],
            userSameIP: []
        }
    },
    created () {
        this.getBuyer(this.$route.params.id)
    },
    methods:{
        //Obtener el comprador
        async getBuyer(id){
            const res = await fetch('http://localhost:8080/graphql', {
                method: 'POST',
                headers: {"Content-Type":"application/json" },
                body: queryOneBuyer(id)
            })
            const buyer = await res.json()
            this.buyer = buyer.data.queryBuyers[0]
            //console.log(this.buyer.transaction[0].productids)
        },
        //Obtener los usuarios con la misma ip del comprador
        async getUserIp(ip){
            const res = await fetch('http://localhost:8080/graphql', {
                method: 'POST',
                headers: {"Content-Type":"application/json" },
                body: queryIp(ip)
            })
            const users = await res.json()
            this.users = users.data.queryTransactions
        },
        //Obtener todos los productos para hacer la recomendacion
        async getProducts(){
            fetch('http://localhost:8080/graphql', {
                method: 'POST',
                headers: {"Content-Type":"application/json" },
                body: queryProducts()
            }).then(res => res.json())
            .then(products => this.products = products.data.queryProducts)
        }
    },
    computed: {
        displayedProducts: function(){
            this.getUserIp(this.buyer.transaction[this.page].ip)
            return this.buyer.transaction[this.page].productids
        },
        displayedUsers: function(){ 
            return this.users
        }
    }
}
</script>