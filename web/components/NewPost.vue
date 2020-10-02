<template>
  <b-container>
    <b-col cols="1"></b-col>
    <b-col>
      <h1 style="margin: 10px 0px 10px 0px">新增</h1>
      <b-row>
        <b-col cols="2"><label for="post-title-input">文章標題</label></b-col>
        <b-col cols="5"
          ><b-form-input id="post-title-input" v-model="post.title"
        /></b-col>
      </b-row>
      <b-row>
        <b-col cols="2">
          <label for="post-tag-input">文章標籤</label>
        </b-col>
        <b-col cols="5">
          <vue-tags-input
            id="post-tag-input"
            class
            v-model="tag"
            :tags="post.tags"
            :autocomplete-items="tags"
            @tags-changed="newTags => addTag(newTags)"
          />
        </b-col>
      </b-row>
      <b-row>
        <label class="col-6" for="post-content-input">內容</label>
        <label class="col-6" for="post-content-input">預覽</label>
      </b-row>
      <b-row>
        <b-textarea
          id="post-content-input"
          @update="updatePreview()"
          v-model="post.content"
          :rows="height / 40"
          class="col-6"
        />
        <div v-html="postPreview" class="col-6 post-content preview"></div>
      </b-row>
      <b-row>
        <b-button-group col="12">
          <b-button @click="save(post)">
            Save
          </b-button>
          <b-button @click="add()">
            Publish
          </b-button>
        </b-button-group>
      </b-row>
    </b-col>
    <b-col cols="1"></b-col>
  </b-container>
</template>

<script>
import { mapGetters, mapMutations } from 'vuex'
import VueTagsInput from '@johmun/vue-tags-input'
import lodash from 'lodash'
import emoji from 'markdown-it-emoji'
import Markdown from 'markdown-it'
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
      tag: '',
      postPreview: ''
    }
  },
  computed: {
    ...mapGetters({
      articles: 'articles/getList',
      tags: 'articles/getTags'
    })
  },
  mounted() {
    window.addEventListener('resize', () => {
      this.height = window.innerHeight
    })
    this.height = window.innerHeight
  },
  methods: {
    ...mapMutations({
      save: 'articles/savePost'
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
    },
    updatePreview() {
      const md = new Markdown({
        html: false,
        linkify: true,
        typographer: true
      }).use(emoji)
      this.postPreview = md.render(this.post.content)
    }
  }
}
</script>
<style>
.row + .row {
  margin-top: 0.5rem;
}

.vue-tags-input {
  background: #324652;
}

.vue-tags-input .ti-new-tag-input {
  background: transparent;
  color: #b7c4c9;
}

.vue-tags-input .ti-input {
  padding: 4px 10px;
  transition: border-bottom 200ms ease;
}

/* we cange the border color if the user focuses the input */
.vue-tags-input.ti-focus .ti-input {
  border: 1px solid #ebde6e;
}

/* some stylings for the autocomplete layer */
.vue-tags-input .ti-autocomplete {
  background: #fafafa;
  border: 1px solid #8b9396;
  border-top: none;
}

/* the selected item in the autocomplete layer, should be highlighted */
.vue-tags-input .ti-item.ti-selected-item {
  background: #ebde6e;
  color: #283944;
}

/* style the placeholders color across all browser */
.vue-tags-input ::-webkit-input-placeholder {
  color: #a4b1b6;
}

.vue-tags-input ::-moz-placeholder {
  color: #a4b1b6;
}

.vue-tags-input :-ms-input-placeholder {
  color: #a4b1b6;
}

.vue-tags-input :-moz-placeholder {
  color: #a4b1b6;
}

/* default styles for all the tags */
.vue-tags-input .ti-tag {
  position: relative;
  background: #ebde6e;
  color: #283944;
}

/* we defined a custom css class in the data model, now we are using it to style the tag */
.vue-tags-input .ti-tag.custom-class {
  background: transparent;
  border: 1px solid #ebde6e;
  color: #ebde6e;
  margin-right: 4px;
  border-radius: 0px;
  font-size: 13px;
}

/* the styles if a tag is invalid */
.vue-tags-input .ti-tag.ti-invalid {
  background-color: #e88a74;
}

/* if the user input is invalid, the input color should be red */
.vue-tags-input .ti-new-tag-input.ti-invalid {
  color: #e88a74;
}

/* if a tag or the user input is a duplicate, it should be crossed out */
.vue-tags-input .ti-duplicate span,
.vue-tags-input .ti-new-tag-input.ti-duplicate {
  text-decoration: line-through;
}

/* if the user presses backspace, the complete tag should be crossed out, to mark it for deletion */
.vue-tags-input .ti-tag:after {
  transition: transform 0.2s;
  position: absolute;
  content: '';
  height: 2px;
  width: 108%;
  left: -4%;
  top: calc(50% - 1px);
  background-color: #000;
  transform: scaleX(0);
}

.vue-tags-input .ti-deletion-mark:after {
  transform: scaleX(1);
}
</style>
