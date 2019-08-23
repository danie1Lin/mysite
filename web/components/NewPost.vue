<template>
  <b-form>
    <b-form-input v-model="post.title" />
    <b-form-textarea v-model="post.content" />
    <b-button @click="add(post)">
      Save
    </b-button>
    <b-button @click="add">
      Publish
    </b-button>
  </b-form>
</template>

<script>
import { mapGetters, mapMutations } from 'vuex'

export default {
  state() {
    return {
      post: {
        title: '',
        content: ''
      }
    }
  },
  computed: mapGetters({
    articles: 'articles/getList',
    classes: 'articles/getClasses'
  }),
  methods: mapMutations({ add: 'articles/add' })
}
</script>
