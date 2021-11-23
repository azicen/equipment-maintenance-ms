import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue';
import LoginPage from "@/views/page/LoginPage.vue";
import CreateMaintainForm from "@/views/component/form/CreateMaintainForm.vue";
// import LoginForm from "@/views/component/form/LoginForm";

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/about',
    name: 'About',
    component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
  },
  {
    path: '/login',
    name: 'Login',
    component: LoginPage,
  },
  {
    path: '/test',
    name: 'Test',
    component: CreateMaintainForm,
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
