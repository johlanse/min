<template>
  <el-container :style="container_style">
    <el-aside :style="aside_width">
      <div
          style="height: 6%"
          class="label_style bor"
      >
        Mins
      </div>
      <el-menu
          router="router"
          class="el-menu-demo"
          background-color="#545c64"
          text-color="#fff"
          active-text-color="#ffd04b"
      >
        <el-menu-item index="add_user">
          添加用户
        </el-menu-item>
        <el-menu-item index="user_manager">
          用户管理
        </el-menu-item>
        <el-menu-item index="activity">
          任务管理
        </el-menu-item>
<!--        <el-menu-item index="requests">-->
<!--          请求事件-->
<!--        </el-menu-item>-->
<!--        <el-menu-item index="meta">-->
<!--          元事件-->
<!--        </el-menu-item>-->

      </el-menu>

    </el-aside>
    <el-container style="padding: 0">
<!--      <el-header :style="header_height">-->
<!--        <el-menu-->
<!--            router="router"-->
<!--            class="el-menu-demo"-->
<!--            mode="horizontal"-->
<!--            background-color="#545c64"-->
<!--            text-color="#fff"-->
<!--            active-text-color="#ffd04b"-->
<!--        >-->
<!--          -->
<!--&lt;!&ndash;          <el-menu-item index="config">&ndash;&gt;-->
<!--&lt;!&ndash;            配置管理&ndash;&gt;-->
<!--&lt;!&ndash;          </el-menu-item>&ndash;&gt;-->
<!--&lt;!&ndash;          <el-menu-item index="suggest">&ndash;&gt;-->
<!--&lt;!&ndash;            <el-link type="info" href="#">意见反馈</el-link>&ndash;&gt;-->
<!--&lt;!&ndash;          </el-menu-item>&ndash;&gt;-->
<!--        </el-menu>-->
<!--      </el-header>-->
      <el-main>
        <router-view/>
      </el-main>
      <el-footer class="bor">
        <el-form ref="form" >
          <el-form-item label="">
            <el-col style="width: 25%;margin-left: 3%;margin-top: 10px;">
              <el-input v-model="user_id" placeholder="请输入用户ID"></el-input>
            </el-col>

          <el-col style="width: 25%;margin-left: 3%;margin-top: 10px;">
            <el-input v-model="course_id" placeholder="请输入课程ID">

            </el-input>
          </el-col>

            <el-col style="width: 30%;margin-left: 3%;margin-top: 10px;">
              <el-button @click="do_active" type="success">提交任务</el-button>
            </el-col>


          </el-form-item>
        </el-form>

      </el-footer>
    </el-container>
  </el-container>
</template>


<script>
import Foo from "@/components/Foo";
import Config from "./config/config"
import {ElMessage} from "element-plus";


export default {
  components:[Foo],
  data(){
    return {
      label:"MINs",
      container_style :{
        margin: "0 auto",
        height : "100%",
      },
      header_height:{
        padding: 0,
        backgroundColor: "#545c64",
        height: "6%"
      },
      aside_width:{
        backgroundColor: "#545c64",
        padding: 0,
        width: "12%"
      },
      link:"",

      user_id: "",
      course_id: ""


    }
  },
  computed : {

  },
  methods:{
    do_active: function () {
      Config.DoActivity(Number(this.user_id),Number(this.course_id)).then(data =>{
        if (data.code === 1200){
          ElMessage.info("提交成功")
          this.course_id = ""
          this.user_id = ""
        }
      })
    }
  }

}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  height: 100%;
}
html, body {
  margin: 0 auto;
  padding: 0;
  height: 100%;
}
.label_style {
  font-size: large;
  font-family: "Microsoft YaHei UI Light",serif;
  font-weight: bold;
  color: #42b983;
}

#content{

  width: 100%;
  height: 100%;
}



</style>
