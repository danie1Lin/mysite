importScripts('/_nuxt/workbox.4c4f5ca6.js')

workbox.precaching.precacheAndRoute([
  {
    "url": "/_nuxt/108545d4065890fe08d0.js",
    "revision": "384f3cc2f6c274fcfdb4426103738f42"
  },
  {
    "url": "/_nuxt/1bac82a6435d8583db32.js",
    "revision": "c22fa859fa81bcfb330789a8e77dc5c2"
  },
  {
    "url": "/_nuxt/2eb683672a061aa7b858.js",
    "revision": "74cec01de0884f716baeba3dd6cdfb4b"
  },
  {
    "url": "/_nuxt/3895418bcb94f61ad2f8.js",
    "revision": "2b2a87f19ede09997ab22199c06ce55c"
  },
  {
    "url": "/_nuxt/3a6a91908214c03d68e8.js",
    "revision": "a02bc4d3589e81c83957ffeb1e2deb77"
  },
  {
    "url": "/_nuxt/4bc416e3e05f556bc9b4.js",
    "revision": "359ca1ed4ce5e9544efcb77aef6c546f"
  },
  {
    "url": "/_nuxt/512453d93889fff09da0.js",
    "revision": "99cb9a3374a296529feecd11bcfff6d1"
  },
  {
    "url": "/_nuxt/65cbc723d640bbe326fd.js",
    "revision": "6c42ca438234f11fc75e39a28c5e4487"
  },
  {
    "url": "/_nuxt/8b75b40422229508b2c8.js",
    "revision": "4ebd4c6ae82ef63e516176a856dc42c3"
  },
  {
    "url": "/_nuxt/a6f673b3dfb4994631b0.js",
    "revision": "dafd5b18ac4560e04e118c952cdcd04a"
  },
  {
    "url": "/_nuxt/d0e0d5a1da4420f4d394.js",
    "revision": "b97af5ff608a3b81fb615e26ed41e189"
  }
], {
  "cacheId": "mysite",
  "directoryIndex": "/",
  "cleanUrls": false
})

workbox.clientsClaim()
workbox.skipWaiting()

workbox.routing.registerRoute(new RegExp('/_nuxt/.*'), workbox.strategies.cacheFirst({}), 'GET')

workbox.routing.registerRoute(new RegExp('/.*'), workbox.strategies.networkFirst({}), 'GET')
