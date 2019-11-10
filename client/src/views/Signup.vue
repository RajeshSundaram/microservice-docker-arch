<template lang="pug">
    .signup-container
        .signup-navigation
            el-link(@click="navToLogin") back to login
        .signup-content
            el-form(ref="form" :model="signupForm" label-width="100px")
                el-form-item(v-if="errors.length")
                  ul.error-class
                    li(v-for="err in errors") {{err}}
                el-form-item(label="Username")
                    el-input(v-model="signupForm.username")
                el-form-item(label="Password")
                    el-input(v-model="signupForm.password")
                el-form-item(label="DateOfBirth")
                    el-date-picker(type="date" v-model="signupForm.dob" placeholder="DOB")
                el-form-item.signup-actions
                    el-button(@click="signup") Signup

</template>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({
  data () {
    return {
      errors: [],
      signupForm: {
        username: '',
        password: '',
        dob: undefined
      },
      messages: [
        'Username required',
        'Password required',
        'Date of birth required'
      ]
    }
  },
  methods: {
    navToLogin: function () {
      this.$router.go(-1)
    },
    signup: async function () {
      if (this.signupForm.username === '') {
        this.addIfNotIn(this.errors, this.messages[0])
      }
      if (this.signupForm.password === '') {
        this.addIfNotIn(this.errors, this.messages[1])
      }
      if (this.signupForm.dob === '') {
        this.addIfNotIn(this.errors, this.messages[2])
      }
      if (this.errors.length === 0) {
        Vue.axios.post('http://localhost:8080/auth/signup', this.signupForm)
          .then(res => console.log(res))
          .catch(e => console.error(e))
      }
    },
    addIfNotIn (array :Array<String>, element: String) {
      if (!array.includes(element)) {
        array.push(element)
      }
      return array
    }
  }
})
</script>

<style lang="scss" scoped>
.signup-container {
  text-align: center;
  .signup-content {
    padding: 30px;
    background: lightblue;
    width: 400px;
    margin: 0 auto;
    ul.error-class {
      color: red;
      list-style-type: square;
    }
  }
}
</style>
