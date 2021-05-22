<template>
    <v-container grid-list-xl>
        <v-row justify="center">
            <v-col cols="12" sm="6" md="4">
      <v-dialog
        ref="dialog"
        v-model="modal"
        :return-value.sync="date"
        persistent
        width="290px"
      >
        <template v-slot:activator="{ on, attrs }">
          <v-text-field
            v-model="date"
            label="Selecciona una fecha"
            readonly
            v-bind="attrs"
            v-on="on"
          ></v-text-field>
        </template>
        <v-date-picker v-model="date" scrollable>
          <v-spacer></v-spacer>
          <v-btn text color="primary" @click="modal = false">Cancelar</v-btn>
          <v-btn text color="primary" @click="$refs.dialog.save(date)">OK</v-btn>
        </v-date-picker>
      </v-dialog>
    </v-col>
        </v-row>
        <v-layout row wrap justify-center>
            <v-flex md6>
                <v-card class="mb-3" >
                    <!--v-form method="post" action="http://localhost:8002/buyers"> -->
                    <v-form @submit.prevent="addFiles"> 
                        <v-file-input
                            v-model="files"
                            color="deep-purple accent-4"
                            counter
                            label="Adjunte los archivos"
                            multiple
                            placeholder="Selecciona tus archivos"
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
            modal: false,
            menu: false,
            date: new Date().toISOString().substr(0, 10),
            info : "",
            files: [],
            snackbar: false,
            text: ""
        }
    },
    methods:{
        addFiles(){
            if(this.files.length === 0){
                this.snackbar = true
                this.text = "No se ha detectado archivos"
            }else{
                this.files.forEach(i => {
                    const date = (new Date(this.picker)).getTime()
                    const formData = new FormData()
                    formData.append('data', i)
                    formData.append('date', date)
                    console.log(i)
                    axios.post("http://localhost:8003/"+i.name, formData)
                })
                this.files = []
                this.snackbar = false 
                this.snackbar = true
                this.text = "Archivos subidos correctamente"
            }
        }
    }
}
</script>