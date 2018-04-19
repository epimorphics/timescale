<template>
  <div>
    <el-card class="logincard">
      <el-form @submit.prevent="submitLogin" id="login" ref="login" :model="loginData" label-width="200px" label-position="top">
        <el-form-item label="Username">
          <el-input type="text" v-model="loginData.username"></el-input>
        </el-form-item>
        <el-form-item label="Password">
          <el-input type="password" v-model="loginData.password"></el-input>
        </el-form-item>
        <el-button type="submit" @click.native="submitLogin">Submit</el-button>
      </el-form>
    </el-card>
  </div>
</template>
<script>
import request from 'superagent'
import Cookies from 'js-cookie'
export default {
  name: 'login',
  data () {
    return {
      loginData: {
        username: '',
        password: '',
      }
    }
  },
  methods: {
    submitLogin () {
      console.log(this.loginData)
      request.post('/token')
        .send(this.loginData)
        .then((resp) => {
          console.log('success')
          console.log(resp)
          Cookies.set('token', resp.body.token, {expires: 1})
          this.$router.push({ path: '/' })
        })
        .catch(() => {
          console.log('failure')
        })
    }
  }
}
</script>
<style>
  .logincard {
    width: 20em;
    margin: auto auto;
  }
</style>
