import Vue from 'vue'
import VueRouter from 'vue-router'
import Main from '../views/Main.vue'
import Settings from '../views/Settings.vue'
import Requests from '../views/Requests.vue'
import Graphs from '../views/Graphs.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Main Page',
    component: Main
  },
  {
    path: '/settings',
    name: 'Settings Page',
    component: Settings
  },
  {
    path: '/requests',
    name: 'Requests Page',
    component: Requests
  },
  {
    path: '/graphs',
    name: 'Graphs Page',
    component: Graphs
  },
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
