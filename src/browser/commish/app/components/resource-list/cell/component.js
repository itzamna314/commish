import Ember from 'ember';

export default Ember.Component.extend({
  mouseEnter () {
    this.set('record.isHovered', true);
  },
  mouseLeave() {
    this.set('record.isHovered', false);
  }
});
