const {Delete} = require("@element-plus/icons-vue");
<template>
<div>
  <el-table :data="activities">
    <el-table-column prop="id" label="ID"/>
    <el-table-column prop="user_id" label="UserID"/>
    <el-table-column prop="course_id" label="CourseID" />
    <el-table-column prop="status" label="Status" />
    <el-table-column label="Action">
      <template #default="scope">

        <el-button type="success" @click="to_log(scope.row.id,scope.row.course_id)">启动</el-button>
        <el-button @click="to_log(scope.row.id,scope.row.course_id)">停止</el-button>
        <el-button type="danger" :icon="Delete" circle>删除</el-button>
      </template>
    </el-table-column>
    <el-table-column label="Operations">
      <template #default="scope">
        <el-button @click="to_log(scope.row.id,scope.row.course_id)">查看日志</el-button>

      </template>
    </el-table-column>
  </el-table>
</div>
</template>

<script>
import config from "@/config/config";
import {
  Search,
  Edit,
  Check,
  Message,
  Star,
  Delete,
} from '@element-plus/icons-vue'
export default {
  name: "Activity",
  components:[
    Search,
    Edit,
    Check,
    Message,
    Star,
    Delete,
  ],
  data(){
    return{
        activities: []
    }
  },
  created() {
    let a = async ()=>{
      let resp = await config.GetActives()

      for (let i = 0; i < resp.data.length; i++) {
        console.log(resp.data[i])
        this.activities.push(resp.data[i])
      }
    }
    a()
  },
  methods:{
      get_log_url:function (activity_id,course_id) {
        return config.baseUrl+"/log/"+activity_id+"_"+course_id+".log"
      },
    to_log:function (activity_id,course_id) {
      this.$router.push({name:"log",params:{link:this.get_log_url(activity_id,course_id)}})
    }
  },
  computed:{

  }
}
</script>

<style scoped>

</style>