<template>
  <div >
    <el-form ref="form" :model="form" label-width="120px">
      <el-form-item label="学校">
        <el-select v-model="status" style="width: 100%">
<!--          <el-option label="燕京理工学院" value="2">燕京理工学院</el-option>-->
<!--          <el-option value="3">首都师范大学科德学院</el-option>-->
<!--          <el-option value="4">北京工商大学嘉华学院</el-option>-->
<!--          <el-option value="5">武汉工程科技学院</el-option>-->
<!--          <el-option value="6">昆明城市学院</el-option>-->
<!--          <el-option value="7">南宁理工学院</el-option>-->
<!--          <el-option value="8">云南城市建设职业学院</el-option>-->
<!--          <el-option value="9">上海立达学院</el-option>-->
<!--          <el-option value="10" >华南农业大学珠江学院</el-option>-->
<!--          <el-option value="11">银川科技学院</el-option>-->
<!--          <el-option value="13">重庆城市科技学院</el-option>-->
<!--          <el-option value="14">哈尔滨广厦学院</el-option>-->
<!--          <el-option value="15">中南林业科技大学涉外学院</el-option>-->
<!--          <el-option value="20">明达职业技术学院</el-option>-->
<!--          <el-option value="21">北方出租</el-option>-->
<!--          <el-option value="22">南京航空航天大学金城学院</el-option>-->
<!--          <el-option value="23">西安工商学院</el-option>-->

<!--          <el-option value="25">云南艺术学院文华学院</el-option>-->
<!--          <el-option value="26">延安大学西安创新学院</el-option>-->
<!--          <el-option value="27">温州商学院</el-option>-->
<!--          <el-option value="28">四川应用技术职业学院</el-option>-->
<!--          <el-option value="29">国际化大课堂</el-option>-->
<!--          <el-option value="30">那曲市职业技术学校</el-option>-->
<!--          <el-option value="31">广西北方出租汽车有限责任公司</el-option>-->
<!--          <el-option value="32">职业技能培训</el-option>-->
<!--          <el-option value="33">实习实训平台</el-option>-->
<!--          <el-option value="34">银川科技学院实训平台</el-option>-->
<!--          <el-option value="35">燕京理工学院实训平台</el-option>-->
<!--          <el-option value="36">重庆城市科技学院实训平台</el-option>-->
<!--          <el-option value="37">北京工商大学嘉华学院实训平台</el-option>-->
<!--          <el-option value="38">华南农业大学珠江学院实训平台</el-option>-->
<!--          <el-option value="39">中南林业科技大学涉外学院实训平台</el-option>-->
<!--          <el-option value="40">首都师范大学科德学院实训平台</el-option>-->
<!--          <el-option value="41">昆明城市学院实训平台</el-option>-->
<!--          <el-option value="42">南宁理工学院实训平台</el-option>-->
<!--          <el-option value="43">武汉工程科技学院实训平台</el-option>-->
<!--          <el-option value="44">云南城市建设职业学院实训平台</el-option>-->
<!--          <el-option value="45">上海立达学院实训平台</el-option>-->
<!--          <el-option value="46">哈尔滨广厦学院实训平台</el-option>-->
          <el-option value="0" label="成都文理学院慕课平台">成都文理学院慕课平台</el-option>
          <el-option value="1" label="成都文理学院实训平台">成都文理学院实训平台</el-option>
          <el-option value="2" label="成都文理学院创能平台">成都文理学院创能实训</el-option>
          <el-option value="3" label="成都文理学院频蓝实训">成都文理学院频蓝实训</el-option>
          <el-option value="4" label="成都文理学院戴希实训">成都文理学院戴希实训</el-option>
<!--          <el-option value="48">西安工商学院实训平台</el-option>-->
<!--          <el-option value="49">四川应用技术职业学院实训平台</el-option>-->
<!--          <el-option value="50">云南艺术学院文华学院实训平台</el-option>-->
<!--          <el-option value="51">南京航空航天大学金城学院实训平台</el-option>-->
<!--          <el-option value="53">温州商学院实训平台</el-option>-->
<!--          <el-option value="54">明达职业技术学院实训平台</el-option>-->
        </el-select>
      </el-form-item>
      <el-form-item label="账号">
        <el-input v-model="account"></el-input>
      </el-form-item>
      <el-form-item label="密码">
        <el-input v-model="password"></el-input>
      </el-form-item>

      <el-form-item label="备注">
        <el-input v-model="remark"></el-input>
      </el-form-item>

    </el-form>
    <el-form-item>
      <el-button @click="click" type="success">提交</el-button>
    </el-form-item>

  </div>
</template>

<script>

import Config from "../config/config";
import {ElMessage} from "element-plus";



export default {
  name: "AddUser",
  data(){
    return{
      link:"",
      id : 0,
      account : "",
      password : "",
      status :"0",
      message : "",
      remark: ""
    }
  },
  methods:{
    click : function () {
      this.Axios.post(Config.baseUrl+Config.Api.Login.url,{account:this.account,password:this.password,status:this.status,remark:this.remark}).then((resp)=>{
        console.log(resp.data)
        let con = resp.data.data
        if (con.data.status) {
          ElMessage.info("登录成功")
          this.account = ""
          this.password = ""
          this.id = con.id
          this.remark = ""
          // this.Axios.post(Config.baseUrl+Config.Api.LoginWeixin.url+"?id="+this.id).then(resp=>{
          //   console.log(resp.data)
          //   this.src = resp.data.data.url
          //   this.link = resp.data.data.link
          //   let i = setInterval(()=> {
          //     Axios.post(Config.baseUrl+Config.Api.CheckQrCode.url,{id:con.id,link:resp.data.data.link}).then((resp)=>{
          //       console.log(resp.data)
          //       if (resp.data.code===1200){
          //         console.log(resp.data)
          //         alert("已成功登录该账号")
          //         this.account = ""
          //         this.password = ""
          //         this.link = ""
          //         clearInterval(i)
          //       }
          //     })
          //   },5000)
          // })


        }
        // eslint-disable-next-line no-empty
        else {
          ElMessage.error(resp.data.data)
        }
      })
    }
  }
}
</script>

<style scoped>

.image{
  width: 300px;
  height: 300px;
}

</style>
