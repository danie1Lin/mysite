<template>
  <b-row>
    <b-col cols="1"></b-col>
    <b-col>
      <h1>New Post</h1>
      <b-form>
        <label>Title</label>
        <b-form-input v-model="post.title" />
        <label>tags</label>
        <vue-tags-input
          v-model="tag"
          :tags="post.tags"
          :autocomplete-items="tags"
          @tags-changed="newTags => addTag(newTags)"
        />
        <label>Content</label>
        <b-form-textarea v-model="post.content" :rows="height / 40" />
        <b-button @click="add(post)">
          Save
        </b-button>
        <b-button @click="add">
          Publish
        </b-button>
      </b-form>
    </b-col>
    <b-col cols="1"></b-col>
  </b-row>
</template>

<script>
import { mapGetters, mapMutations } from 'vuex'
import VueTagsInput from '@johmun/vue-tags-input'
import lodash from 'lodash'

export default {
  components: {
    VueTagsInput
  },
  data() {
    return {
      height: 0,
      post: {
        title: '',
        content: '',
        tags: []
      },
      tag: ''
    }
  },
  computed: mapGetters({
    articles: 'articles/getList',
    tags: 'articles/getTags'
  }),
  mounted() {
    window.addEventListener('resize', () => {
      this.height = window.innerHeight
    })
    this.height = window.innerHeight
  },
  methods: {
    ...mapMutations({
      add: 'articles/addPost'
    }),
    addTag(PostTags) {
      const newPostTag = PostTags.map(element => {
        if (!lodash.hasIn(element, 'id')) {
          this.$store.commit('articles/addTag', element)
          return element
        } else {
          return element
        }
      })
      this.post.tags = newPostTag
    }
  }
}
</script>
