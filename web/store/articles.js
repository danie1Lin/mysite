import marked from 'marked'
import DOMPurify from 'DOMPurify'
import lodash from 'lodash'
export const state = () => ({
  tags: [],
  list: [],
  post: null
})

export const mutations = {
  addTag(state, tag) {
    // TODO: post tag/new
    tag.id = state.tags.length
    state.tags.push(tag)
  },
  addPost(state, post) {
    post.id = state.list.length
    state.list.push(post)
  },
  selectPost(state, post) {
    const find = lodash.find(state.list, i => {
      return i.id === post.id
    })
    state.post = find
  }
}

export const getters = {
  getTags(state) {
    return state.tags
  },
  getList(state) {
    return state.list
  },
  getPost(state) {
    if (!state.post) {
      return null
    }
    let cleanHtml = DOMPurify.sanitize(marked(state.post.content))
    const dom = new DOMParser().parseFromString(cleanHtml, 'text/html')

    const h1s = dom.getElementsByTagName('h1')
    const navs = {}
    for (let i = 0; i < h1s.length; i++) {
      const id = 'h1-' + i.toString()
      navs[id] = h1s[i].innerText
      h1s[i].id = id
    }
    console.log(dom)
    console.log(navs)
    cleanHtml = dom.documentElement.innerHTML
    return { content: cleanHtml, title: state.post.title, navs: navs }
  }
}
