import Vue from 'vue';
import Router from 'vue-router';
import Home from './views/Home.vue';
import Timescale from './views/Timescale.vue';
import Login from './views/Login.vue';

Vue.use(Router);

export default new Router({
  routes: [
    {
      path: '/',
      name: 'home',
      component: Timescale,
    },
    {
      path: '/login',
      name: 'login',
      component: Login,
    }
  ],
});
