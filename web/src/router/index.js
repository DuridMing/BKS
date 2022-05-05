import { createRouter, createWebHistory } from 'vue-router'
import home from '../views/components/HomePage.vue'
import Dashboard from '../views/components/DashboardItem.vue'
import Search from "../views/components/SearchItem.vue";
import NotFound from '../views/NotFound.vue'
import Login from "../views/login/Login.vue"


const routes = [
  {
    path: '/',
    name: 'index',
    redirect :"/login",
  },
  {
    path: '/login',
    name: 'login',
    component: Login
  },
  {
    path: '/home',
    name: 'home',
    component: home
  },
  {
    path: '/search',
    name: 'search',
    component: Search
  },
  {
    path: '/dashboard',
    name: 'dashboard',
    component: Dashboard
  },
  // the last router rule.
  {
    path: '/:pathMatch(.*)*',
    name: 'notfound',
    component: NotFound
  }
]

export const Routers = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

