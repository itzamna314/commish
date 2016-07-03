import Ember from 'ember';

export default Ember.Controller.extend({
  actions: {
    selected(player) {
      this.set('selectedPlayer', player);
    },
    create() {
      this.set('selectedPlayer', this.store.createRecord('players'));
    },
    submit() {
      this.get('selectedPlayer').save();
      this.set('selectedPlayer', null);
    }
  }
});
