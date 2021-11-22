import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue';
import WordRotation from "@/views/component/data-display/WordRotation";
import LoginPage from "@/views/page/LoginPage";
import TestPage from "@/views/page/TestPage";

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/about',
    name: 'About',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
  },
  {
    path: '/login',
    name: 'Login',
    component: LoginPage,
  },
  {
    path: '/test',
    component: TestPage,
    children: [
      {
        // 当 /user/:id/profile 匹配成功，
        // UserProfile 会被渲染在 User 的 <router-view> 中
        path: 'word-rotation',
        component: WordRotation
      },
      // {
      //   path: 'posts',
      //   component: UserPosts
      // }
    ]
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
