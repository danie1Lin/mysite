// import axios from 'axios'

export default {
  state() {
    return { userInfo: null }
  },
  mutations: {
    setUser(state, userInfo) {
      state.userInfo = userInfo
    }
  },
  actions: {
    // nuxtServerInit is called by Nuxt.js before server-rendering every page
    login({ commit }, { username, password }) {
      // const { data } = await axios.post('api/login', {
      //   email: username,
      //   pswd: password
      // })
      // TODO: api
      commit('setUser', {
        userName: 'daniel',
        role: 'admin',
        sessionID: 'ghxdjsjknxk'
      })
    },
    logout({ commit }) {
      // await axios.post('/api/logout')
      commit('setUser', null)
    }
  }
}
