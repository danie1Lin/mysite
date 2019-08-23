import marked from 'marked'
import DOMPurify from 'DOMPurify'

export const state = () => ({
  classes: [{ name: 'vue', id: '1' }, { name: 'golang', id: '2' }],
  list: [{ title: 'hi', content: 'haha' }, { title: 'yo', content: 'nono' }],
  title: 'xsaxa',
  post: {
    title: 'test',
    content: `
How To Use The Demo
-------------------
**Preview:**
<p>我有毒<iframe//src=jAva&Tab;script:alert(3)>def
dcs`
  }
})

export const mutations = {
  add(state, post) {
    state.list.push(post)
  }
}

export const getters = {
  getClasses(state) {
    return state.classes
  },
  getList(state) {
    return state.list
  },
  getPost(state) {
    const cleanHtml = DOMPurify.sanitize(marked(state.post.content))
    return { content: cleanHtml, title: state.post.title }
  }
}
