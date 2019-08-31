import Vue from 'vue'
import VueLodash from 'vue-lodash'
import lodash from 'lodash'

const options = { name: 'lodash' } // customize the way you want to call it

Vue.use(VueLodash, lodash, options) // options is optional
