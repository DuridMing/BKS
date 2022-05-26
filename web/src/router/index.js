import { createRouter, createWebHistory } from 'vue-router'
import home from '../views/pages/HomePage.vue'
import NotFound from '../views/NotFound.vue'
import Login from "../views/login/Login.vue"

import AddBook from '../views/pages/AddBookPage.vue'
import Search from "../views/pages/SearchPage.vue";

import store from "../store"


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
    path: '/addbook',
    name: 'addbook',
    component: AddBook
  },
  // the last router rule.
  {
    path: '/:pathMatch(.*)*',
    name: 'notfound',
    component: NotFound
  }
]

const Routers = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
});

Routers.beforeEach((to) => {
  if (to.name != 'login' && !store.state.isAuthenticated){
    return {name: 'login'}
  }
});

export {Routers}

