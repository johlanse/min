import { createRouter, createWebHashHistory } from 'vue-router'
import AddUser from "@/views/AddUser";
import UserManager from "@/views/UserManager";
import Activity from "@/views/Activity";
import Log from "@/views/Log";



const routes = [
  {
    path: "/add_user",
    name: "add_user",
    component: AddUser
  },
  {
    path: "/user_manager",
    name: "user_manager",
    component: UserManager
  },
  {
    path: "/activity",
    name: "activity",
    component: Activity
  },
  {
    path: "/log",
    name: "log",
    component: Log
  },
  // {
  //   path: '/',
  //   name: 'Home',
  //   component: Home
  // },
  // {
  //   path: '/about',
  //   name: 'About',
  //   // route level code-splitting
  //   // this generates a separate chunk (about.[hash].js) for this route
  //   // which is lazy-loaded when the route is visited.
  //   component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
  // }
]

const router = createRouter({
  history: createWebHashHistory(process.env.BASE_URL),
  routes
})

export default router
