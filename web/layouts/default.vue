<template>
  <div>
    <nav-bar :login="login" :userInfo="userInfo" />
    <b-container fluid id="nuxt-page">
      <nuxt />
    </b-container>
  </div>
</template>

<script>
import { mapState } from 'vuex'
import NavBar from '~/components/NavBar.vue'

export default {
  components: {
    NavBar
  },
  data() {
    return {}
  },
  computed: mapState({
    userInfo: state => state.auth.userInfo
  }),
  methods: {
    async login() {
      await this.$store.dispatch('login', {
        username: this.email,
        password: this.password
      })
    },
    handleAuthentication(data) {
      console.log(this.$cookie.get('session-id'))
    }
  }
}
</script>

<style>
html {
  font-family: 'Source Sans Pro', -apple-system, BlinkMacSystemFont, 'Segoe UI',
    Roboto, 'Helvetica Neue', Arial, sans-serif;
  font-size: 16px;
  word-spacing: 1px;
  -ms-text-size-adjust: 100%;
  -webkit-text-size-adjust: 100%;
  -moz-osx-font-smoothing: grayscale;
  -webkit-font-smoothing: antialiased;
  box-sizing: border-box;
}

#nuxt-page {
  margin-top: 20px;
}
</style>
