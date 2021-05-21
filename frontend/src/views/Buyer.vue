<template>
    <v-container grid-list-lg>
        <v-layout row>
            <v-flex xs12 class="text-center display-1 font-weight-black my-5">Cliente {{buyer.name}}</v-flex>
        </v-layout>
        <v-layout row wrap  justify-center>
            <v-flex class="text-center">
                <h3>Historial de compras</h3>
                <v-card elevation="6" > 
                    <p>IP</p>
                    <v-list-item v-for="(item, index) in userProducts" :key="index">
                        <v-list-item-content>
                            <v-list-item-title>{{item.name}} || {{item.price}}</v-list-item-title>
                        </v-list-item-content>
                    </v-list-item>
                </v-card>

                <h3>Usuarios usando tu misma ip</h3>
                <v-list-item v-for="(user, index) in users" :key="index">
                    <v-list-item-content>
                        <v-list-item-title>{{user.buyerid.name}} || {{user.buyerid.age}}</v-list-item-title>
                    </v-list-item-content>
                </v-list-item>
                                
                <v-btn @click="page--"  class="align-self-start" color="info">Anterior</v-btn>
                <v-btn @click="page++"  class="ml-10 " color="info">Siguiente</v-btn>

                <p>Recomendaciones</p>
                <v-card elevation="6" > 
                    <v-card-text >
                        <p>Age: {{buyer.age}}</p>
                        <p>ID: {{id}}</p>
                    </v-card-text>
                </v-card>

            </v-flex>
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
            buyer : [],
            users : [],
            allProducts: [],
            userProducts: [],
            userSameIP: []
        }
    },
    created () {
        this.getBuyer(this.$route.params.id)
    },
    methods:{
        getBuyer(id){
            fetch('http://localhost:8080/graphql', {
                method: 'POST',
                headers: {"Content-Type":"application/json" },
                body: queryOneBuyer(id)
            }).then(res => res.json())
            .then(buyer => this.buyer = buyer.data.queryBuyers[0])
            .then(() => this.userProducts = this.buyer.transaction[0].productids)
            .then(() => this.getUserIp(this.buyer.transaction[0].ip))
        },
        getUserIp(ip){
            fetch('http://localhost:8080/graphql', {
                method: 'POST',
                headers: {"Content-Type":"application/json" },
                body: queryIp(ip)
            }).then(res => res.json())
            .then(users => this.users = users.data.queryTransactions)
        },
        getProducts(){
            fetch('http://localhost:8080/graphql', {
                method: 'POST',
                headers: {"Content-Type":"application/json" },
                body: queryProducts()
            }).then(res => res.json())
            .then(products => this.products = products.data.queryProducts)
        },
        imprimir(){
            console.log(this.users)
        }
    }
}
</script>

