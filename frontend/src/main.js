import Vue from 'vue'
import App from './App.vue'
import router from './router'
import axios from 'axios'
import store from './store'

import 'bootstrap'; 
import 'bootstrap/dist/css/bootstrap.min.css';

import VueAxios from 'vue-axios'

const API_URL = process.env.API_URL || 'http://localhost:5000/'

export const HTTP = axios.create({
  baseURL: API_URL,
  headers: {
    'Content-Type': 'application/json',
    'Authorization': 'Bearer ' + localStorage.token
  }
})

Vue.use(VueAxios, HTTP)

Vue.config.productionTip = false

new Vue({
  router,
  axios,
  store,
  render: h => h(App)
}).$mount('#app')
