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

        <el-button type="success" @click="do_active(scope.row.user_id,scope.row.course_id)">启动</el-button>
        <el-button @click="stop_active(scope.row.id)">停止</el-button>
        <el-button type="danger" @click="delete_active(scope.row.id)" :icon="Delete" circle>删除</el-button>
      </template>
    </el-table-column>
    <el-table-column label="Operations">
      <template #header>
        <el-button type="primary" @click="a()">刷新</el-button>
      </template>
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
  // Search,
  // Edit,
  // Check,
  // Message,
  // Star,
  Delete,

} from '@element-plus/icons-vue'
import {ElMessage} from "element-plus";
import Config from "@/config/config";
export default {
  name: "Activity",
  components:[
    // Search,
    // Edit,
    // Check,
    // Message,
    // Star,
    Delete,
  ],
  data(){
    return{
        activities: [],
      a : async ()=>{
          this.activities= []
        let resp = await config.GetActives()

        for (let i = 0; i < resp.data.length; i++) {
          this.activities.push(resp.data[i])
        }
      }
    }
  },
  created() {
    this.a()
  },
  methods:{
      get_log_url:function (activity_id,course_id) {
        return config.baseUrl+"/log/"+activity_id+"_"+course_id+".log"
      },
    to_log:function (activity_id,course_id) {
      this.$router.push({name:"log",params:{link:this.get_log_url(activity_id,course_id)}})
    },
    do_active:function (user_id,course_id) {
      Config.DoActivity(Number(user_id),Number(course_id)).then(data =>{
        if (data.code === 1200){
          ElMessage.info("提交成功")
          this.course_id = ""
          this.user_id = ""
        }
        this.a()
      })
    },
    stop_active:function (id) {
      config.StopActivity(id).then(data=>{
        if (data.code === 1200){
          ElMessage.info("操作成功")
          this.a()
        }
      })
    },
    delete_active:function (id) {
      config.DeleteActivity(id).then(data=>{
        if (data.code === 1200){
          ElMessage.info("操作成功")
          this.a()
        }
      })
    }
  }
}
</script>

<style scoped>

</style>
