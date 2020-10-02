<template>
  <li :class="'timeline-item ' + item.action_needed">
    <div :class="'timeline-badge ' + item.icon_status">
      <i :class="item.icon_class"></i>
    </div>
    <div
      :class="
        'timeline-panel ' + item.element_status + ' ' + item.element_day_marker
      "
    >
      <div class="timeline-heading">
        <h4 :class="'timeline-title ' + item.text_status">
          <div v-html="item.title"></div>
        </h4>
        <div class="timeline-panel-controls">
          <div class="controls">
            <TimelineControl
              v-for="control in item.controls"
              :control="control"
              :key="control.href"
            />
          </div>
          <div class="timestamp">
            <small class="">{{ item.created }}</small>
          </div>
        </div>
      </div>
      <div class="timeline-body"><div v-html="item.body"></div></div>
    </div>
  </li>
</template>
<script>
import TimelineControl from '~/components/timeline/TimelineControl.vue'
export default {
  components: {
    TimelineControl
  },
  props: ['item'],

  methods: {
    delete: function() {
      this.$dispatch('timeline-delete-item', this.item.id)
    },

    edit: function() {}
  },

  events: {
    'timeline-delete': function() {
      this.delete()
    },
    'timeline-edit': function() {
      this.edit()
    }
  }
}
</script>
