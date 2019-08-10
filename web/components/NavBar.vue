<template>
  <b-navbar pills>
    <b-navbar-brand>{{ brandName }}</b-navbar-brand>
    <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>
    <b-collapse id="nav-collapse" is-nav>
      <b-navbar-nav v-for="(item, index) in leftItems" v-bind:key="index">
        <b-nav-item v-bind:href="item.href">{{ item.name }}</b-nav-item>
      </b-navbar-nav>
      <b-navbar-nav
        class="ml-auto"
        v-for="(item, index) in rightItems"
        v-bind:key="index"
      >
        <b-nav-item v-bind:href="item.href" @click="item.click">{{
          item.name
        }}</b-nav-item>
      </b-navbar-nav>
    </b-collapse>
  </b-navbar>
</template>
<script>
export default {
  props: ['brandName', 'leftItems', 'rightItems'],
  methods: {}
}
</script>
