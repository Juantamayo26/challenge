<template>
    <v-container grid-list-xl>
        <v-layout row wrap>
            <v-flex md6>
                <v-card class="mb-3" v-for="(item, index) in buyers" :key="index"> 
                    <v-card-text>
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
                        <p>{{item.age}}</p>
                        <p>{{item.id}}</p>
                        <v-btn @click="consultar(item.id)" class="ml-0" color="info">Consultar</v-btn>
                    </v-card-text>
                </v-card>
            </v-flex>
        </v-layout>
    </v-container>
</template>

<script>
import {queryBuyers} from "../assets/query"

export default {
    data() {
        return {
            info : "",
            buyers: [],
            files: [],
            snackbar: false,
            text: ""
        }
    },
    mounted () {
        fetch('http://localhost:8080/graphql', {
            method: 'POST',
            headers: {"Content-Type":"application/json" },
            body: queryBuyers()
        }).then(res => res.json())
          .then(buyers => this.buyers = buyers.data.queryBuyers)
    },
    methods:{
        consultar(id){
            console.log(id)
        }
    }
}
</script>