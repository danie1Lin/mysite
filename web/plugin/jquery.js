import Vue from 'vue'
import jquery from 'jquery'

Vue.use({
  install(Vue, options) {
    Vue.prototype.$ = jquery
  }
}) // options is optional
