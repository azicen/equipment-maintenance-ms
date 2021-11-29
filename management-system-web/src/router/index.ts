import {createRouter, createWebHistory, RouteRecordRaw} from 'vue-router'
import {adminRoutes} from "@/router/admin";

import Home from '@/views/Home.vue'
import Admin from "@/views/Admin.vue"
import LoginPage from "@/views/Login.vue"
import TestPage from "@/views/TestView.vue"

export const routes: RouteRecordRaw[] = [
    {
        path: '/',
        name: 'Home',
        component: Home,
    },
    {
        path: '/admin',
        name: 'Admin',
        component: Admin,
    },
    {
        path: '/login',
        name: 'Login',
        component: LoginPage,
    },
    {
        path: '/test',
        name: 'Test',
        component: TestPage,
    },
    {
        path: '/admin',
        name: 'Admin',
        component: Admin,
        children: adminRoutes,
    }
]

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes
})

export default router
