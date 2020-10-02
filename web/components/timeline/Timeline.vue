<template>
  <ul class="timeline">
    <TimelineItem
      v-for="(item, index) in items"
      is="timeline-item"
      :item="item"
      :key="index"
    />
  </ul>
</template>
<script>
import TimelineItem from '~/components/Timeline/TimelineItem.vue'
export default {
  components: {
    TimelineItem
  },

  props: ['items'],

  events: {
    'delete-item': function() {
      return true // forward to parent
    }
  }
}
</script>
