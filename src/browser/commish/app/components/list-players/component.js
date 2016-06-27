import Ember from 'ember';

export default Ember.Component.extend({
  actions: {
    selected (player) {
      this.get('selected')(player);
    }
  }
});
