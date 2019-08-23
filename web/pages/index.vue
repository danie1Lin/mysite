<template>
  <b-container>
    <logo class="row justify-content-center" />
    <h1 class="title">
      Daniel Lin's Blog
    </h1>
    <h2 class="subtitle">
      my personal site
    </h2>
    <div class="links">
      <a
        href="https://github.com/daniel840829"
        target="_blank"
        class="button--grey"
      >
        GitHub
      </a>
    </div>
  </b-container>
</template>
<script>
import Logo from '~/components/Logo.vue'
export default {
  components: {
    Logo
  }
}
</script>

<style>
.container {
  margin: 0 auto;
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: center;
  flex-direction: row;
}

.title {
  font-family: 'Quicksand', 'Source Sans Pro', -apple-system, BlinkMacSystemFont,
    'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif;
  display: block;
  font-weight: 300;
  font-size: 100px;
  color: #35495e;
  letter-spacing: 1px;
}

.subtitle {
  font-weight: 300;
  font-size: 42px;
  color: #526488;
  word-spacing: 5px;
  padding-bottom: 15px;
}

.links {
  padding-top: 15px;
}
</style>
