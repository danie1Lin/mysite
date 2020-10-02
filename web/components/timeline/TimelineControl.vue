<template>
  <a href="#" @click="handleClick">
    <i :class="control.icon_class"></i>
  </a>
</template>
<script>
export default {
  props: ['control'],

  methods: {
    handleClick() {
      if (this.control.method === 'delete') {
        this.$dispatch('timeline-delete')
      } else if (this.control.method === 'edit') {
        this.$dispatch('timeline-edit')
      } else {
        console.log('Unknown method ' + this.control.method)
      }
    }
  }
}
</script>
