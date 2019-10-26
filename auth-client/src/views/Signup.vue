<template lang="pug">
    .signup-container
        .signup-navigation
            el-link(@click="navToLogin") back to login
        .signup-content
            el-form(ref="form" :model="signupForm" label-width="100px")
                el-form-item(label="username")
                    el-input(v-model="signupForm.username")
                el-form-item(label="password")
                    el-input(v-model="signupForm.password")
                el-form-item(label="Date Of Birth")
                    el-date-picker(type="date" v-model="signupForm.dob" placeholder="DOB")
                el-form-item.signup-actions
                    el-button(@click="signup") Signup

</template>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({
  data () {
    return {
      signupForm: {
        username: '',
        password: '',
        dob: undefined
      }
    }
  },
  methods: {
    navToLogin: function () {
      this.$router.go(-1)
    },
    signup: function () {
      console.dir(this.signupForm)
    }
  }
})
</script>

<style lang="scss" scoped>
</style>
