(function(e){function t(t){for(var r,u,o=t[0],i=t[1],l=t[2],s=0,d=[];s<o.length;s++)u=o[s],Object.prototype.hasOwnProperty.call(c,u)&&c[u]&&d.push(c[u][0]),c[u]=0;for(r in i)Object.prototype.hasOwnProperty.call(i,r)&&(e[r]=i[r]);b&&b(t);while(d.length)d.shift()();return a.push.apply(a,l||[]),n()}function n(){for(var e,t=0;t<a.length;t++){for(var n=a[t],r=!0,o=1;o<n.length;o++){var i=n[o];0!==c[i]&&(r=!1)}r&&(a.splice(t--,1),e=u(u.s=n[0]))}return e}var r={},c={app:0},a=[];function u(t){if(r[t])return r[t].exports;var n=r[t]={i:t,l:!1,exports:{}};return e[t].call(n.exports,n,n.exports,u),n.l=!0,n.exports}u.m=e,u.c=r,u.d=function(e,t,n){u.o(e,t)||Object.defineProperty(e,t,{enumerable:!0,get:n})},u.r=function(e){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(e,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(e,"__esModule",{value:!0})},u.t=function(e,t){if(1&t&&(e=u(e)),8&t)return e;if(4&t&&"object"===typeof e&&e&&e.__esModule)return e;var n=Object.create(null);if(u.r(n),Object.defineProperty(n,"default",{enumerable:!0,value:e}),2&t&&"string"!=typeof e)for(var r in e)u.d(n,r,function(t){return e[t]}.bind(null,r));return n},u.n=function(e){var t=e&&e.__esModule?function(){return e["default"]}:function(){return e};return u.d(t,"a",t),t},u.o=function(e,t){return Object.prototype.hasOwnProperty.call(e,t)},u.p="/static/dist/";var o=window["webpackJsonp"]=window["webpackJsonp"]||[],i=o.push.bind(o);o.push=t,o=o.slice();for(var l=0;l<o.length;l++)t(o[l]);var b=i;a.push([0,"chunk-vendors"]),n()})({0:function(e,t,n){e.exports=n("56d7")},"2b5c":function(e,t,n){"use strict";n("74d1")},"4bb3":function(e,t,n){"use strict";n("7638")},"56d7":function(e,t,n){"use strict";n.r(t);n("e260"),n("e6cf"),n("cca6"),n("a79d"),n("d3b7");var r=n("7a23"),c=Object(r["n"])("div",{style:{height:"6%"},class:"label_style bor"}," Mins ",-1),a=Object(r["p"])(" 添加用户 "),u=Object(r["p"])(" 用户管理 "),o=Object(r["p"])(" 任务管理 "),i=Object(r["p"])("提交任务"),l=Object(r["p"])("刷新token");function b(e,t,n,b,s,d){var f=Object(r["R"])("el-menu-item"),p=Object(r["R"])("el-menu"),O=Object(r["R"])("el-aside"),j=Object(r["R"])("router-view"),m=Object(r["R"])("el-main"),g=Object(r["R"])("el-input"),h=Object(r["R"])("el-col"),_=Object(r["R"])("el-button"),v=Object(r["R"])("el-form-item"),w=Object(r["R"])("el-form"),k=Object(r["R"])("el-footer"),y=Object(r["R"])("el-container");return Object(r["I"])(),Object(r["k"])(y,{style:Object(r["z"])(s.container_style)},{default:Object(r["gb"])((function(){return[Object(r["q"])(O,{style:Object(r["z"])(s.aside_width)},{default:Object(r["gb"])((function(){return[c,Object(r["q"])(p,{router:"router",class:"el-menu-demo","background-color":"#545c64","text-color":"#fff","active-text-color":"#ffd04b"},{default:Object(r["gb"])((function(){return[Object(r["q"])(f,{index:"add_user"},{default:Object(r["gb"])((function(){return[a]})),_:1}),Object(r["q"])(f,{index:"user_manager"},{default:Object(r["gb"])((function(){return[u]})),_:1}),Object(r["q"])(f,{index:"activity"},{default:Object(r["gb"])((function(){return[o]})),_:1})]})),_:1})]})),_:1},8,["style"]),Object(r["q"])(y,{style:{padding:"0"}},{default:Object(r["gb"])((function(){return[Object(r["q"])(m,null,{default:Object(r["gb"])((function(){return[Object(r["q"])(j)]})),_:1}),Object(r["q"])(k,{class:"bor"},{default:Object(r["gb"])((function(){return[Object(r["q"])(w,{ref:"form"},{default:Object(r["gb"])((function(){return[Object(r["q"])(v,{label:""},{default:Object(r["gb"])((function(){return[Object(r["q"])(h,{style:{width:"25%","margin-left":"3%","margin-top":"10px"}},{default:Object(r["gb"])((function(){return[Object(r["q"])(g,{modelValue:s.user_id,"onUpdate:modelValue":t[0]||(t[0]=function(e){return s.user_id=e}),placeholder:"请输入用户ID"},null,8,["modelValue"])]})),_:1}),Object(r["q"])(h,{style:{width:"25%","margin-left":"3%","margin-top":"10px"}},{default:Object(r["gb"])((function(){return[Object(r["q"])(g,{modelValue:s.course_id,"onUpdate:modelValue":t[1]||(t[1]=function(e){return s.course_id=e}),placeholder:"请输入课程ID"},null,8,["modelValue"])]})),_:1}),Object(r["q"])(h,{style:{width:"30%","margin-left":"3%","margin-top":"10px"}},{default:Object(r["gb"])((function(){return[Object(r["q"])(_,{onClick:d.do_active,type:"success"},{default:Object(r["gb"])((function(){return[i]})),_:1},8,["onClick"])]})),_:1}),Object(r["q"])(h,null,{default:Object(r["gb"])((function(){return[Object(r["q"])(_,{type:"danger",onClick:d.flush_token},{default:Object(r["gb"])((function(){return[l]})),_:1},8,["onClick"])]})),_:1})]})),_:1})]})),_:1},512)]})),_:1})]})),_:1})]})),_:1},8,["style"])}n("a9e3");function s(e,t,n,c,a,u){var o=Object(r["R"])("el-input"),i=Object(r["R"])("el-form-item"),l=Object(r["R"])("el-form");return Object(r["I"])(),Object(r["k"])(l,{ref:"form"},{default:Object(r["gb"])((function(){return[Object(r["q"])(i,{label:"12"},{default:Object(r["gb"])((function(){return[Object(r["q"])(o,{placeholder:"请输入用户ID"}),Object(r["q"])(o,{placeholder:"请输入课程ID"})]})),_:1})]})),_:1},512)}var d={name:"Foo",data:function(){return{}}},f=n("6b0d"),p=n.n(f);const O=p()(d,[["render",s]]);var j=O,m=n("1da1"),g=(n("96cf"),n("bc3a")),h=n.n(g),_={baseUrl:"",Api:{Login:{url:"/min/login"},LoginWeixin:{url:"/min/logician"},CheckQrCode:{url:"/min/check_qr_code"},GetLoginUser:{url:"/admin/get_login_user"},Info:{url:"/min/info"},GetCourseList:{url:"/min/get_course_list"},GetCourses:{url:"/min/get_courses"}},DoActivity:function(){var e=Object(m["a"])(regeneratorRuntime.mark((function e(t,n){var r;return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:return e.next=2,h.a.post(this.baseUrl+"/active/do_active",{user_id:t,course_id:n});case 2:return r=e.sent,e.abrupt("return",r.data);case 4:case"end":return e.stop()}}),e,this)})));function t(t,n){return e.apply(this,arguments)}return t}(),GetActives:function(){var e=Object(m["a"])(regeneratorRuntime.mark((function e(){var t;return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:return e.next=2,h.a.post(this.baseUrl+"/active/get_actives");case 2:return t=e.sent,e.abrupt("return",t.data);case 4:case"end":return e.stop()}}),e,this)})));function t(){return e.apply(this,arguments)}return t}(),StopActivity:function(){var e=Object(m["a"])(regeneratorRuntime.mark((function e(t){var n;return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:return e.next=2,h.a.post(this.baseUrl+"/active/stop_active?active_id="+t);case 2:return n=e.sent,e.abrupt("return",n.data);case 4:case"end":return e.stop()}}),e,this)})));function t(t){return e.apply(this,arguments)}return t}(),DeleteActivity:function(){var e=Object(m["a"])(regeneratorRuntime.mark((function e(t){var n;return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:return e.next=2,h.a.post(this.baseUrl+"/active/delete_active?active_id="+t);case 2:return n=e.sent,e.abrupt("return",n.data);case 4:case"end":return e.stop()}}),e,this)})));function t(t){return e.apply(this,arguments)}return t}(),DeleteUser:function(){var e=Object(m["a"])(regeneratorRuntime.mark((function e(t){var n;return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:return e.next=2,h.a.post(this.baseUrl+"/admin/delete_user?id="+t);case 2:return n=e.sent,e.abrupt("return",n.data);case 4:case"end":return e.stop()}}),e,this)})));function t(t){return e.apply(this,arguments)}return t}(),GetLoginUser:function(){var e=Object(m["a"])(regeneratorRuntime.mark((function e(){var t;return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:return e.next=2,h.a.post(this.baseUrl+"/admin/get_login_user");case 2:return t=e.sent,e.abrupt("return",t.data);case 4:case"end":return e.stop()}}),e,this)})));function t(){return e.apply(this,arguments)}return t}()},v=n("7864"),w={components:[j,v["b"]],data:function(){return{label:"MINs",container_style:{margin:"0 auto",height:"100%"},header_height:{padding:0,backgroundColor:"#545c64",height:"6%"},aside_width:{backgroundColor:"#545c64",padding:0,width:"12%"},link:"",user_id:"",course_id:""}},computed:{},methods:{do_active:function(){var e=this;_.DoActivity(Number(this.user_id),Number(this.course_id)).then((function(t){1200===t.code&&(v["a"].info("提交成功"),e.course_id="",e.user_id="")}))},flush_token:function(){v["b"].prompt("请输入页面token","提示").then((function(e){null!==e&&window.localStorage.setItem("min_token",e.value)}))}}};n("ba1f");const k=p()(w,[["render",b]]);var y=k,q=n("6c02"),R=Object(r["p"])("成都文理学院慕课平台"),x=Object(r["p"])("成都文理学院实训平台"),I=Object(r["p"])("成都文理学院创能实训"),C=Object(r["p"])("成都文理学院频蓝实训"),U=Object(r["p"])("成都文理学院戴希实训"),S=Object(r["p"])("提交");function V(e,t,n,c,a,u){var o=Object(r["R"])("el-option"),i=Object(r["R"])("el-select"),l=Object(r["R"])("el-form-item"),b=Object(r["R"])("el-input"),s=Object(r["R"])("el-form"),d=Object(r["R"])("el-button");return Object(r["I"])(),Object(r["m"])("div",null,[Object(r["q"])(s,{ref:"form",model:e.form,"label-width":"120px"},{default:Object(r["gb"])((function(){return[Object(r["q"])(l,{label:"学校"},{default:Object(r["gb"])((function(){return[Object(r["q"])(i,{modelValue:a.status,"onUpdate:modelValue":t[0]||(t[0]=function(e){return a.status=e}),style:{width:"100%"}},{default:Object(r["gb"])((function(){return[Object(r["q"])(o,{value:"0",label:"成都文理学院慕课平台"},{default:Object(r["gb"])((function(){return[R]})),_:1}),Object(r["q"])(o,{value:"1",label:"成都文理学院实训平台"},{default:Object(r["gb"])((function(){return[x]})),_:1}),Object(r["q"])(o,{value:"2",label:"成都文理学院创能平台"},{default:Object(r["gb"])((function(){return[I]})),_:1}),Object(r["q"])(o,{value:"3",label:"成都文理学院频蓝实训"},{default:Object(r["gb"])((function(){return[C]})),_:1}),Object(r["q"])(o,{value:"4",label:"成都文理学院戴希实训"},{default:Object(r["gb"])((function(){return[U]})),_:1})]})),_:1},8,["modelValue"])]})),_:1}),Object(r["q"])(l,{label:"账号"},{default:Object(r["gb"])((function(){return[Object(r["q"])(b,{modelValue:a.account,"onUpdate:modelValue":t[1]||(t[1]=function(e){return a.account=e})},null,8,["modelValue"])]})),_:1}),Object(r["q"])(l,{label:"密码"},{default:Object(r["gb"])((function(){return[Object(r["q"])(b,{modelValue:a.password,"onUpdate:modelValue":t[2]||(t[2]=function(e){return a.password=e})},null,8,["modelValue"])]})),_:1}),Object(r["q"])(l,{label:"备注"},{default:Object(r["gb"])((function(){return[Object(r["q"])(b,{modelValue:a.remark,"onUpdate:modelValue":t[3]||(t[3]=function(e){return a.remark=e})},null,8,["modelValue"])]})),_:1})]})),_:1},8,["model"]),Object(r["q"])(l,null,{default:Object(r["gb"])((function(){return[Object(r["q"])(d,{onClick:u.click,type:"success"},{default:Object(r["gb"])((function(){return[S]})),_:1},8,["onClick"])]})),_:1})])}var A={name:"AddUser",data:function(){return{link:"",id:0,account:"",password:"",status:"0",message:"",remark:""}},methods:{click:function(){var e=this;this.Axios.post(_.baseUrl+_.Api.Login.url,{account:this.account,password:this.password,status:this.status,remark:this.remark}).then((function(t){console.log(t.data);var n=t.data.data;n.data.status?(v["a"].info("登录成功"),e.account="",e.password="",e.id=n.id,e.remark=""):v["a"].error(t.data.data)}))}}};n("2b5c");const D=p()(A,[["render",V],["__scopeId","data-v-17e6ad44"]]);var M=D,P={class:"data"},G=Object(r["p"])("添加任务"),L=["href"],N=["href"],B=Object(r["p"])("刷新"),z=Object(r["p"])("删除用户"),T={key:0},J={key:1},Q={key:2},W={key:3},F={key:4};function $(e,t,n,c,a,u){var o=Object(r["R"])("el-table-column"),i=Object(r["R"])("el-button"),l=Object(r["R"])("el-table");return Object(r["I"])(),Object(r["k"])(l,{data:a.users,style:{width:"100%"}},{default:Object(r["gb"])((function(){return[Object(r["q"])(o,{type:"expand"},{default:Object(r["gb"])((function(e){return[Object(r["n"])("div",P,[Object(r["q"])(l,{border:"",data:e.row.courses},{default:Object(r["gb"])((function(){return[Object(r["q"])(o,{label:"ID",prop:"id"}),Object(r["q"])(o,{label:"名称",prop:"name"}),Object(r["q"])(o,{label:"进度",prop:"progress"}),Object(r["q"])(o,{label:"链接",prop:"link"}),Object(r["q"])(o,{label:"操作"},{default:Object(r["gb"])((function(t){return[Object(r["q"])(i,{onClick:function(n){return u.do_active(e.row.Id,t.row.id)}},{default:Object(r["gb"])((function(){return[G]})),_:2},1032,["onClick"])]})),_:2},1024)]})),_:2},1032,["data"])])]})),_:1}),Object(r["q"])(o,{label:"ID",prop:"Id"}),Object(r["q"])(o,{label:"姓名",prop:"name"}),Object(r["q"])(o,{label:"学号",prop:"student_id"}),Object(r["q"])(o,{label:"备注"},{default:Object(r["gb"])((function(e){return[u._isMobile()?(Object(r["I"])(),Object(r["m"])("a",{key:0,href:"mqqwpa://im/chat?chat_type=wpa&uin="+e.row.Remark+"&version=1&src_type=web&web_src=oicqzone.com"},Object(r["V"])(e.row.Remark),9,L)):(Object(r["I"])(),Object(r["m"])("a",{key:1,href:"tencent://message/?Menu=yes&uin="+e.row.Remark+"&Site=80fans&Service=300&sigT=45a1e5847943b64c6ff3990f8a9e644d2b31356cb0b4ac6b24663a3c8dd0f8aa12a545b1714f9d45"+e.row.Remark+"&Site=80fans&Service=300&sigT=45a1e5847943b64c6ff3990f8a9e644d2b31356cb0b4ac6b24663a3c8dd0f8aa12a545b1714f9d45"},Object(r["V"])(e.row.Remark),9,N))]})),_:1}),Object(r["q"])(o,{label:"Operations"},{header:Object(r["gb"])((function(){return[Object(r["q"])(i,{type:"primary",onClick:t[0]||(t[0]=function(e){return u.flush()})},{default:Object(r["gb"])((function(){return[B]})),_:1})]})),default:Object(r["gb"])((function(e){return[Object(r["q"])(i,{onClick:function(t){return u.delete_user(e.row.Id)}},{default:Object(r["gb"])((function(){return[z]})),_:2},1032,["onClick"])]})),_:1}),Object(r["q"])(o,{label:"平台"},{default:Object(r["gb"])((function(e){return[0===e.row.State?(Object(r["I"])(),Object(r["m"])("span",T,"慕课平台")):1===e.row.State?(Object(r["I"])(),Object(r["m"])("span",J,"实训平台")):2===e.row.State?(Object(r["I"])(),Object(r["m"])("span",Q,"创能平台")):3===e.row.State?(Object(r["I"])(),Object(r["m"])("span",W,"频蓝实训")):4===e.row.State?(Object(r["I"])(),Object(r["m"])("span",F,"戴希实训")):Object(r["l"])("",!0)]})),_:1})]})),_:1},8,["data"])}n("b0c0"),n("ac1f"),n("466d");var E={name:"UserManager",data:function(){var e=this;return{users:[],a:function(){var t=Object(m["a"])(regeneratorRuntime.mark((function t(){var n,r,c,a,u,o;return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:return e.users=[],n=[],t.next=4,h.a.post(_.baseUrl+_.Api.GetLoginUser.url);case 4:r=t.sent,c=r.data.data,a=0;case 7:if(!(a<c.length)){t.next=23;break}return t.next=10,h.a.post(_.baseUrl+_.Api.Info.url+"?id="+c[a].Id);case 10:return u=t.sent,c[a].name=u.data.name,c[a].student_id=u.data.student_id,c[a].grade=u.data.grade,n.push(c[a]),t.next=17,h.a.post(_.baseUrl+_.Api.GetCourses.url+"?id="+c[a].Id);case 17:o=t.sent,c[a].courses=o.data.data,e.users=n;case 20:a++,t.next=7;break;case 23:case"end":return t.stop()}}),t)})));function n(){return t.apply(this,arguments)}return n}()}},created:function(){this.a()},methods:{do_active:function(e,t){_.DoActivity(e,t).then((function(e){1200===e.code&&v["a"].info("提交成功")}))},flush:function(){this.a()},delete_user:function(e){var t=this;console.log("删除用户"+e),_.DeleteUser(e).then((function(){t.flush()}))},_isMobile:function(){return navigator.userAgent.match(/(phone|pad|pod|iPhone|iPod|ios|iPad|Android|Mobile|BlackBerry|IEMobile|MQQBrowser|JUC|Fennec|wOSBrowser|BrowserNG|WebOS|Symbian|Windows Phone)/i)}},computed:{}};n("4bb3");const H=p()(E,[["render",$],["__scopeId","data-v-395dd37e"]]);var K=H,X=Object(r["p"])("启动"),Y=Object(r["p"])("停止"),Z=Object(r["p"])("删除"),ee=Object(r["p"])("刷新"),te=Object(r["p"])("查看日志");function ne(e,t,n,c,a,u){var o=Object(r["R"])("el-table-column"),i=Object(r["R"])("el-button"),l=Object(r["R"])("el-table");return Object(r["I"])(),Object(r["m"])("div",null,[Object(r["q"])(l,{data:a.activities},{default:Object(r["gb"])((function(){return[Object(r["q"])(o,{prop:"id",label:"ID"}),Object(r["q"])(o,{prop:"user_id",label:"UserID"}),Object(r["q"])(o,{prop:"course_id",label:"CourseID"}),Object(r["q"])(o,{prop:"status",label:"Status"}),Object(r["q"])(o,{label:"Action"},{default:Object(r["gb"])((function(t){return[Object(r["q"])(i,{type:"success",onClick:function(e){return u.do_active(t.row.user_id,t.row.course_id)}},{default:Object(r["gb"])((function(){return[X]})),_:2},1032,["onClick"]),Object(r["q"])(i,{onClick:function(e){return u.stop_active(t.row.id)}},{default:Object(r["gb"])((function(){return[Y]})),_:2},1032,["onClick"]),Object(r["q"])(i,{type:"danger",onClick:function(e){return u.delete_active(t.row.id)},icon:e.Delete,circle:""},{default:Object(r["gb"])((function(){return[Z]})),_:2},1032,["onClick","icon"])]})),_:1}),Object(r["q"])(o,{label:"Operations"},{header:Object(r["gb"])((function(){return[Object(r["q"])(i,{type:"primary",onClick:t[0]||(t[0]=function(e){return a.a()})},{default:Object(r["gb"])((function(){return[ee]})),_:1})]})),default:Object(r["gb"])((function(e){return[Object(r["q"])(i,{onClick:function(t){return u.to_log(e.row.id,e.row.course_id)}},{default:Object(r["gb"])((function(){return[te]})),_:2},1032,["onClick"])]})),_:1})]})),_:1},8,["data"])])}var re=n("df9f"),ce={name:"Activity",components:[re["a"]],data:function(){var e=this;return{activities:[],a:function(){var t=Object(m["a"])(regeneratorRuntime.mark((function t(){var n,r;return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:return e.activities=[],t.next=3,_.GetActives();case 3:for(n=t.sent,r=0;r<n.data.length;r++)e.activities.push(n.data[r]);case 5:case"end":return t.stop()}}),t)})));function n(){return t.apply(this,arguments)}return n}()}},created:function(){this.a()},methods:{get_log_url:function(e,t){return _.baseUrl+"/log/"+e+"_"+t+".log"},to_log:function(e,t){this.$router.push({name:"log",params:{link:this.get_log_url(e,t)}})},do_active:function(e,t){var n=this;_.DoActivity(Number(e),Number(t)).then((function(e){1200===e.code&&(v["a"].info("提交成功"),n.course_id="",n.user_id=""),n.a()}))},stop_active:function(e){var t=this;_.StopActivity(e).then((function(e){1200===e.code&&(v["a"].info("操作成功"),t.a())}))},delete_active:function(e){var t=this;_.DeleteActivity(e).then((function(e){1200===e.code&&(v["a"].info("操作成功"),t.a())}))}}};const ae=p()(ce,[["render",ne]]);var ue=ae;function oe(e,t,n,c,a,u){var o=Object(r["R"])("el-input");return Object(r["I"])(),Object(r["m"])("div",null,[Object(r["q"])(o,{modelValue:a.content,"onUpdate:modelValue":t[0]||(t[0]=function(e){return a.content=e}),autosize:"",type:"textarea",disabled:!0},null,8,["modelValue"])])}n("9911"),n("a15b"),n("1276");var ie={name:"Log",data:function(){return{content:"",link:"",interval:{},height:0,scrollbar:{}}},mounted:function(){},created:function(){var e=this;this.link=this.$route.params.link,h.a.get(this.link).then((function(t){e.content=t.data.split().reverse().join()})),this.interval=setInterval((function(){h.a.get(e.link).then((function(t){e.content=t.data.split().reverse().join()}))}),1e4)},methods:{},unmounted:function(){clearInterval(this.interval)}};const le=p()(ie,[["render",oe]]);var be=le,se=[{path:"/add_user",name:"add_user",component:M},{path:"/user_manager",name:"user_manager",component:K},{path:"/activity",name:"activity",component:ue},{path:"/log",name:"log",component:be}],de=Object(q["a"])({history:Object(q["b"])("/static/dist/"),routes:se}),fe=de,pe=(n("c69f"),n("3ef0")),Oe=n.n(pe),je=function(e){e.use(v["c"],{locale:Oe.a})},me=Object(r["j"])(y);me.config.globalProperties.Axios=h.a,je(me),me.use(fe).mount("#app");var ge=window.localStorage.getItem("min_token");null===ge&&v["b"].prompt("请输入页面token","提示").then((function(e){null!==e&&window.localStorage.setItem("min_token",e.value)})),h.a.interceptors.request.use((function(e){return console.log(e.url),e.headers.token=window.localStorage.getItem("min_token"),e}),(function(e){return Promise.reject(e)}))},"74d1":function(e,t,n){},7638:function(e,t,n){},"9f76":function(e,t,n){},ba1f:function(e,t,n){"use strict";n("9f76")},c69f:function(e,t,n){}});
//# sourceMappingURL=app.7c26f925.js.map