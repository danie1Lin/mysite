export const state = () => ({
  list: [{ name: 'hi' }, { name: 'yo' }],
  title: 'xsaxa'
})

export const mutations = {
  add(state, article) {
    state.list.push({ name: article.title })
  }
}

export const getters = {
  getList(state) {
    return state.list
  }
}
