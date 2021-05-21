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
                            {{item.Name}}
                        </v-chip>
                        <p>{{item.Age}}</p>
                        <p>{{item.ID}}</p>
                        <v-btn @click="consultar(item.id)" class="ml-0" color="info">Consultar</v-btn>
                    </v-card-text>
                </v-card>
            </v-flex>

            <v-flex md6>
                <v-card class="mb-3">
                    <!--v-form method="post" action="http://localhost:8002/buyers"> -->
                    <v-form @submit.prevent="addFiles"> 
                        <v-file-input
                            v-model="files"
                            color="deep-purple accent-4"
                            counter
                            label="File input"
                            multiple
                            placeholder="Select your files"
                            prepend-icon="mdi-paperclip"
                            outlined
                            :show-size="1000"
                        >
                            <template v-slot:selection="{ index, text }">
                            <v-chip
                                v-if="index < 2"
                                color="deep-purple accent-4"
                                dark
                                label
                                small
                            >
                                {{ text }}
                            </v-chip>

                            <span
                                v-else-if="index === 2"
                                class="overline grey--text text--darken-3 mx-2"
                            >
                                +{{ files.length - 2 }} File(s)
                            </span>
                            </template>
                        </v-file-input>
                    <v-btn type="submit" block color="success">Agregar archivos</v-btn>
                    </v-form>

                    <v-snackbar
                    v-model="snackbar"
                    :timeout="2000"
                    >
                    {{ text }}
                        <v-btn
                        color="pink"
                        text
                        @click="snackbar = false"
                        >
                        Cerrar
                        </v-btn>
                    </v-snackbar>

                </v-card>
            </v-flex>
        </v-layout>
    </v-container>
</template>

<script>
import axios from "axios";

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
        axios.get('http://localhost:8003/buyers')
        .then(response => this.buyers = response.data)
    },
    methods:{
        addFiles(){
            //console.log(this.files)
            if(this.files.length === 0){
                this.snackbar = true
                this.text = "No se ha detectado archivos"
            }else{
                this.files.forEach(i => {
                    const formData = new FormData()
                    formData.append('data', i)
                    console.log(i)
                    axios.post("http://localhost:8003/"+i.name, formData)
                })
                this.files = []
                this.snackbar = false 
                this.snackbar = true
                this.text = "Archivos subidos correctamente"
            }
        },
        consultar(id){
            console.log(id)
        }
    }
}
</script>