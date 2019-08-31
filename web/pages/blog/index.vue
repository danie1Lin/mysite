<template>
  <div>
    <b-row>
      <b-col cols="2" id="post-list">
        <b-navbar class="flex-column">
          <b-navbar-brand>
            文章列表
          </b-navbar-brand>
          <b-nav pills>
            <b-nav-item
              v-for="(article, index) in articles"
              :key="index"
              @click="selectPost(article)"
              >{{ article.title }}</b-nav-item
            >
          </b-nav>
        </b-navbar>
      </b-col>
      <b-col id="post-content">
        <div>
          <template v-if="post">
            <h1>{{ _.get(post, 'title') }}</h1>
            <div v-html="_.get(post, 'content')"></div>
          </template>
        </div>
      </b-col>
      <b-col v-b-scrollspy:post-content id="post-outline" cols="2">
        <b-navbar class="flex-column">
          <b-navbar-brand>
            Outline
          </b-navbar-brand>
          <b-nav pills class="flex-column">
            <template v-if="post">
              <b-nav-item
                class="post-outline-item"
                v-for="(text, id) in _.get(post, 'navs')"
                :key="id"
                :href="'#' + id"
              >
                {{ text }}
              </b-nav-item>
            </template>
          </b-nav>
        </b-navbar>
      </b-col>
    </b-row>
  </div>
</template>

<script>
import { mapGetters, mapMutations } from 'vuex'

export default {
  data() {
    return { title: 'defult' }
  },
  methods: mapMutations({
    selectPost: 'articles/selectPost'
  }),
  computed: mapGetters({
    articles: 'articles/getList',
    tags: 'articles/getTags',
    post: 'articles/getPost'
  }),
  mounted() {
    document.getElementById('post-outline')
  }
}
</script>
<style>
#post-list {
  border-right: 3px solid;
  border-color: #b8daff;

  position: -webkit-sticky;
  position: sticky;
  top: 4rem;
  height: calc(100vh - 4rem);
}

#post-content {
  position: -webkit-sticky;
  position: sticky;
  overflow-y: auto;
  top: 4rem;
  height: calc(100vh - 4rem);
}

#post-outline {
  border-left: 3px solid;
  border-color: #b8daff;
  overflow-y: auto;
  position: -webkit-sticky;
  position: sticky;
  top: 4rem;
  height: calc(100vh - 4rem);
}
</style>
