<template>
<div>


      <el-input
          v-model="content"
          autosize
          type="textarea"
          :disabled="true"
      >

      </el-input>

</div>
</template>

<script>
import axios from "axios";


export default {
  name: "Log",
  data(){
    return{
      content: "",
      link: "",
      interval:{},

      height: 0,
      scrollbar:{}
    }
  },
  mounted() {

  },
  created() {
    this.link = this.$route.params.link
    axios.get(this.link).then(resp=>{
      this.content = resp.data
    })

    this.interval = setInterval(()=>{
      axios.get(this.link).then(resp=>{
        this.content = resp.data

      })
    },10000)
  },
  methods:{

  },
  unmounted() {
    clearInterval(this.interval)
  }

}
</script>

<style scoped>

</style>