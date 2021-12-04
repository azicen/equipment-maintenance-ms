import {RouteRecordRaw} from "vue-router"
import Home from "@/views/admin/Home.vue"
import User from "@/views/admin/User.vue"
import Equipment from "@/views/admin/Equipment.vue"
import Maintain from "@/views/admin/Maintain.vue"

export const adminRoutes: RouteRecordRaw[] = [
    {
        path: "/admin",
        name: "AdminHome",
        component: Home,
    },
    {
        path: "/admin/user",
        name: "AdminUser",
        component: User,
    },
    {
        path: "/admin/equipment",
        name: "AdminEquipment",
        component: Equipment,
    },
    {
        path: "/admin/maintain",
        name: "AdminMaintain",
        component: Maintain,
    },
]