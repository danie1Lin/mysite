<template>
  <b-row id="login-panel">
    <b-col cols="12"><h1>Login</h1></b-col>
    <b-col></b-col>
    <b-col cols="10">
      <label>Email</label>
      <b-form-input type="email" v-model="username" />
      <label>Password</label>
      <b-form-input type="password" v-model="password" />
      <b-button
        md="auto"
        @click="login({ username, password })"
        variant="primary"
      >
        LOGIN
      </b-button>
    </b-col>
    <b-col></b-col>
  </b-row>
</template>

<script>
import { mapActions } from 'vuex'
export default {
  data() {
    return {
      password: '',
      username: ''
    }
  },
  methods: mapActions('auth', ['login'])
}
</script>
