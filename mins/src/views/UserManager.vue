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

        ]
    }
  },
  created() {
    let a = async ()=>{
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
    a()
      // this.Axios.post(config.baseUrl+config.Api.GetLoginUser.url).then((resp)=>{
      //   let con = resp.data.data
      //   for (let i=0;i<con.length;i++) {
      //     this.Axios.post(config.baseUrl+config.Api.Info.url+"?id="+con[i].Id).then((resp)=>{
      //       console.log(resp.data)
      //         con[i].name = resp.data.name
      //         con[i].student_id = resp.data.student_id
      //         con[i].grade = resp.data.grade
      //     })
      //     this.Axios.post(config.baseUrl+config.Api.GetCourses.url+"?id="+con[i].Id).then((resp)=>{
      //       con[i].courses = resp.data.data
      //     })
      //   }
      //   for (let i = 0; i < con.length; i++) {
      //     this.users.push(con[i])
      //   }
      // })
  },

}
</script>

<style scoped>

.data{
  background-color: aquamarine;
}

</style>