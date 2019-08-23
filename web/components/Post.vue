<template>
  <div>
    <h1>{{ blog.title }}</h1>
    <div v-html="blog.content"></div>
  </div>
</template>

<script>
export default {
  props: ['blog']
}
</script>
