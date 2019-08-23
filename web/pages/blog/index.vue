<template>
  <div>
    <b-nav pills>
      <b-nav-item v-for="cls in classes" :key="cls.id">{{
        cls.name
      }}</b-nav-item>
    </b-nav>
    <b-row>
      <b-col class="blog-post-content" md="9">
        <b-container><Post :blog="post"/></b-container>
      </b-col>
      <b-col cols="3">
        <RightColumn :articles="articles" height="300" />
      </b-col>
    </b-row>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import RightColumn from '~/components/RightColumn'
import Post from '~/components/Post'

export default {
  components: {
    RightColumn,
    Post
  },
  data() {
    return { title: 'defult', windowHeight: 0 }
  },
  computed: mapGetters({
    articles: 'articles/getList',
    classes: 'articles/getClasses',
    post: 'articles/getPost'
  }),
  methods: {
    add() {
      this.$store.commit('articles/add', { title: this.title })
    }
  }
}
</script>
