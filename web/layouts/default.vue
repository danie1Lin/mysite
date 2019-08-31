<template>
  <div>
    <nav-bar id="header-nav" :login="login" :userInfo="userInfo" />
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
    handleAuthentication(data) {}
  }
}
</script>

<style>
* {
  font-family: 'Helvetica', 'Arial', 'LiHei Pro', '黑體-繁', '微軟正黑體',
    sans-serif;
}

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
}

#login-panel {
  padding: 30px;
  margin: 10px;
  background: #90d2dd;
  border: #27a1b8;
  border-radius: 20px;
}
button {
  margin-top: 10px;
}

input + label {
  margin-top: 10px;
}

.blog-post-content {
  padding: 30px;
}
::-webkit-scrollbar {
  width: 0px;
}
::-webkit-scrollbar-track {
  box-shadow: inset 0 0 5px grey;
  border-radius: 0px;
}

#header-nav {
  height: 4rem;
  top: 0px;
  position: sticky;
  position: -webkit-sticky;
}
</style>
