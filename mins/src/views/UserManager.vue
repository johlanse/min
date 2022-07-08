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
            <el-table-column label="操作">
              <template #default="scope">
                <el-button @click="do_active(props.row.Id,scope.row.id)">添加任务</el-button>
              </template>
            </el-table-column>
          </el-table>
        </div>
      </template>
    </el-table-column>
    <el-table-column label="ID" prop="Id" />
    <el-table-column label="姓名" prop="name" />
    <el-table-column label="学号" prop="student_id" />
<!--    <el-table-column label="班级" prop="grade" />-->
    <el-table-column label="备注">
      <template #default="scope">
        <a v-if="_isMobile()" v-bind:href="'mqqwpa://im/chat?chat_type=wpa&uin='+scope.row.Remark+'&version=1&src_type=web&web_src=oicqzone.com'">{{scope.row.Remark}}</a>
        <a v-else v-bind:href="'tencent://message/?Menu=yes&uin=1743224847&Site=80fans&Service=300&sigT=45a1e5847943b64c6ff3990f8a9e644d2b31356cb0b4ac6b24663a3c8dd0f8aa12a545b1714f9d45'+scope.row.Remark+'&Site=80fans&Service=300&sigT=45a1e5847943b64c6ff3990f8a9e644d2b31356cb0b4ac6b24663a3c8dd0f8aa12a545b1714f9d45'">{{scope.row.Remark}}</a>
      </template>
    </el-table-column>
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
        <span v-else-if="scope.row.State===1">实训平台</span>
        <span v-else-if="scope.row.State===2">创能平台</span>
      </template>
    </el-table-column>
  </el-table>
</template>

<script>
import config from "../config/config";
import axios from "axios";
import Config from "@/config/config";
import {ElMessage} from "element-plus";


export default {
  name: "UserManager",
  data(){
    return{
        users : [

        ],
      a: async ()=>{
        this.users = []
        let temp = []
        let resp = await axios.post(config.baseUrl+config.Api.GetLoginUser.url)
        let con = resp.data.data
        for (let i = 0; i < con.length; i++) {
          let resp1 = await axios.post(config.baseUrl+config.Api.Info.url+"?id="+con[i].Id)
          con[i].name = resp1.data.name
          con[i].student_id = resp1.data.student_id
          con[i].grade = resp1.data.grade
          // this.users.push(con[i])
          temp.push(con[i])
          let resp2 = await axios.post(config.baseUrl+config.Api.GetCourses.url+"?id="+con[i].Id)
          con[i].courses = resp2.data.data
          this.users = temp
        }
      }
    }
  },
  created() {
    this.a()
  },
  methods:{

    do_active: function (user_id,course_id) {
      Config.DoActivity(user_id,course_id).then(data =>{
        if (data.code === 1200){
          ElMessage.info("提交成功")
        }
      })
    },

    flush: function () {
        this.a()
    },
    delete_user:function (id) {
      console.log("删除用户"+id)
        config.DeleteUser(id).then(()=>{
          this.flush()
        })

    },
    _isMobile() {
      return navigator.userAgent.match(/(phone|pad|pod|iPhone|iPod|ios|iPad|Android|Mobile|BlackBerry|IEMobile|MQQBrowser|JUC|Fennec|wOSBrowser|BrowserNG|WebOS|Symbian|Windows Phone)/i);
    }

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
