import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'
import RootHome from '../views/root/RootHome.vue'

const routes = [{
  path: '/',
  name: 'Home',
  component: Home
},
{
  path: '/about',
  name: 'About',
  component: () =>
    import( /* webpackChunkName: "about" */ '../views/About.vue')
},
{
  path: '/test',
  name: 'Test',
  component: () =>
    import('../views/Test.vue')
},
{
  path: '/root_home',
  name: 'RootHome',
  component: RootHome,
  children: [{
    path: "/root_home/add_user",
    name: "add-user",
    component: () =>
      import("../components/root/page/AddUserPage.vue")
  },
  {
    path: "/root_home/get_user",
    name: "get-user",
    component: () =>
      import("../components/root/page/GetUserListPage.vue")
  },
  {
    path: "/root_home/add_equipment",
    name: "add-equipment",
    component: () =>
      import("../components/root/page/AddEquipmentPage.vue")
  },
  {
    path: "/root_home/get_equipment",
    name: "get-equipment",
    component: () =>
      import("../components/root/page/GetEquipmentListPage.vue")
  },
  {
    path: "/root_home/add_group",
    name: "add-group",
    component: () =>
      import("../components/root/page/AddGroupPage.vue")
  },
  {
    path: "/root_home/add_equipment_type",
    name: "add-equipment-type",
    component: () =>
      import("../components/root/page/AddEquipmentTypePage.vue")
  }
  ]
}
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router