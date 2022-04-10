<template>
  <el-table :data="users" style="width: 100%">
    <el-table-column type="expand">
      <template #default="props">
        <div class="data">
          <el-table border :data="props.row.courses" >
            <el-table-column label="ID" prop="id" />
            <el-table-column label="名称" prop="name" />
            <el-table-column label="进度" prop="progress" />
            <el-table-column label="链接" prop="link" />

          </el-table>
        </div>

      </template>
    </el-table-column>
    <el-table-column label="ID" prop="Id" />
    <el-table-column label="姓名" prop="name" />
    <el-table-column label="学号" prop="student_id" />
    <el-table-column label="班级" prop="grade" />
    <el-table-column label="Operations">
      <template #header>
        <el-button type="primary" @click="flush()">刷新</el-button>
      </template>
      <template #default="scope">
        <el-button @click="delete_user(scope.row.Id)">删除用户</el-button>

      </template>
    </el-table-column>
    <el-table-column label="平台">
      <template #default="scope">
        <span v-if="scope.row.State===0">慕课平台</span>
        <span v-else>实训平台</span>
      </template>
    </el-table-column>
  </el-table>
</template>

<script>
import config from "../config/config";
import axios from "axios";


export default {
  name: "UserManager",
  data(){
    return{
        users : [

        ],
      a: async ()=>{
        this.users = []
        let resp = await axios.post(config.baseUrl+config.Api.GetLoginUser.url)
        let con = resp.data.data
        for (let i = 0; i < con.length; i++) {
          let resp1 = await axios.post(config.baseUrl+config.Api.Info.url+"?id="+con[i].Id)
          con[i].name = resp1.data.name
          con[i].student_id = resp1.data.student_id
          con[i].grade = resp1.data.grade
          this.users.push(con[i])
          let resp2 = await axios.post(config.baseUrl+config.Api.GetCourses.url+"?id="+con[i].Id)
          con[i].courses = resp2.data.data
        }
      }
    }
  },
  created() {
    this.a()
  },
  methods:{
    flush: function () {
        this.a()
    },
    delete_user:function (id) {
      console.log("删除用户"+id)
        config.DeleteUser(id).then(()=>{
          this.flush()
        })

    },

  },
  computed:{

  }

}
</script>

<style scoped>

.data{
  background-color: aquamarine;
}

</style>
