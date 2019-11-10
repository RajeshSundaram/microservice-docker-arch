<template lang="pug">
    .login-container
        el-form(:model="loginCredintials" label-width="60px" ref="form")
            el-form-item(label="Username")
                el-input(placeholder="hello" v-model="loginCredintials.username" size="mini")
            el-form-item(label="Password")
                el-input(placeholder="password" v-model="loginCredintials.password" show-password size="mini")
            el-form-item.login-actions
                el-button(@click="signin") Login
                el-button(@click="signup") Signup
</template>>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({
  data () {
    return {
      loginCredintials: {
        username: '',
        password: ''
      }
    }
  },
  methods: {
    signin: () => {
      console.log('signing in')
    },
    signup: function () {
      this.$router.push('/signup')
    }
  }
})
</script>

<style lang="scss" scoped>
.login-container {
  // border: 2px solid;
  text-align: center;
  .el-form {
    margin: 0 auto;
    height: 100%;
    width: 350px;
    background: lightblue;
    .el-form-item {
      display: flex;
      flex-direction: row;
      justify-content: center;
      &.login-actions .el-button{
        margin: 0px 10px;
      }
    }
  }
}
</style>
