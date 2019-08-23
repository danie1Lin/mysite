import qs from 'qs'

export default function({ $axios, app }) {
  // console.log(app.$cookies.get('token'))  //每次請求獲取cooie
  $axios.onRequest(config => {
    // 獲取cookie放在頭部傳到後端
    config.headers.session = app.$cookies.get('session-id')
    console.log(app.$cookies.get('session-id'))
    config.data = qs.stringify(config.data, {
      allowDots: true // Option allowDots can be used to enable dot notation
    })
    return config
  })

  $axios.onResponse(response => {
    return Promise.resolve(response.data)
  })

  $axios.onError(error => {
    return Promise.reject(error)
  })
}
