import axios from "axios";

export default {
    baseUrl : "http://127.0.0.1:13450",
    Api : {
        Login:{
            url:"/min/login"
        },
        LoginWeixin : {
            url: "/min/logician"
        },
        CheckQrCode : {
            url: "/min/check_qr_code"
        },
        GetLoginUser:{
            url:"/admin/get_login_user"
        },
        Info:{
            url:"/min/info"
        },
        GetCourseList:{
            url:"/min/get_course_list"
        },
        GetCourses:{
            url:"/min/get_courses"
        },
    },
    // 执行一个任务
    DoActivity: async function(userId,courseId){
        let resp = await axios.post(this.baseUrl+"/active/do_active",{user_id:userId,course_id:courseId})
        return resp.data
    },
    GetActives : async function(){
        let resp = await axios.post(this.baseUrl+"/active/get_actives")
        return resp.data
    }
}