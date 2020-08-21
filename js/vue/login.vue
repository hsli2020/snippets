// login.vue
<template>
  <div>
    <input type="text" v-model="loginForm.username" placeholder="用户名"/>
    <input type="text" v-model="loginForm.password" placeholder="密码"/>
    <button @click="login">登录</button>
  </div>
</template>
 
<script>
import { mapMutations } from 'vuex';

export default {
  data () {
    return {
      loginForm: { username: '', password: '' },
      userToken: ''
    };
  },
 
  methods: {
    ...mapMutations(['changeLogin']),
    login () {
      let _this = this;
      if (this.loginForm.username === '' || this.loginForm.password === '') {
        alert('账号或密码不能为空');
      } else {
        this.axios({ method: 'post', url: '/user/login', data: _this.loginForm })
        .then(res => {
          console.log(res.data);
          _this.userToken = 'Bearer ' + res.data.data.body.token;
          // 将用户token保存到vuex中
          _this.changeLogin({ Authorization: _this.userToken });
          _this.$router.push('/home');
          alert('登陆成功');
        }).catch(error => {
          alert('账号或密码错误');
          console.log(error);
        });
      }
    }
  }
};
</script>

// store文件夹下的index.js
import Vue from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);
 
const store = new Vuex.Store({
  state: {
	// 存储token
	Authorization: localStorage.getItem('Authorization') || : ''
  },
 
  mutations: {
	// 修改token，并将token存入localStorage
	changeLogin (state, user) {
	  state.Authorization = user.Authorization;
	  localStorage.setItem('Authorization', user.Authorization);
	}
  }
});
 
export default store;

// router文件夹下的index.js
import Vue from 'vue';
import Router from 'vue-router';
import login from '@/components/login';
import home from '@/components/home';
 
Vue.use(Router);
 
const router = new Router({
  routes: [
	{ path: '/',      redirect: '/login' },
	{ path: '/login', name: 'login', component: login },
	{ path: '/home',  name: 'home',  component: home }
  ]
});
 
// 使用 router.beforeEach 注册一个全局前置守卫，判断用户是否登陆
router.beforeEach((to, from, next) => {
  if (to.path === '/login') {
	next();
  } else {
	let token = localStorage.getItem('Authorization');
 
	if (token === null || token === '') {
	  next('/login');
	} else {
	  next();
	}
  }
});
 
export default router;

// 请求头加token // 添加请求拦截器，在请求头中加token
axios.interceptors.request.use(
  config => {
	if (localStorage.getItem('Authorization')) {
	  config.headers.Authorization = localStorage.getItem('Authorization');
	}
	return config;
  },
  error => {
	return Promise.reject(error);
  }
);

// 如果前端拿到状态码为401，就清除token信息并跳转到登录页面

localStorage.removeItem('Authorization');
this.$router.push('/login');
