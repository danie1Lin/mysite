<template>
  <b-navbar toggleable="md" type="dark" variant="info" sticky>
    <b-navbar-brand href="/">Daniel's Blog</b-navbar-brand>
    <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>
    <b-collapse id="nav-collapse" is-nav>
      <b-navbar-nav>
        <b-nav-item href="/blog">Blog</b-nav-item>
        <b-nav-item href="/gallery">Gallery</b-nav-item>
        <!-- <b-nav-item href="/blog/new">New Post</b-nav-item> -->
      </b-navbar-nav>
      <!-- <b-navbar-nav class="ml-auto"> -->
      <!--   <b-nav-item v-show="_.isNull(userInfo)" href="/login" right -->
      <!--     >Login</b-nav-item -->
      <!--   > -->
      <!--   <b-nav-item-dropdown -->
      <!--     v-show="!_.isNull(userInfo)" -->
      <!--     :text="_.get(userInfo, 'userName')" -->
      <!--     right -->
      <!--   > -->
      <!--     <b-dropdown-item href="/account">Account</b-dropdown-item> -->
      <!--     <b-dropdown-item @click="logout()">log out</b-dropdown-item> -->
      <!--   </b-nav-item-dropdown> -->
      <!-- </b-navbar-nav> -->
    </b-collapse>
  </b-navbar>
</template>
<script>
import { mapState, mapActions } from 'vuex'
export default {
  computed: mapState({
    userInfo: state => {
      return state.auth.userInfo
    }
  }),
  methods: mapActions('auth', ['logout'])
}
</script>
