<template>
  <div class="container atasnya">
    <div>
        Masukkan url secreto target 👇
    </div>
    <div>
        <div class="form-group">
            <div class="d-flex justify-content-center">
                    <textarea v-model="isinya.target" class="form-control" rows="4" id="comment" style="width:700px"></textarea>
            </div>
            <small>
                <i>
                URL target boleh lebih dari satu dengan cara ENTER (dibawahnya)
                </i>
            </small>                    
        </div>
    </div>
    <div style="padding-top: 20px">
        Masukkan pesan yang ingin disampaikan ke target ✉️
    </div>
    <div class="form-group">
            <div class="d-flex justify-content-center">
                    <textarea v-model="isinya.pesan" class="form-control" rows="3" id="comment" style="width:500px;"></textarea>
            </div>                  
    </div>
    <div style="padding-top: 20px">
        Masukkan jumlah pesan yang ingin dikirim 🔥
    </div>
    <div>
        <div class="form-group">
            <div class="d-flex justify-content-center">
                <input type="number" v-model="isinya.jumlah" class="form-control" id="usr" style="width:100px;">
            </div>
        </div>
    </div>
    <div style="padding-top: 20px">
        <span v-if="isOnDutty">
            On Dutty ...
        </span>
        <div v-if="isDone">
            <span>
            -- Selesai --
            </span> <br>
        </div>
        <button type="button" class="btn btn-warning" @click="bersihkan"  v-if="progress.length > 0 && isDone">Clear</button>  
        <button type="button" class="btn btn-primary" style="margin-left: 10px" @click="kirim" v-if="!isKirim">Kirim !</button> 
    </div>
    <div style="padding-top: 30px;" class="d-flex justify-content-center" v-if="progress.length > 0">
        <div style="overflow-y: scroll; height:200px;" class="scroll">
            <div v-for="(item, index) in progress" v-bind:key="index">
            <span>
                {{ item }}
            </span>
        </div>
        </div>
        
    </div>
  </div>
</template>

<script>
import {YukKirim} from '../../wailsjs/go/spamin/Spamin'
import {EventsOn} from '../../wailsjs/runtime/runtime'
export default {
    data() {
        return {
            isinya: {
                target: "",
                pesan: "",
                jumlah: ""
            },
            progress: [],
            isKirim: false,
            isOnDutty: false,
            isDone: false
        }
    },
    methods: {
        kirim() {
            // console.log(this.isinya);
            this.isKirim = true
            YukKirim(JSON.stringify(this.isinya, null, 2));
        },
        bersihkan() {
            this.progress = []
            this.isinya.pesan = ''
            this.isKirim = false
            this.isOnDutty = false
            this.isDone = false
        }
    },
    mounted() {
        EventsOn("pesannya", (pesan) => {
            this.progress.push(pesan)
        });
        EventsOn("mulai", () => {
            this.isOnDutty = true
        })
        EventsOn("selesai", () => {
            this.isOnDutty = false
            this.isDone = true
        })
    },
}
</script>

<style>
.atasnya{
    padding-top: 20px
}
.scroll{
    background-color: rgb(43, 62, 92);
    width: 700px;
}
</style>