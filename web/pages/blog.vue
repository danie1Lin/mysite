<template>
  <div>
    <b-row>
      <b-col>
        <b-input-group>
          <b-input v-model="title" />
          <b-button @click="add" />
        </b-input-group>
      </b-col>
      <b-col>
        <RightColumn :articles="articles" height="300" />
      </b-col>
    </b-row>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import RightColumn from '~/components/RightColumn'

export default {
  components: {
    RightColumn
  },
  data() {
    return { title: 'defult', windowHeight: 0 }
  },
  computed: mapGetters({
    articles: 'articles/getList'
  }),
  methods: {
    add() {
      this.$store.commit('articles/add', { title: this.title })
    }
  }
}
</script>
