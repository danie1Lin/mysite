<template>
  <b-list-group :style="{ 'overflo-y': 'auto', height: height }">
    <b-list-group-item v-for="(article, index) in articles" :key="index">{{
      article.name
    }}</b-list-group-item>
  </b-list-group>
</template>

<script>
export default {
  props: ['articles', 'height']
}
</script>
